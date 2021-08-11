package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

type Reservation struct {
	ID int												`bson:"id" json:"id"`
	FromDate int64										`bson:"from_date" json:"from_date"`
	ToDate int64										`bson:"to_date" json:"to_date"`
	CreatedBy int										`bson:"created_by" json:"created_by"`
	CreatedAt int64										`bson:"created_at" json:"created_at"`
	User User											`bson:"user" json:"user"`
	Resource Resource									`bson:"resource" json:"resource"`
}

func GetReservations(order, sortBy, q string, paginateBy, page int64) []Reservation {

	collection := DB.Database(Database).Collection("reservations")

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

	reservations, err := collection.Find(context.Background(), filter, findOptions)

	if err != nil {
		panic("Something went wrong!")
	}
	var returnReservations []Reservation
	for reservations.Next(context.Background()) {
		var reservation Reservation
		if err = reservations.Decode(&reservation); err != nil {
			panic(err)
		}
		returnReservations = append(returnReservations, reservation)
	}

	return returnReservations
}

func GetTotalReservations(q string) int {
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

func SaveReservation(fromDate, toDate int64, resourceID int, createdBy int, createdAt int64) error {

	var user User
	collection := DB.Database(Database).Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"id": createdBy}).Decode(&user)

	if err != nil {
		log.Println(err)
		return &errorString{err.Error()}
	}

	resource := GetTrimResource(resourceID)

	collection = DB.Database(Database).Collection("reservations")
	user.Password = ""
	_, err = collection.InsertOne(context.Background(),
		bson.M{
			"id": 			reservationMaxID() + 1,
			"from_date": 	fromDate,
			"to_date": 		toDate,
			"created_by": 	createdBy,
			"created_at": 	createdAt,
			"user": 		user,
			"resource": 	resource,
		})

	if err != nil {
		return err
	}

	return nil
}

func reservationMaxID() int {
	collection := DB.Database(Database).Collection("reservations")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"id", -1}})
	findOptions.SetLimit(1)

	reservations, err := collection.Find(context.Background(), bson.M{}, findOptions)

	if err != nil {
		return 0
	}

	for reservations.Next(context.Background()) {
		var reservation Reservation
		if err = reservations.Decode(&reservation); err != nil {
			log.Panic(err)
			return 0
		}
		return reservation.ID
	}
	return 0
}

func GetReservation(id int) Reservation {

	collection := DB.Database(Database).Collection("reservations")

	reservation := collection.FindOne(context.Background(), bson.M{"id": id})

	var r Reservation
	_ = reservation.Decode(&r)

	log.Println(r)

	return r
}

func UpdateReservation(id int, data map[string]interface{}) error {

	collection := DB.Database(Database).Collection("reservations")
	var set bson.D

	for key, value := range data {
		switch value.(type) {
		case string:
			v, err := strconv.Atoi(value.(string))
			if err != nil {
				set = append(set, bson.E{Key: key, Value: value})
			} else {
				set = append(set, bson.E{Key: key, Value: v})
			}
			break
		case int64:
			set = append(set, bson.E{Key: key, Value: value})
		default:
			return &errorString{"Podaci nisu odgovarajuci!"}
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

func DeleteReservation(id int) int64 {
	collection := DB.Database(Database).Collection("reservations")
	result, err := collection.DeleteOne(context.Background(), bson.M{"id": id})

	if err != nil {
		return 0
	}

	return result.DeletedCount
}