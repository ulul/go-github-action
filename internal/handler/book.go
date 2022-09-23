package handler

import (
	"hexagonal-architecture/internal/app/book"
	"hexagonal-architecture/pkg/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBook(c echo.Context) error {
	books, err := h.bookService.Get()

	if err != nil {
		errorMessage := constant.FormatValidationError(err)

		response := constant.APIResponse("Error retrive data", http.StatusInternalServerError, false, errorMessage)
		return c.JSON(http.StatusInternalServerError, response)
	}

	formatter := book.FormatBooks(books)
	response := constant.APIResponse("Book retrieved", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *bookHandler) GetBookByID(e echo.Context) error {
	id := e.Param("id")
	bookDetail, err := h.bookService.FindByID(id)

	if err != nil {
		errorMessage := constant.FormatValidationError(err)
		response := constant.APIResponse("Error retrive data", http.StatusInternalServerError, false, errorMessage)
		return e.JSON(http.StatusInternalServerError, response)

	}

	formatter := book.FormatBook(bookDetail)
	response := constant.APIResponse("Book Retrieved", http.StatusOK, true, formatter)
	return e.JSON(http.StatusOK, response)

}

func (h *bookHandler) CreateBook(e echo.Context) error {
	var inputUser book.CreateBookRequest
	e.Bind(&inputUser)

	if err := e.Validate(inputUser); err != nil {
		errorMessage := constant.FormatValidationError(err)
		response := constant.APIResponse("Error create data", http.StatusBadRequest, false, errorMessage)
		return e.JSON(http.StatusBadRequest, response)
	}

	bookDetail, err := h.bookService.Create(inputUser)

	if err != nil {
		errorMessage := err.Error()
		response := constant.APIResponse("Error create data", http.StatusInternalServerError, false, errorMessage)
		return e.JSON(http.StatusInternalServerError, response)
	}

	formatter := book.FormatBook(bookDetail)
	response := constant.APIResponse("Book Created", http.StatusCreated, true, formatter)
	return e.JSON(http.StatusCreated, response)

}

func (h *bookHandler) UpdateBook(e echo.Context) error {
	var inputUser book.UpdateBookRequest
	e.Bind(&inputUser)

	if err := e.Validate(inputUser); err != nil {
		errorMessage := constant.FormatValidationError(err)
		response := constant.APIResponse("Error update data", http.StatusBadRequest, false, errorMessage)
		return e.JSON(http.StatusBadRequest, response)
	}

	id := e.Param("id")
	bookDetail, err := h.bookService.UpdateByID(id, inputUser)

	if err != nil {
		errorMessage := err.Error()
		response := constant.APIResponse("Error update data", http.StatusInternalServerError, false, errorMessage)
		return e.JSON(http.StatusInternalServerError, response)
	}

	formatter := book.FormatBook(bookDetail)
	response := constant.APIResponse("Book Updated", http.StatusOK, true, formatter)
	return e.JSON(http.StatusOK, response)

}

func (h *bookHandler) DeleteBook(e echo.Context) error {
	id := e.Param("id")

	err := h.bookService.Delete(id)

	if err != nil {
		errorMessage := constant.FormatValidationError(err)
		response := constant.APIResponse("Error delete data", http.StatusBadRequest, false, errorMessage)
		return e.JSON(http.StatusBadRequest, response)
	}

	response := constant.APIResponse("Book Deleted", http.StatusOK, true, nil)
	return e.JSON(http.StatusOK, response)

}
