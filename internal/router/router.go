package router

import (
	"LearnECHO/internal/handlers"
	"LearnECHO/models/requestAndresponse"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func NewRouter(todoListHandler handlers.TodoListHandler) *echo.Echo {
	e := echo.New()

	e.POST("/api.todolist.com/todolist/managed-todolist", func(ctx echo.Context) error {
		return todoListHandler.Create(ctx, requestAndresponse.TodoListCreateRequest{})
	})

	e.GET("/api.todolist.com/todolists/managed-todolists", func(ctx echo.Context) error {
		return todoListHandler.ReadAll(ctx)
	})

	e.GET("/api.todolist.com/todolist/managed-todolist/:todolistId", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.ReadById(ctx, todolistId)

	})

	e.PUT("/api.todolist.com/todolists/managed-todolists/:todolistId", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.UpdateTitleAndDescription(ctx, todolistId, requestAndresponse.TodoListUpdateTitleDescription{})
	})

	e.PATCH("/api.todolist.com/todolist/managed-todolist/:todolistId", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.UpdateStatus(ctx, todolistId, requestAndresponse.TodoListUpdateStatus{})
	})

	e.DELETE("/api.todolist.com/todolist/manage-todolist/:todolistId", func(ctx echo.Context) error {
		todolistId, err := strconv.Atoi(ctx.Param("todolistId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Todolist ID")
		}
		return todoListHandler.Delete(ctx, todolistId)
	})

	return e
}
