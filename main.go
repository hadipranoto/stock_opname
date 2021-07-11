package main

import (
	"fmt"

	"github.com/hadipranoto/stock_opname/routes"
	"github.com/labstack/echo"
)


func main (){
	fmt.Println("Welcome to the server")

	e := echo.New()
	

	routes.RegisterRoutes(e)

	e.Start(":8000")

}