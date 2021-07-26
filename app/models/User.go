package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/url"
	"strconv"
)

type User struct {
	ID int				`bson:"id"`
	Username string		`bson:"username"`
	Password string		`bson:"password"`
	FirstName string	`bson:"first_name"`
	LastName string		`bson:"last_name"`
	Email string		`bson:"email"`
	DateOfBirth string	`bson:"date_of_birth"`
	Status int			`bson:"status"`
}

func SaveUser(username, password, firstName, lastName, dob, email string) error {

	password, err := HashPassword(password)
	if err != nil {
		return &errorString{"Doslo je do greske na hashiranju sifre!"}
	}

	var user User
	collection := DB.Database(Database).Collection("users")
	err = collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)

	if err == nil {
		log.Println(err)
		return &errorString{"Korisnik vec postoji!"}
	}

	_, err = collection.InsertOne(context.Background(),
		bson.M{
			"id": maxId() + 1,
			"username": username,
			"password": password,
			"first_name": firstName,
			"last_name": lastName,
			"date_of_birth": dob,
			"email": email,
			"status": 1,
		})

	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckCredentials(username, password string) int {

	var user User
	collection := DB.Database(Database).Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"username": username, "status": 1}).Decode(&user)

	if err != nil {
		log.Println("Korisnik ne postoji")
		return 0
	}

	if !CheckPasswordHash(password, user.Password) {
		return 0
	}

	return user.ID
}

func maxId() int {

	collection := DB.Database(Database).Collection("users")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"id", -1}})
	findOptions.SetLimit(1)

	users, err := collection.Find(context.Background(), bson.M{}, findOptions)

	if err != nil {
		return 0
	}

	for users.Next(context.Background()) {
		var user User
		if err = users.Decode(&user); err != nil {
			log.Panic(err)
		}
		log.Println(user.LastName)
		return user.ID
	}
	return 0
}

func GetUsers(q, sortBy, order string, paginateBy, page int64) []User {

	collection := DB.Database(Database).Collection("users")

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
			bson.M{"username": primitive.Regex{Pattern: "^" + q, Options: ""}},
			bson.M{"first_name": primitive.Regex{Pattern: "^" + q, Options: ""}},
			bson.M{"last_name": primitive.Regex{Pattern: "^" + q, Options: ""}},
			bson.M{"email": primitive.Regex{Pattern: "^" + q, Options: ""}},
			bson.M{"id": qInt}},
		}
	}

	users, err := collection.Find(context.Background(), filter, findOptions)

	if err != nil {
		panic("Something went wrong!")
	}
	var returnUsers []User
	for users.Next(context.Background()) {
		var user User
		if err = users.Decode(&user); err != nil {
			panic(err)
		}
		returnUsers = append(returnUsers, user)
	}

	return returnUsers
}

func GetTotal() int {
	collection := DB.Database(Database).Collection("users")
	filter := bson.M{}
	data, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	return int(data)
}

func UpdateUser(id int, data url.Values) error {

	collection := DB.Database(Database).Collection("users")

	var set bson.D

	for key, value := range data {
		v, err := strconv.Atoi(value[0])
		if err != nil {
			set = append(set, bson.E{Key: key, Value: value[0]})
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