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

	e.GET("/records", handlers.GetRecords(db))
	e.POST("/records", handlers.PutRecords(db))
	e.PUT("/records", handlers.EditRecords(db))
	e.DELETE("/records/:id", handlers.DeleteRecords(db))

	e.Logger.Fatal(e.Start(":14045"))
}


