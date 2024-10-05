package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/inventory/db"
	"github.com/Mayankrai449/ecom-microservice/inventory/models"
	"github.com/labstack/echo/v4"
)

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}
	product.ID = 0

	result := db.GetDB().Model(&models.Product{}).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to update product",
			"message": result.Error.Error(),
		})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Product not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product updated successfully",
		"product": product,
	})
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	result := db.GetDB().Delete(&models.Product{}, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to delete product",
			"message": result.Error.Error(),
		})
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Product not found",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
