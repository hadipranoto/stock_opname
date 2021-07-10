package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)


func main (){
	fmt.Println("Welcome to the server")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {return c.String(http.StatusOK, "Hello from the web side.")})

	e.Start(":8000")

}