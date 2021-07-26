package controllers

import (
	"github.com/revel/revel"
	"r_res/app/models"
)

type User struct {
	*revel.Controller
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (c User) Index() revel.Result{
	users := models.GetUsers()
	return c.RenderJSON(users)
}