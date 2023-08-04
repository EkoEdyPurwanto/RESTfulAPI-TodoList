package router

import (
	"RESTfulAPI-TodoList/internal/handlers"
	"RESTfulAPI-TodoList/models/domain"
	"RESTfulAPI-TodoList/models/requestAndresponse"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

func NewRouter(todoListHandler handlers.TodoListHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set up authentication middleware
	config := middleware.JWTConfig{
		Claims:     &requestAndresponse.JwtCustomClaims{},
		SigningKey: []byte("my-secret-key"),
	}
	authMiddleware := middleware.JWTWithConfig(config)

	// TodoList endpoints
	todoListGroup := e.Group("/api.todolist.com/todolist", authMiddleware)

	todoListGroup.POST("/managed-todolist", func(ctx echo.Context) error {
		return todoListHandler.Create(ctx, requestAndresponse.TodoListCreateRequest{})
	})

	todoListGroup.GET("/managed-todolists", func(ctx echo.Context) error {
		return todoListHandler.ReadAll(ctx)
	})

	todoListGroup.GET("/managed-todolist/:todolistId", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.ReadById(ctx, todolistId)
	})

	todoListGroup.PATCH("/managed-todolists/:todolistId", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.UpdateTitleAndDescription(ctx, todolistId, requestAndresponse.TodoListUpdateTitleDescription{})
	})

	todoListGroup.PUT("/managed-todolist/:todolistId", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.UpdateStatus(ctx, todolistId, requestAndresponse.TodoListUpdateStatus{})
	})

	todoListGroup.DELETE("/manage-todolist/:todolistId", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.Delete(ctx, todolistId)
	})

	todoListGroup.POST("/managed-todolist/:todolistId/upload-picture", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.UploadPicture(ctx, todolistId)
	})

	todoListGroup.GET("/managed-todolist/:todolistId/picture/:pictureId", func(ctx echo.Context) error {
		pictureID, err := strconv.Atoi(ctx.Param("pictureId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Picture ID")
		}
		return todoListHandler.GetPicture(ctx, pictureID)
	})

	todoListGroup.POST("/managed-todolist/upload-s3", func(ctx echo.Context) error {
		// Call your UploadS3 function from here
		return todoListHandler.UploadS3(ctx)
	})

	// User endpoints
	userGroup := e.Group("/api.todolist.com/user")

	userGroup.POST("/register", func(ctx echo.Context) error {
		return todoListHandler.Register(ctx, domain.Users{})
	})

	userGroup.POST("/login", func(ctx echo.Context) error {
		return todoListHandler.Login(ctx, domain.Users{})
	})

	return e
}
