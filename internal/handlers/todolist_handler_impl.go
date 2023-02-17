package handlers

import (
	"LearnECHO/helper"
	"LearnECHO/models/domain"
	"LearnECHO/models/requestAndresponse"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type TodoListHandlerImpl struct {
	DB *sql.DB
}

func NewTodoListHandlerImpl(DB *sql.DB) *TodoListHandlerImpl {
	return &TodoListHandlerImpl{DB: DB}
}

func (handler *TodoListHandlerImpl) Create(ctx echo.Context, request requestAndresponse.TodoListCreateRequest) error {

	err := ctx.Bind(&request)
	if err != nil {
		helper.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		helper.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	SQL, err := handler.DB.Exec("INSERT INTO TodoList(title, description) VALUES(?, ?)", request.Title, request.Description)
	if err != nil {
		helper.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	lastID, err := SQL.LastInsertId()
	if err != nil {
		helper.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	response := domain.Response{
		Status:  http.StatusCreated,
		Message: "you have successfully created todo list with ID: " + strconv.FormatInt(lastID, 10),
	}
	log.Print(response.Message)

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	ctx.Response().WriteHeader(response.Status)
	helper.WriteToResponseBody(ctx, response)

	return nil
}

func (handler *TodoListHandlerImpl) ReadAll(ctx echo.Context) error {
	var todos requestAndresponse.TodoListResponse
	var sliceTodos []requestAndresponse.TodoListResponse

	rows, err := handler.DB.Query("SELECT id, title, description, status FROM TodoList")
	if err != nil {
		helper.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&todos.Id, &todos.Title, &todos.Description, &todos.Status)
		if err != nil {
			log.Fatal(err)
		} else {
			sliceTodos = append(sliceTodos, todos)
		}
	}

	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    sliceTodos,
	}
	log.Print("Read All Todo successfully")

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	ctx.Response().WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(ctx, apiResponse)

	return nil
}

func (handler *TodoListHandlerImpl) ReadById(ctx echo.Context, todolistId int) error {
	var todos requestAndresponse.TodoListResponse
	var arrTodos []requestAndresponse.TodoListResponse

	var count int
	if err := handler.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", todolistId).Scan(&count); err != nil {
		helper.InternalServerError(err, ctx)
		log.Error("Failed to check Todo existence in the database")
		return err
	}

	if count == 0 {
		helper.NotFound(errors.New(" id not found in db"), ctx)
		return nil
	}

	rows, err := handler.DB.Query("SELECT id, title, description, status FROM TodoList WHERE id = ?", todolistId)
	if err != nil {
		helper.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	for rows.Next() {
		rows.Scan(&todos.Id, &todos.Title, &todos.Description, &todos.Status)

		if err != nil {
			log.Fatal(err)
		} else {
			arrTodos = append(arrTodos, todos)
		}
	}

	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    todos,
	}
	log.Info("Read Id Todo successfully")

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	ctx.Response().WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(ctx, apiResponse)

	return nil
}

func (handler *TodoListHandlerImpl) UpdateTitleAndDescription(ctx echo.Context, todolistId int, request requestAndresponse.TodoListUpdateTitleDescription) error {

	err := ctx.Bind(&request)
	if err != nil {
		log.Error(err)
		panic(err)
		return err
	}

	var count int
	if err := handler.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", todolistId).Scan(&count); err != nil {
		helper.InternalServerError(err, ctx)
		log.Error("Failed to check Todo existence in the database")
		return err
	}

	if count == 0 {
		helper.NotFound(errors.New(" id not found in db"), ctx)
		return errors.New("id not found")
	}

	validate := validator.New()
	err = validate.Struct(requestAndresponse.TodoListUpdateTitleDescription{
		Title:       request.Title,
		Description: request.Description,
	})

	if err != nil {
		helper.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	if request.Title != "" && request.Description == "" {
		_, err = handler.DB.Exec("UPDATE TodoList SET title=? WHERE id=?", request.Title, todolistId)
	} else if request.Description != "" && request.Title == "" {
		_, err = handler.DB.Exec("UPDATE TodoList SET description=? WHERE id=?", request.Description, todolistId)
	} else {
		_, err = handler.DB.Exec("UPDATE TodoList SET title=?, description=? WHERE id=?", request.Title, request.Description, todolistId)
	}

	if err != nil {
		helper.InternalServerError(err, ctx)
		log.Print(err)
	}

	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Success",
	}
	log.Info("Update Title & Description Todo successfully")

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	ctx.Response().WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(ctx, apiResponse)

	return nil

}

func (handler *TodoListHandlerImpl) UpdateStatus(ctx echo.Context, todolistId int, request requestAndresponse.TodoListUpdateStatus) error {
	err := ctx.Bind(&request)
	if err != nil {
		log.Error(err)
		panic(err)
		return err
	}

	var count int
	if err := handler.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", todolistId).Scan(&count); err != nil {
		helper.InternalServerError(err, ctx)
		log.Error("Failed to check Todo existence in the database")
		return err
	}

	if count == 0 {
		helper.NotFound(errors.New(" id not found in db"), ctx)
		return errors.New("id not found")
	}

	validate := validator.New()
	err = validate.Struct(requestAndresponse.TodoListUpdateStatus{
		Status: request.Status,
	})

	if err != nil {
		helper.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	_, err = handler.DB.Exec("UPDATE TodoList SET status=? WHERE id=?", request.Status, todolistId)

	if err != nil {
		helper.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Success",
	}
	log.Print("Update Status Todo successfully")

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	ctx.Response().WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(ctx, apiResponse)

	return nil
}

func (handler *TodoListHandlerImpl) Delete(ctx echo.Context, todolistId int) error {

	var count int
	if err := handler.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", todolistId).Scan(&count); err != nil {
		helper.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	if count == 0 {
		helper.NotFound(errors.New(" id not found in the db"), ctx)
		return errors.New("id not found")
	}

	if _, err := handler.DB.Exec("DELETE FROM TodoList WHERE id=?", todolistId); err != nil {
		helper.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Todo with id " + strconv.Itoa(todolistId) + " has been deleted",
	}
	log.Info("Delete Todo successfully")

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	ctx.Response().WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(ctx, apiResponse)

	return nil
}
