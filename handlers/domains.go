package handlers

import (
	"database/sql"
	"net/http"
	"coredns-helper/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDomains(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetDomains(db))
	}
}

func PutDomains(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var domain models.Domain
		c.Bind(&domain)
		if (domain.Name == "" || domain.Type == ""){
			return c.JSON(http.StatusBadRequest, H{
				"message": "Ada data yang kosong",
			})
		}
		id, err := models.PutDomain(db, domain.Name, domain.Type)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

func EditDomain(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var domain models.Domain
		c.Bind(&domain)
		_, err := models.EditDomain(db, domain.ID, domain.Name, domain.Type)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": domain,
			})
		} else {
			return err
		}
	}
}

func DeleteDomain(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteDomain(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
