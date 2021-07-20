package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type User struct {
	ID int
	Username string
	Password string
	FirstName string
	LastName string
}

func SaveUser(username, password, firstName, lastName, dob, email string) error {

	log.Println(Database)
	collection := DB.Database(Database).Collection("users")

	_, err := collection.InsertOne(context.Background(),
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