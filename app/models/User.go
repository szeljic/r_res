package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID int				`bson:"id"`
	Username string		`bson:"username"`
	Password string		`bson:"password"`
	FirstName string	`bson:"first_name"`
	LastName string		`bson:"last_name"`
	Email string		`bson:"email"`
	DateOfBirth primitive.DateTime
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
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)

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
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"id", -1}})
	findOptions.SetLimit(1)

	users, err := collection.Find(context.Background(), bson.M{}, findOptions)

	if err != nil {
		log.Println(err, "ASD")
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

func GetUsers() []User {

	collection := DB.Database(Database).Collection("users")

	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"id", 1}})

	users, err := collection.Find(context.Background(), bson.M{}, findOptions)

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