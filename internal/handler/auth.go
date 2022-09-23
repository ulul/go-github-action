package handler

import (
	"hexagonal-architecture/internal/app/user"
	"hexagonal-architecture/internal/middleware"
	"hexagonal-architecture/pkg/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	userService user.Service
}

func NewAuthHandler(userService user.Service) *authHandler {
	return &authHandler{userService}
}

func (h *authHandler) Login(e echo.Context) error {
	var loginInput user.LoginRequest
	e.Bind(&loginInput)

	if err := e.Validate(loginInput); err != nil {
		errorMessage := constant.FormatValidationError(err)
		response := constant.APIResponse("Error login", http.StatusBadRequest, false, errorMessage)
		return e.JSON(http.StatusBadRequest, response)
	}

	userLogin, err := h.userService.Login(loginInput)

	if err != nil {
		errorMessage := err.Error()
		response := constant.APIResponse("Error login", http.StatusBadRequest, false, errorMessage)
		return e.JSON(http.StatusBadRequest, response)
	}

	token, err := middleware.CreateToken(int(userLogin.ID))

	if err != nil {
		errorMessage := err.Error()
		response := constant.APIResponse("Error login", http.StatusInternalServerError, false, errorMessage)
		return e.JSON(http.StatusInternalServerError, response)
	}

	formatter := user.FormatLogin(userLogin, token)
	response := constant.APIResponse("Login Success", http.StatusCreated, true, formatter)
	return e.JSON(http.StatusCreated, response)

}
