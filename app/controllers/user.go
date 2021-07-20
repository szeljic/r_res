package controllers

import "github.com/revel/revel"

type User struct {
	*revel.Controller
}

type Response struct {
	message string
	code int
}

func (c User) Registration() revel.Result {

	r := Response{
		message: "success",
		code: 200,
	}

	return c.RenderJSON(r)
}