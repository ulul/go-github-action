package handler

import (
	"hexagonal-architecture/internal/app/user"
	"hexagonal-architecture/pkg/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUser(e echo.Context) error {
	users, err := h.userService.Get()

	if err != nil {
		errorMessage := constant.FormatValidationError(err)

		response := constant.APIResponse("Error retrive data", http.StatusInternalServerError, false, errorMessage)
		return e.JSON(http.StatusInternalServerError, response)

	}

	formatter := user.FormatUsers(users)
	response := constant.APIResponse("User retrieved", http.StatusOK, true, formatter)
	return e.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUserByID(e echo.Context) error {
	id := e.Param("id")
	userDetail, err := h.userService.FindByID(id)

	if err != nil {
		errorMessage := constant.FormatValidationError(err)
		response := constant.APIResponse("Error retrive data", http.StatusInternalServerError, false, errorMessage)
		return e.JSON(http.StatusInternalServerError, response)

	}

	formatter := user.FormatUser(userDetail)
	response := constant.APIResponse("User Retrieved", http.StatusOK, true, formatter)
	return e.JSON(http.StatusOK, response)

}

func (h *userHandler) CreateUser(e echo.Context) error {
	var inputUser user.CreateUserRequest
	e.Bind(&inputUser)

	if err := e.Validate(inputUser); err != nil {
		errorMessage := constant.FormatValidationError(err)
		response := constant.APIResponse("Error create data", http.StatusBadRequest, false, errorMessage)
		return e.JSON(http.StatusBadRequest, response)
	}

	userDetail, err := h.userService.Create(inputUser)

	if err != nil {
		errorMessage := err.Error()
		response := constant.APIResponse("Error create data", http.StatusInternalServerError, false, errorMessage)
		return e.JSON(http.StatusInternalServerError, response)
	}

	formatter := user.FormatUser(userDetail)
	response := constant.APIResponse("User Created", http.StatusCreated, true, formatter)
	return e.JSON(http.StatusCreated, response)

}

func (h *userHandler) UpdateUser(e echo.Context) error {
	var inputUser user.UpdateUserRequest
	e.Bind(&inputUser)

	if err := e.Validate(inputUser); err != nil {
		errorMessage := constant.FormatValidationError(err)
		response := constant.APIResponse("Error update data", http.StatusBadRequest, false, errorMessage)
		return e.JSON(http.StatusBadRequest, response)
	}

	id := e.Param("id")
	userDetail, err := h.userService.Update(id, inputUser)

	if err != nil {
		errorMessage := err.Error()
		response := constant.APIResponse("Error update data", http.StatusInternalServerError, false, errorMessage)
		return e.JSON(http.StatusInternalServerError, response)
	}

	formatter := user.FormatUser(userDetail)
	response := constant.APIResponse("User Updated", http.StatusOK, true, formatter)
	return e.JSON(http.StatusOK, response)

}

func (h *userHandler) DeleteUser(e echo.Context) error {
	id := e.Param("id")

	err := h.userService.Delete(id)

	if err != nil {
		errorMessage := err.Error()
		response := constant.APIResponse("Error delete data", http.StatusBadRequest, false, errorMessage)
		return e.JSON(http.StatusBadRequest, response)
	}

	response := constant.APIResponse("User Deleted", http.StatusOK, true, nil)
	return e.JSON(http.StatusOK, response)

}
