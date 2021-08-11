package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"log"
	"net/http"
	"r_res/app/models"
	"strconv"
)

type User struct {
	*revel.Controller
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ResponseUsers struct {
	Page int				`json:"page"`
	PaginateBy int			`json:"paginateBy"`
	Total int				`json:"total"`
	Items *[]models.User	`json:"items"`
}
func (c User) Index() revel.Result{
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

	total := models.GetTotal(q)
	log.Println(total)

	users := models.GetUsers(q, sortBy, order, int64(paginateBy), int64(page))

	r := ResponseUsers{
		Page: 		page,
		PaginateBy: paginateBy,
		Items: 		&users,
		Total: 		total,
	}

	if users == nil {
		emptyUser := make([]models.User, 0)
		r.Items = &emptyUser
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(r)
}



func (c User) Update() revel.Result {

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

	err = models.UpdateUser(id, data)
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}

		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}
	r := Response{
		Message: "Success",
		Code:    200,
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(r)
}

func (c User) Show() revel.Result {

	id, err := strconv.Atoi(c.Params.Route.Get("id"))

	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}

		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	user := models.GetUser(id)

	if user == nil {
		return c.RenderJSON(make(map[string]string))
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(&user)
}

