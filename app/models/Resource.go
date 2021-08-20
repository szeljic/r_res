package models

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"strconv"
	"time"
)



type Resource struct {
	ID int											`bson:"id" json:"id"`
	CreatedAt int64									`bson:"created_at" json:"created_at"`
	Name string										`bson:"name" json:"name"`
	User User										`bson:"user" json:"user"`
}

func SaveResource(data map[string]interface{}, user User, categoryID int) error {

	var category Category
	collection := DB.Database(Database).Collection("categories")
	err := collection.FindOne(context.Background(), bson.M{"id": categoryID}).Decode(&category)
	if err != nil {
		log.Println(err)
		return &errorString{err.Error()}
	}
	err = validateDataByCategory(category, data, "create")

	if err != nil {
		return &errorString{err.Error()}
	}

	data = removeUnneededData(category, data)

	collection = DB.Database(Database).Collection("resources")
	var set bson.D
	for key, value := range data {
		set = append(set, bson.E{Key: key, Value: value})
	}
	user.Password = ""
	category.User.Password = ""
	log.Println(user)
	set = append(set, bson.E{Key: "user", Value: user})
	set = append(set, bson.E{Key: "category", Value: category})
	set = append(set, bson.E{Key: "id", Value: resourceMaxID() + 1})
	set = append(set, bson.E{Key: "created_at", Value: time.Now().Unix()})
	set = append(set, bson.E{Key: "deleted_at", Value: nil})

	_, err = collection.InsertOne(context.Background(), set)
	if err != nil {
		return err
	}

	return nil
}

func resourceMaxID() int {
	collection := DB.Database(Database).Collection("resources")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"id", -1}})
	findOptions.SetLimit(1)

	resources, err := collection.Find(context.Background(), bson.M{}, findOptions)

	if err != nil {
		return 0
	}

	for resources.Next(context.Background()) {
		resource := make(map[string]interface{})
		err = resources.Decode(&resource)
		if err != nil {
			log.Panic(err)
			return 0
		}
		log.Println(resource["id"])
		if resource["id"] != nil {
			return int(resource["id"].(int32))
		}
	}
	return 0
}

func removeUnneededData(category Category, data map[string]interface{}) map[string]interface{} {
	for key, _ := range data {
		if key == "name" || key == "category_id" {
			continue
		}
		exists := false
		for _, v := range category.SpecificFields {
			if key == v.SCName {
				exists = true
			}
		}
		if !exists {
			delete(data, key)
		}
	}
	return data
}

