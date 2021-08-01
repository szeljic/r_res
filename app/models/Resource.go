package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"strconv"
	"time"
)

func SaveResource(data map[string]interface{}, user User, categoryID int) error {

	var category Category
	collection := DB.Database(Database).Collection("categories")
	err := collection.FindOne(context.Background(), bson.M{"id": categoryID}).Decode(&category)
	if err != nil {
		log.Println(err)
		return &errorString{err.Error()}
	}
	err = validateDataByCategory(category, data)

	if err != nil {
		return &errorString{err.Error()}
	}

	collection = DB.Database(Database).Collection("resources")

	var set bson.D

	for key, value := range data {
		set = append(set, bson.E{Key: key, Value: value})
	}
	set = append(set, bson.E{Key: "user", Value: user})
	set = append(set, bson.E{Key: "category", Value: category})

	_, err = collection.InsertOne(context.Background(), set)
	if err != nil {
		return err
	}

	return nil
}

func validateDataByCategory(category Category, data map[string]interface{}) error {

	for _, value := range category.SpecificFields {

		if value.Required && (data[value.SCName] == "" || data[value.SCName] == nil) {
			return &errorString{value.Name + " je obavezno polje!"}
		}

		switch value.DataType {
		case "integer":
			_, ok := data[value.SCName].(float64)
			if !ok {
				return &errorString{value.Name + " nije ispravan podatak!"}
			}
			break
		case "float":
			if reflect.TypeOf(data[value.SCName]).Name() != "float64" {
				return &errorString{value.Name + " nije ispravan podatak!"}
			}
			break
		case "string":
			break
		case "text":
			break
		case "date":
			_, err := time.Parse("2006-01-02", data[value.SCName].(string))
			if err != nil {
				return &errorString{value.Name + " nije ispravan podatak!"}
			}
			break
		case "datetime":
			_, err := time.Parse("2006-01-02 15:04:05", data[value.SCName].(string))
			if err != nil {
				return &errorString{value.Name + " nije ispravan podatak!"}
			}
			break
		case "boolean":
			_, ok := data[value.SCName].(bool)
			if !ok {
				return &errorString{value.Name + " nije ispravan podatak!"}
			}
			break

		default:
			return &errorString{"Nepoznat tip podatka!"}
		}
	}

	return nil
}

func GetTotalResources(q string) int {
	collection := DB.Database(Database).Collection("resources")

	filter := bson.M{}
	if q != "" {
		qInt, _ := strconv.Atoi(q)
		filter = bson.M{"$or": []interface{}{
			bson.M{"name": primitive.Regex{Pattern: "^" + q, Options: ""}},
			bson.M{"user.username": primitive.Regex{Pattern: "^" + q, Options: ""}},
			bson.M{"id": qInt},
			bson.M{"user.id": qInt}},
		}
	}

	data, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	return int(data)
}

func GetResources(order, sortBy, q string, paginateBy, page int64) []map[string]interface{} {

	collection := DB.Database(Database).Collection("resources")

	findOptions := options.Find()
	if order == "desc" {
		findOptions.SetSort(bson.D{{sortBy, -1}})
	} else {
		findOptions.SetSort(bson.D{{sortBy, 1}})
	}

	findOptions.SetLimit(paginateBy)
	findOptions.SetSkip(paginateBy * (page - 1))

	filter := bson.M{}
	if q != "" {
		qInt, _ := strconv.Atoi(q)
		filter = bson.M{"$or": []interface{}{
			bson.M{"name": primitive.Regex{Pattern: "^" + q, Options: ""}},
			bson.M{"user.username": primitive.Regex{Pattern: "^" + q, Options: ""}},
			bson.M{"id": qInt},
			bson.M{"user.id": qInt}},
		}
	}

	resources, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		panic("Something went wrong!")
	}

	returnResources := make([]map[string]interface{}, 0)
	for resources.Next(context.Background()) {
		var resource map[string]interface{}
		if err = resources.Decode(&resource); err != nil {
			panic(err)
		}
		returnResources = append(returnResources, resource)
	}

	return returnResources
}