package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type User struct {
	ID int
	Username string
	Password string
	FirstName string
	LastName string
	Email string
	DateOfBirth primitive.DateTime
}

func SaveUser(username, password, firstName, lastName, dob, email string) error {

	log.Println(Database)
	var user User
	collection := DB.Database(Database).Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)

	if err == nil {
		log.Println(err)
		return &errorString{"Korisnik vec postoji!"}
	}

	_, err = collection.InsertOne(context.Background(),
		bson.M{
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