package controllers

import (
	"context"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"r_res/app/models"
)

type User struct {
	*revel.Controller
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (c User) Registration() revel.Result {
	e := models.DB.Ping(context.TODO(), readpref.Primary())

	if e != nil {
		panic(e)
	}

	username := c.Params.Get("username")
	password := c.Params.Get("password")
	firstName := c.Params.Get("first_name")
	lastName := c.Params.Get("last_name")
	dob := c.Params.Get("date_of_birth")
	email := c.Params.Get("email")

	err := models.SaveUser(username, password, firstName, lastName, dob, email)

	if err != nil {
		r := Response{
			Message: "failed",
			Code:    -1,
		}
		return c.RenderJSON(r)
	}

	r := Response{
		Message: "success",
		Code:    200,
	}
	return c.RenderJSON(r)
}