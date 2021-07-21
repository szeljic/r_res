package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type TokenResponse struct {
	Success bool	`json:"success"`
	Logged bool		`json:"logged"`
}

func (c App) TokenValidation() revel.Result {
	r := TokenResponse{
		Success: true,
		Logged: true,
	}
	return c.RenderJSON(r)
}