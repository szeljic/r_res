package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"log"
	"net/http"
	"r_res/app/models"
	"strconv"
)

type Resource struct {
	*revel.Controller
}

type ResponseResources struct {
	Page int						`json:"page"`
	PaginateBy int					`json:"paginateBy"`
	Total int						`json:"total"`
	Items []map[string]interface{}	`json:"items"`
}

func (c Resource) Index() revel.Result {
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

	total := models.GetTotalResources(q)

	resources := models.GetResources(order, sortBy, q, int64(paginateBy), int64(page))

	r := ResponseResources{
		Page:       page,
		PaginateBy: paginateBy,
		Total:      total,
		Items: 		resources,
	}

	if resources == nil {
		emptyCategory := make([]map[string]interface{}, 0)
		r.Items = emptyCategory
	}

	return c.RenderJSON(r)
}
func (c Resource) Show() revel.Result {

	id, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	resource := models.GetResource(id)
	return c.RenderJSON(resource)

}
func (c Resource) Create() revel.Result {

	user := models.GetLoggedUser(c.Request.Header.Get("x-token"))

	if user == (models.User{}) {
		log.Println("USER IS NOT LOGGED IN!!!")
		r := TokenResponse{
			Logged: false,
		}
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(r)
	}

	rq := c.Params.JSON
	var data map[string]interface{}
	_ = json.Unmarshal(rq, &data)

	if data["name"] == "" {
		r := Response{
			Message: "Ime je obavezno polje!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	cID, ok := data["category_id"].(float64)
	if !ok {
		log.Println("AAA", cID)
	}

	categoryID := int(cID)

	if categoryID <= 0 {
		r := Response{
			Message: "Kategorija je obavezno polje!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	err := models.SaveResource(data, user, categoryID)

	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusInternalServerError
		return c.RenderJSON(r)
	}

	r := Response{
		Message: "Success!",
		Code:    200,
	}

	return c.RenderJSON(r)
}
func (c Resource) Update() revel.Result {

	rq := c.Params.JSON
	var data map[string]interface{}
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

	err = models.UpdateResource(id, data)
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
func (c Resource) Delete() revel.Result {

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

	n := models.DeleteResource(id)
	if n > 0 {
		r = Response{
			Message: "Success",
			Code:    200,
		}
	} else {
		c.Response.Status = http.StatusBadRequest
		r = Response{
			Message: "Record not found!",
			Code:    0,
		}
	}
	return c.RenderJSON(r)
}