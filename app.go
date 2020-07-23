package main

import (
	"coredns-helper/handlers"
	"coredns-helper/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := utils.InitDB("dns.db")

	e := echo.New()
  	e.Use(middleware.Logger())

	e.GET("/domains", handlers.GetDomains(db))
	e.POST("/domains", handlers.PutDomains(db))
	e.PUT("/domains", handlers.EditDomain(db))
	e.DELETE("/domains/:id", handlers.DeleteDomain(db))

	e.Logger.Fatal(e.Start(":14045"))
}


