package models

import "go.mongodb.org/mongo-driver/mongo"

var (
	DB *mongo.Client
	Database string
)

type errorString struct {
	message string
}
func (e *errorString) Error() string {
	return e.message
}