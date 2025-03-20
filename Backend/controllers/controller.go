package controllers

import (
	"Backend/connection"
	"Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	rows, err := connection.DB.Query(c, "SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time FROM item")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la base de datos"})
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Ticker, &item.Company, &item.Brokerage, &item.Action, &item.RatingFrom, &item.RatingTo, &item.TargetFrom, &item.TargetTo, &item.Time)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear los datos"})
			return
		}
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}
