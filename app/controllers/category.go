package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"log"
	"net/http"
	"r_res/app/models"
	"strconv"
	"time"
)

type Category struct {
	*revel.Controller
}

type ResponseCategories struct {
	Page int						`json:"page"`
	PaginateBy int					`json:"paginateBy"`
	Total int						`json:"total"`
	Items *[]models.Category		`json:"items"`
}



func (c Category) Index() revel.Result {

	q := c.Params.Query.Get("q")
	paginateBy, err := strconv.Atoi(c.Params.Query.Get("paginate-by"))
	if err != nil {
		paginateBy = 20
	}
	page, err := strconv.Atoi(c.Params.Query.Get("page"))
	if err != nil {
		page = 1
	}
	sortBy := c.Params.Query.Get("sort-by")
	if sortBy == "" {
		sortBy = "id"
	}
	order := c.Params.Query.Get("order")
	if order == "" {
		order = "asc"
	}

	total := models.GetTotalCategories(q)

	categories := models.GetCategories(order, sortBy, q, int64(paginateBy), int64(page))


	r := ResponseCategories{
		Page:       page,
		PaginateBy: paginateBy,
		Total:      total,
		Items: 		&categories,
	}

	if categories == nil {
		emptyCategory := make([]models.Category, 0)
		r.Items = &emptyCategory
	}

	return c.RenderJSON(r)
}

type CreateStruct struct {
	Name		string `json:"name"`
	Description string `json:"description"`
}

func (c Category) Create() revel.Result {

	user := models.GetLoggedUser(c.Request.Header.Get("x-token"))

	if user == (models.User{}) {
		log.Println("USER IS NOT LOGGED IN!!!")
		r := TokenResponse{
			Logged: false,
		}
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(r)
	}

	var createStruct CreateStruct
	err := c.Params.BindJSON(&createStruct)

	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.RenderJSON(r)
	}

	if createStruct.Name == "" {
		r := Response{
			Message: "Ime je obavezno polje!",
			Code:    0,
		}
		c.RenderJSON(r)
	}

	if createStruct.Description == "" {
		r := Response{
			Message: "Opis je obavezno polje!",
			Code:    0,
		}
		c.RenderJSON(r)
	}
	createdAt := time.Now()
	err = models.SaveCategory(createStruct.Name, createStruct.Description, user.ID, createdAt)

	return c.RenderJSON(true)
}

func (c Category) Update() revel.Result {

	rq := c.Params.JSON
	var data map[string]string
	_ = json.Unmarshal(rq, &data)

	if len(data) == 0 {
		r := Response{
			Message: "Empty data",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	id, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	err = models.UpdateCategory(id, data)
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	return c.RenderJSON(data)
}

func (c Category) Show() revel.Result {

	id, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	category := models.GetCategory(id)
	return c.RenderJSON(category)
}

func (c Category) Delete() revel.Result {

	var r Response

	id, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		r = Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	n := models.DeleteCategory(id)
	if n > 0 {
		r = Response{
			Message: "Success",
			Code:    200,
		}
	} else {
		r = Response{
			Message: "Record not found!",
			Code:    0,
		}
	}
	return c.RenderJSON(r)
}