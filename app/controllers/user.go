package controllers

import (
	"github.com/revel/revel"
	"log"
	"net/http"
	"r_res/app/models"
)

type User struct {
	*revel.Controller
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type RegistrationStruct struct {
	Username       string                    `json:"username"`
	Password       string                    `json:"password"`
	FirstName      string                    `json:"first_name"`
	LastName       string                    `json:"last_name"`
	DateOfBirth    string                    `json:"date_of_birth"`
	Email          string                    `json:"email"`
}

func (c User) Registration() revel.Result {

	var registrationStruct RegistrationStruct
	c.Params.BindJSON(&registrationStruct)

	err := models.SaveUser(registrationStruct.Username, registrationStruct.Password,
		registrationStruct.FirstName, registrationStruct.LastName,
		registrationStruct.DateOfBirth, registrationStruct.Email)

	if err != nil {
		log.Println(err)
		r := Response{
			Message: err.Error(),
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

type LoginStruct struct {
	Username    string                    `json:"username"`
	Password    string                    `json:"password"`
}

func (c User) Login() revel.Result {

	var loginStruct LoginStruct
	c.Params.BindJSON(&loginStruct)

	var r Response

	userId := models.CheckCredentials(loginStruct.Username, loginStruct.Password)
	log.Println(userId, "KRATAK")
	if userId > 0 {

		c.Session["user"] = loginStruct.Username
		c.Session.SetDefaultExpiration()

		token, err := models.CreateToken(userId)

		if err != nil {
			return c.RenderJSON(http.StatusUnprocessableEntity)
		}
		c.RenderJSON(token)
		log.Println(token, err, "AAA")

		c.Response.Status = http.StatusOK
		return c.RenderJSON(token)
	}

	r = Response{
		Message: "Korisnicko ime ili lozinka nisu odgovarajuci!",
		Code: 403,
	}

	return c.RenderJSON(r)
}