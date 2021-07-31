package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

type Category struct {
	ID int											`bson:"id" json:"id"`
	Name string										`bson:"name" json:"name"`
	Description string								`bson:"description" json:"description"`
	CreatedBy int									`bson:"created_by" json:"created_by"`
	CreatedAt time.Time								`bson:"created_at" json:"created_at"`
	User User										`bson:"user" json:"user"`
	SpecificFields []SpecificField 					`bson:"specific_fields" json:"specific_fields"`
}

type SpecificField struct {
	Name 		string		`bson:"name" json:"name"`
	Required 	bool		`bson:"required" json:"required"`
	DataType	string		`bson:"data_type" json:"data_type"`
}

func SaveCategory(name, description string, specificFields []SpecificField, createdBy int, createdAt time.Time) error {

	var user User
	collection := DB.Database(Database).Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"id": createdBy}).Decode(&user)

	if err != nil {
		log.Println(err)
		return &errorString{err.Error()}
	}

	collection = DB.Database(Database).Collection("categories")
	_, err = collection.InsertOne(context.Background(),
		bson.M{
			"id": categoryMaxID() + 1,
			"name": name,
			"description": description,
			"created_by": createdBy,
			"created_at": createdAt,
			"user": user,
			"specific_fields":specificFields,
		})

	if err != nil {
		return err
	}

	return nil
}

func categoryMaxID() int {
	collection := DB.Database(Database).Collection("categories")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"id", -1}})
	findOptions.SetLimit(1)

	categories, err := collection.Find(context.Background(), bson.M{}, findOptions)

	if err != nil {
		return 0
	}

	for categories.Next(context.Background()) {
		var category Category
		if err = categories.Decode(&category); err != nil {
			log.Panic(err)
		}
		log.Println(category.Name)
		return category.ID
	}
	return 0
}

func GetCategories(order, sortBy, q string, paginateBy, page int64) []Category {

	collection := DB.Database(Database).Collection("categories")

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

	users, err := collection.Find(context.Background(), filter, findOptions)

	if err != nil {
		panic("Something went wrong!")
	}
	var returnCategories []Category
	for users.Next(context.Background()) {
		var category Category
		if err = users.Decode(&category); err != nil {
			panic(err)
		}
		returnCategories = append(returnCategories, category)
	}

	return returnCategories
}

func GetTotalCategories(q string) int {
	collection := DB.Database(Database).Collection("categories")

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

func GetCategory(id int) Category {

	collection := DB.Database(Database).Collection("categories")

	category := collection.FindOne(context.Background(), bson.M{"id": id})

	var c Category
	_ = category.Decode(&c)

	log.Println(c)

	return c
}

func UpdateCategory(id int, data map[string]string) error {

	collection := DB.Database(Database).Collection("categories")
	var set bson.D

	for key, value := range data {
		v, err := strconv.Atoi(value)
		if err != nil {
			set = append(set, bson.E{Key: key, Value: value})
		} else {
			set = append(set, bson.E{Key: key, Value: v})
		}
	}

	_, err := collection.UpdateOne(context.Background(),
		bson.M{"id": id},
		bson.D{
			{"$set", set},
		})

	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id int) int64 {
	collection := DB.Database(Database).Collection("categories")
	result, err := collection.DeleteOne(context.Background(), bson.M{"id": id})

	if err != nil {
		return 0
	}

	return result.DeletedCount
}