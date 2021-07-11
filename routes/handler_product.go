package routes

import (
	"github.com/labstack/echo"
)


func (h *Handler) AddProduct(c echo.Context) error {
	var (
		data interface{}
	)
	return h.Response.SendSuccess(c, "OK", data)
}