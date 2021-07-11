package routes

import (
	midL "github.com/hadipranoto/stock_opname/middleware"
	"github.com/labstack/echo"
)

func (h *Handler) GetSummaryOpname (c echo.Context) error {	
	var (
		data interface{}
		err error
	)
	
	data, err = h.services.SummaryOpname()

	if err != nil {
		return h.Response.SendBadRequest(c, err.Error(), data)
	}
	
	return h.Response.SendSuccess(c, "OK", data)	
}

func (h *Handler) AddInventoryReport (c echo.Context) error {		
	var (
		data interface{}
	)
	if midL.ReqReport.Kind == "kind_initial_inventory" {
		return h.Response.SendSuccess(c, "OK", data)
	}
	if midL.ReqReport.Kind == "kind_sale" {
		return h.Response.SendSuccess(c, "OK", data)
	}

	if midL.ReqReport.Kind == "kind_incoming_product" {
		return h.Response.SendSuccess(c, "OK", data)
	}

	if midL.ReqReport.Kind == "kind_final_inventory" {
		return h.Response.SendSuccess(c, "OK", data)
	}

	if midL.ReqReport.Kind == "kind_warehouse_inventory" {
		return h.Response.SendSuccess(c, "OK", data)
	}	
	
	return h.Response.SendBadRequest(c, "Error", "Kind is not defined")
}