func validateDataByCategory(category Category, data map[string]interface{}, action string) error {

	for _, value := range category.SpecificFields {

		if action != "update" && value.Required && (data[value.SCName] == "" || data[value.SCName] == nil) {
			return &errorString{value.Name + " je obavezno polje!"}
		}

		if _, ok := data[value.SCName]; !ok && action == "update" {
			continue
		}

		log.Println("BORA", value.DataType, value)

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
		case "checkbox":
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

func GetTotalResources(q string, categoryId int) int {
	collection := DB.Database(Database).Collection("resources")

	filter := bson.D{bson.E{
		Key:   "deleted_at",
		Value: nil,
	}}
	if q != "" {
		qInt, _ := strconv.Atoi(q)
		filter = bson.D{{"$or", bson.A{
			bson.D{{"name", primitive.Regex{Pattern: "^" + q, Options: ""}}},
			bson.D{{"user.username", primitive.Regex{Pattern: "^" + q, Options: ""}}},
			bson.D{{"id", qInt}},
			bson.D{{"user.id", qInt}},
		},
		}, {Key: "deleted_at", Value: nil},
		}
	}

	if categoryId > 0 {
		filter = append(filter, bson.E{Key: "category_id", Value: categoryId})
	}

	data, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	return int(data)
}

func GetResources(order, sortBy, q string, paginateBy, page int64, categoryId int) []map[string]interface{} {

	collection := DB.Database(Database).Collection("resources")

	findOptions := options.Find()
	if order == "desc" {
		findOptions.SetSort(bson.D{{sortBy, -1}})
	} else {
		findOptions.SetSort(bson.D{{sortBy, 1}})
	}

	findOptions.SetLimit(paginateBy)
	findOptions.SetSkip(paginateBy * (page - 1))

	filter := bson.D{bson.E{
		Key:   "deleted_at",
		Value: nil,
	}}
	if q != "" {
		qInt, _ := strconv.Atoi(q)
		filter = bson.D{{"$or", bson.A{
					bson.D{{"name", primitive.Regex{Pattern: "^" + q, Options: ""}}},
					bson.D{{"user.username", primitive.Regex{Pattern: "^" + q, Options: ""}}},
					bson.D{{"id", qInt}},
					bson.D{{"user.id", qInt}},
				},
			}, {Key: "deleted_at", Value: nil},
		}
	}

	if categoryId > 0 {
		filter = append(filter, bson.E{Key: "category_id", Value: categoryId})
	}

	log.Println("KOJAJJAJAJA")
	log.Println(filter)

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

func GetResource(id int) map[string]interface{} {

	collection := DB.Database(Database).Collection("resources")

	resource := collection.FindOne(context.Background(), bson.M{"id": id, "deleted_at": nil})

	r := make(map[string]interface{})
	err := resource.Decode(&r)

	if err != nil {
		return nil
	}

	return r
}

func GetTrimResource(id int) *Resource {

	collection := DB.Database(Database).Collection("resources")

	resource := collection.FindOne(context.Background(), bson.M{"id": id, "deleted_at": nil})


	r := make(map[string]interface{})
	err := resource.Decode(&r)

	if err != nil {
		return nil
	}

	res := Resource{}

	user := User{}
	err = mapstructure.Decode(r["user"], &user)

	if err != nil {
		panic(err)
	}

	res.ID = int(r["id"].(int32))
	res.Name = r["name"].(string)
	res.User = user
	if r["create_at"] != nil {
		res.CreatedAt = r["created_at"].(int64)
	}

	return &res
}

func DeleteResource(id int) int64 {

	session, err := DB.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	var modified int64
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := DB.Database(Database).Collection("resources")
		now := time.Now().Unix()
		result, err := collection.UpdateOne(context.Background(),
			bson.M{"id": id},
			bson.D{
				{"$set", bson.M{"deleted_at": now},
				}})

		if err != nil {
			return err
		}

		err = DeleteReservationsOfResource(id)

		if err != nil {
			return err
		}
		modified = result.ModifiedCount
		return nil
	})

	if err != nil {
		return 0
	}

	return modified
}

func UpdateResource(id int, data map[string]interface{}) error {

	r := GetResource(id)

	var category Category
	collection := DB.Database(Database).Collection("categories")
	err := collection.FindOne(context.Background(), bson.M{"id": r["category_id"]}).Decode(&category)
	if err != nil {
		log.Println(err)
		return &errorString{err.Error()}
	}
	err = validateDataByCategory(category, data, "update")
	if err != nil {
		return &errorString{err.Error()}
	}
	data = removeUnneededData(category, data)
	collection = DB.Database(Database).Collection("resources")
	var set bson.D
	for key, value := range data {
		set = append(set, bson.E{Key: key, Value: value})
	}

	if len(data) < 1 {
		return &errorString{"Nema podataka za promjenu!"}
	}

	_, err = collection.UpdateOne(context.Background(),
		bson.M{"id": id},
		bson.D{
			{"$set", set},
		})

	if err != nil {
		return err
	}

	return nil
}

func DeleteResourcesOfCategory(categoryID int) error {
	collection := DB.Database(Database).Collection("resources")
	findOptions := options.Find()
	filter := bson.M{"category.id": categoryID}
	resources, err := collection.Find(context.Background(), filter, findOptions)
	log.Println("D")

	if err != nil {
		return err
	}
	log.Println("E")

	for resources.Next(context.Background()) {
		log.Println("AAS")
		log.Println("AAS")
		var resource Resource
		if err = resources.Decode(&resource); err != nil {
			return err
		}
		log.Println(resource.ID)
		go DeleteResource(resource.ID)
	}

	return nil
}