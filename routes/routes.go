package routes

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	midL "github.com/hadipranoto/stock_opname/middleware"
	"github.com/hadipranoto/stock_opname/services"
	"github.com/labstack/echo"
)



type Handler struct {
	Response    midL.HTTPResponseHelper		
	services	services.OpnameServiceInterface
}

func (h *Handler) HelloRoot (c echo.Context) error {		
	return h.Response.SendSuccess(c, "OK", "Hello from the web side.")	
}


func RegisterRoutes (e *echo.Echo)  {

	//conn sql 
	sqlConn, err := SqlConnect()
	
	if err != nil {
		panic(fmt.Sprintf("Sql failed to load %v",err.Error()))
	}

	serv := services.NewOpnameService(sqlConn)	
	handler := Handler{services:serv}


	midle := midL.InitMiddleware()
	
	e.GET("/", handler.HelloRoot)

	//showing summaries opname
	e.GET("/summary-opname", handler.GetSummaryOpname)

	//adding product to tables 		
	e.POST("/product", handler.AddProduct, midle.ProductMiddleware)
	
	//api for inserting report data from different kind of documents
	e.POST("/initial-inventory", handler.AddInventoryReport, midle.ProductMiddleware)
	e.POST("/sale", handler.AddInventoryReport, midle.ProductMiddleware)
	e.POST("/incoming-product", handler.AddInventoryReport, midle.ProductMiddleware)
	e.POST("/final-inventory", handler.AddInventoryReport, midle.ProductMiddleware)
	e.POST("/warehouse-inventory", handler.AddInventoryReport, midle.ProductMiddleware)

}



func SqlConnect() (*sql.DB, error) {
	user := "root"
	password := "tarosabi"
	host := "127.0.0.1"
	port := "3306"
	dbName := "stock_opname"

	sqlSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, password,host,port,dbName)

	db, err := sql.Open("mysql", sqlSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}


