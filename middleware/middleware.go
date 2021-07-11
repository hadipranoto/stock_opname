package middleware

import (
	"github.com/labstack/echo"
)



type GoMiddleware struct {
	response HTTPResponseHelper
}

func InitMiddleware() *GoMiddleware{
	return &GoMiddleware{}
}

func (m *GoMiddleware) ProductMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		err          error
		response     HTTPResponseHelper		
	)

	return func(c echo.Context) error {		

		if err = c.Bind(&ReqProduct); err != nil {
			return response.SendBadRequest(c, err.Error(), nil)
		}

		if ReqProduct.ProductCode == "" {
			return response.SendBadRequest(c, "Product code is empty", nil)
		}
		
		return next(c)
	}
}

func (m *GoMiddleware) InitialInventoryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		err          error
		response     HTTPResponseHelper		
	)

	return func(c echo.Context) error {

		if err = c.Bind(&ReqReport); err != nil {
			return response.SendBadRequest(c, err.Error(), nil)
		}

		if ReqReport.ProductCode == "" {
			return response.SendBadRequest(c, "Product code is empty", nil)
		}
		
		ReqReport.Kind = "kind_initial_inventory"		
		
		return next(c)
	}
}
func (m *GoMiddleware) SaleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		err          error
		response     HTTPResponseHelper		
	)

	return func(c echo.Context) error {

		if err = c.Bind(&ReqReport); err != nil {
			return response.SendBadRequest(c, err.Error(), nil)
		}

		if ReqReport.ProductCode == "" {
			return response.SendBadRequest(c, "Product code is empty", nil)
		}
		ReqReport.Kind = "kind_sale"
		
		return next(c)
	}
}

func (m *GoMiddleware) IncomingProductMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		err          error
		response     HTTPResponseHelper		
	)

	return func(c echo.Context) error {		

		if err = c.Bind(&ReqReport); err != nil {
			return response.SendBadRequest(c, err.Error(), nil)
		}

		if ReqReport.ProductCode == "" {
			return response.SendBadRequest(c, "Product code is empty", nil)
		}
		ReqReport.Kind = "kind_incoming_product"
		return next(c)
	}
}


func (m *GoMiddleware) FinalInventoryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		err          error
		response     HTTPResponseHelper		
	)

	return func(c echo.Context) error {		

		if err = c.Bind(&ReqReport); err != nil {
			return response.SendBadRequest(c, err.Error(), nil)
		}

		if ReqReport.ProductCode == "" {
			return response.SendBadRequest(c, "Product code is empty", nil)
		}
		
		ReqReport.Kind = "kind_final_inventory"

		return next(c)
	}
}


func (m *GoMiddleware) WarehouseInventoryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		err          error
		response     HTTPResponseHelper		
	)

	return func(c echo.Context) error {		

		if err = c.Bind(&ReqReport); err != nil {
			return response.SendBadRequest(c, err.Error(), nil)
		}

		if ReqReport.ProductCode == "" {
			return response.SendBadRequest(c, "Product code is empty", nil)
		}
		
		ReqReport.Kind = "kind_warehouse_inventory"

		return next(c)
	}
}
