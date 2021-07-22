package controllers

import (
	"github.com/revel/revel"
	"log"
)

type User struct {
	*revel.Controller
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (c User) Home() revel.Result{

	log.Println("HOME")


	return c.RenderJSON(true)
}