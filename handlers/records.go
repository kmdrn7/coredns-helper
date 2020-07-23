package handlers

import (
	"database/sql"
	"net/http"
	"coredns-helper/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetRecords(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetRecords(db))
	}
}

func PutRecords(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var record models.Record
		c.Bind(&record)
		if (record.Name == "" || record.Type == "" || record.Content == ""){
			return c.JSON(http.StatusBadRequest, H{
				"message": "Ada data yang kosong",
			})
		}
		id, err := models.PutRecords(db, record.DomainID, record.Name, record.Type, record.Content, record.Ttl, record.Disabled)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

func EditRecords(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var record models.Record
		c.Bind(&record)
		_, err := models.EditRecords(db, record.ID, record.Name, record.Type, record.Content, record.Ttl, record.Disabled)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": record,
			})
		} else {
			return err
		}
	}
}

func DeleteRecords(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteRecords(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
