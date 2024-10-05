package handlers

import (
	"net/http"

	"github.com/Mayankrai449/ecom-microservice/users/db"
	"github.com/Mayankrai449/ecom-microservice/users/models"
	"github.com/Mayankrai449/ecom-microservice/users/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Hash the Password")
	}
	user.Password = string(hashedPassword)

	if err := db.GetDB().Create(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "User already exists!")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Generate Token")
	}

	user.Password = ""
	return c.JSON(http.StatusCreated, models.AuthResponse{Token: token, User: *user})
}

func Login(c echo.Context) error {
	req := new(models.LoginRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := db.GetDB().Where("email = ?", req.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Credentials! Email not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Credentials! Wrong Password")
	}

	token, err := utils.GenerateJWT(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Generate Token")
	}

	user.Password = ""
	return c.JSON(http.StatusOK, models.AuthResponse{Token: token, User: user})
}
