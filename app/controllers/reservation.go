package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"net/http"
	"r_res/app/models"
	"strconv"
	"time"
)

type Reservation struct {
	*revel.Controller
}

type ResponseReservations struct {
	Page int						`json:"page"`
	PaginateBy int					`json:"paginateBy"`
	Total int						`json:"total"`
	Items *[]models.Reservation		`json:"items"`
}

func (c Reservation) Index() revel.Result {

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

	total := models.GetTotalReservations(q)
	reservations := models.GetReservations(order, sortBy, q, int64(paginateBy), int64(page))
	r := ResponseReservations{
		Page:       page,
		PaginateBy: paginateBy,
		Total:      total,
		Items: 		&reservations,
	}

	if reservations == nil {
		emptyReservation := make([]models.Reservation, 0)
		r.Items = &emptyReservation
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(r)
}

func (c Reservation) Show() revel.Result {
	id, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	reservation := models.GetReservation(id)
	c.Response.Status = http.StatusOK
	if reservation == nil {
		return c.RenderJSON(make(map[string]string))
	}
	return c.RenderJSON(&reservation)
}

type CreateReservation struct {
	FromDate				string 						`json:"from_date"`
	ToDate 					string 						`json:"to_date"`
	ResourceID				int							`json:"resource_id"`
}

func (c Reservation) Create() revel.Result {

	user := models.GetLoggedUser(c.Request.Header.Get("x-token"))

	if user == (models.User{}) {
		r := TokenResponse{
			Logged: false,
		}
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(r)
	}

	var createReservation CreateReservation
	err := c.Params.BindJSON(&createReservation)

	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	if createReservation.FromDate == "" {
		r := Response{
			Message: "Datum je obavezno polje!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	if createReservation.ToDate == "" {
		r := Response{
			Message: "Datum je obavezno polje!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	if createReservation.ResourceID <= 0 {
		r := Response{
			Message: "Resurs je obavezno polje!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	layout := "2006-01-02 15:04:05"
	fromDate, err := time.Parse(layout, createReservation.FromDate)
	if err != nil {
		r := Response{
			Message: "Format datuma nije validan!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	toDate, err := time.Parse(layout, createReservation.ToDate)
	if err != nil {
		r := Response{
			Message: "Format datuma nije validan!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	if fromDate.Unix() >= toDate.Unix() {
		r := Response{
			Message: "Datumi nisu u redu!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	if !models.IsResourceAvailable(fromDate.Unix(), toDate.Unix(), createReservation.ResourceID) {
		r := Response{
			Message: "Resurs je zauzet u datom periodu!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	createdAt := time.Now().Unix()
	err = models.SaveReservation(fromDate.Unix(), toDate.Unix(), createReservation.ResourceID, user.ID, createdAt)

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

func (c Reservation) Update() revel.Result {
	rq := c.Params.JSON
	var data map[string]interface{}
	_ = json.Unmarshal(rq, &data)

	id, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	_, fromDateOK := data["from_date"]
	_, toDateOK := data["to_date"]

	if fromDateOK != toDateOK {
		r := Response{
			Message: "Molimo posaljite oba datuma!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}
	layout := "2006-01-02 15:04:05"
	if val, ok := data["from_date"]; ok {
		fromDate, err := time.Parse(layout, val.(string))
		if err != nil {
			r := Response{
				Message: "Format datuma nije validan!",
				Code:    0,
			}
			c.Response.Status = http.StatusBadRequest
			return c.RenderJSON(r)
		}
		data["from_date"] = fromDate.Unix()
	}
	if val, ok := data["to_date"]; ok {
		toDate, err := time.Parse(layout, val.(string))
		if err != nil {
			r := Response{
				Message: "Format datuma nije validan!",
				Code:    0,
			}
			c.Response.Status = http.StatusBadRequest
			return c.RenderJSON(r)
		}
		data["to_date"] = toDate.Unix()
	}

	if fromDateOK && toDateOK && data["from_date"].(int64) >= data["to_date"].(int64) {
		r := Response{
			Message: "Datumi nisu u redu!",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)

	}

	if _, ok := data["user"]; ok {
		delete(data, "user")
	}

	if _, ok := data["created_by"]; ok {
		delete(data, "created_by")
	}

	if _, ok := data["resource"]; ok {
		delete(data, "resource")
	}

	if _, ok := data["id"]; ok {
		delete(data, "id")
	}

	if _, ok := data["_id"]; ok {
		delete(data, "_id")
	}

	if _, ok := data["created_at"]; ok {
		delete(data, "created_at")
	}

	if _, ok := data["resource_id"]; ok {
		delete(data, "resource_id")
	}

	if len(data) == 0 {
		r := Response{
			Message: "Empty data",
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	err = models.UpdateReservation(id, data)
	if err != nil {
		r := Response{
			Message: err.Error(),
			Code:    0,
		}
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(r)
	}

	c.Response.Status = http.StatusOK
	r := Response{
		Message: "Success!",
		Code:    200,
	}
	return c.RenderJSON(r)
}

func (c Reservation) Delete() revel.Result {
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

	n := models.DeleteReservation(id)
	if n > 0 {
		c.Response.Status = http.StatusOK
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