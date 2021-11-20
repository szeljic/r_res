package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"net/http"
	"r_res/app/common"
	"r_res/app/models"
	"time"
)

type Auth struct {
	*revel.Controller
}

type LoginResponse struct {
	Message 	string `json:"message"`
	Code    	int    `json:"code"`
	AccessToken string `json:"access_token"`
}

type RegistrationStruct struct {
	Username       string                    `json:"username"`
	Password       string                    `json:"password"`
	FirstName      string                    `json:"first_name"`
	LastName       string                    `json:"last_name"`
	DateOfBirth    string                    `json:"date_of_birth"`
	Email          string                    `json:"email"`
}

func (c Auth) Registration() revel.Result {

	var registrationStruct RegistrationStruct
	err := c.Params.BindJSON(&registrationStruct)
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    -1,
		}
		return c.RenderJSON(r)
	}

	err = models.SaveUser(registrationStruct.Username, registrationStruct.Password,
		registrationStruct.FirstName, registrationStruct.LastName,
		registrationStruct.DateOfBirth, registrationStruct.Email)

	if err != nil {
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

func (c Auth) Login() revel.Result {

	var loginStruct LoginStruct
	err := c.Params.BindJSON(&loginStruct)
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    -1,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	userID := models.CheckCredentials(loginStruct.Username, loginStruct.Password)
	if userID > 0 {

		expirationTime := time.Now().Add(1500 * time.Minute)

		claims := &common.Claims{
			Username: loginStruct.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := t.SignedString(common.JwtKey)
		if err != nil {
			r := LoginResponse {
				Message: "false",
				Code: http.StatusUnprocessableEntity,
				AccessToken: "",
			}
			c.Response.Status = http.StatusUnprocessableEntity
			return c.RenderJSON(r)
		}

		r := LoginResponse {
			Message: "success",
			Code: 200,
			AccessToken: tokenString,
		}
		c.Response.Status = http.StatusOK
		return c.RenderJSON(r)
	}

	r := Response{
		Message: "Korisničko ime ili lozinka nisu odgovarajući!",
		Code: 403,
	}
	c.Response.Status = http.StatusBadRequest
	return c.RenderJSON(r)
}
type TokenResponse struct {
	Success bool	`json:"success"`
	Logged bool		`json:"logged"`
}

func (c Auth) TokenValidation() revel.Result {

	token := c.Request.Header.Get("x-token")
	claims := &common.Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return common.JwtKey, nil
	})

	if err != nil {
		r := TokenResponse{
			Success: false,
			Logged: false,
		}
		if err == jwt.ErrSignatureInvalid {
			c.Response.Status = http.StatusUnauthorized
			return c.RenderJSON(r)
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	if !tkn.Valid {
		r := TokenResponse{
			Success: false,
			Logged: false,
		}
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(r)
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 120*time.Second {
		r := TokenResponse{
			Success: false,
			Logged: true,
		}
		c.Response.Status = http.StatusOK
		return c.RenderJSON(r)
	}

	c.Response.Status = http.StatusOK
	r := TokenResponse{
		Success: true,
		Logged: true,
	}
	return c.RenderJSON(r)
}

type ActiveResponse struct {
	UserID int			`json:"id"`
	Username string		`json:"username"`
	FirstName string	`json:"first_name"`
	LastName string		`json:"last_name"`
	UserType string 	`json:"user_type"`
}

func (c Auth) ActiveUser() revel.Result {
	user := models.GetLoggedUser(c.Request.Header.Get("x-token"))

	if user == (models.User{}) {
		r := TokenResponse{
			Logged: false,
		}
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(r)
	}

	res := ActiveResponse{
		UserID: user.ID,
		Username: user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserType:  "user",
	}
	if user.IsAdmin == true {
		res.UserType = "admin"
	}

	return c.RenderJSON(res)
}