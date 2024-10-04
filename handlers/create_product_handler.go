package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/db"
	"github.com/Mayankrai449/ecom-microservice/models"
	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}
	result := db.GetDB().Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create product",
			"error":   result.Error.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Product created successfully",
		"product": product,
	})
}
