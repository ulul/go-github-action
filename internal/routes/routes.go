package routes

import (
	"hexagonal-architecture/database"
	"hexagonal-architecture/internal/app/book"
	"hexagonal-architecture/internal/app/user"
	"hexagonal-architecture/internal/handler"
	"hexagonal-architecture/internal/middleware"
	"os"

	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	middleware.LogMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to rest api")
	})

	userRepository := user.NewRepository(database.DB)
	bookRepository := book.NewRepository(database.DB)

	userService := user.NewUserService(userRepository)
	bookService := book.NewBookService(bookRepository)

	userHandler := handler.NewUserHandler(userService)
	bookHandler := handler.NewBookHandler(bookService)
	authHandler := handler.NewAuthHandler(userService)

	e.POST("/login", authHandler.Login)
	e.POST("/users", userHandler.CreateUser)

	e.GET("/users", userHandler.GetUser, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/users/:id", userHandler.GetUserByID, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.PUT("/users/:id", userHandler.UpdateUser, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.DELETE("/users/:id", userHandler.DeleteUser, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))

	e.GET("/books", bookHandler.GetBook, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/books/:id", bookHandler.GetBookByID, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.POST("/books", bookHandler.CreateBook, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.PUT("/books/:id", bookHandler.UpdateBook, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.DELETE("/books/:id", bookHandler.DeleteBook, echoMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))

	return e
}
