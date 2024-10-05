package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/inventory/db"
	"github.com/Mayankrai449/ecom-microservice/inventory/models"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	var products []models.Product
	result := db.GetDB().Find(&products)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to retrieve products",
			"message": result.Error.Error(),
		})
	}

	if len(products) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "No products found",
			"products": products,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Products retrieved successfully",
		"products": products,
	})
}

func GetProductByID(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	result := db.GetDB().First(&product, id)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error":   "Product not found",
				"message": "No product found with the specified ID",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to retrieve product",
			"message": result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product retrieved successfully",
		"product": product,
	})
}
