package handlers

import (
	"LearnECHO/models/domain"
	"LearnECHO/models/requestAndresponse"
	"LearnECHO/utils"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
)

type TodoListHandlerImpl struct {
	DB *sql.DB
}

func NewTodoListHandlerImpl(DB *sql.DB) *TodoListHandlerImpl {
	return &TodoListHandlerImpl{DB: DB}
}

func (handler *TodoListHandlerImpl) Create(ctx echo.Context, request requestAndresponse.TodoListCreateRequest) error {
	// Check authentication
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		utils.UnauthorizedError(errors.New("missing token"), ctx)
		return errors.New("missing token")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateJWTToken(tokenString)
	if err != nil {
		utils.UnauthorizedError(err, ctx)
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	err = ctx.Bind(&request)
	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	// Insert TodoList item
	SQL := `INSERT INTO TodoList(user_id, title, description) VALUES($1, $2, $3) RETURNING todo_id`
	var id int64
	err = handler.DB.QueryRowContext(ctx.Request().Context(), SQL, int64(userID), request.Title, request.Description).Scan(&id)
	if err != nil {
		utils.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	response := domain.Response{
		Status:  http.StatusCreated,
		Message: "you have successfully created todo list with ID: " + strconv.FormatInt(id, 10),
	}
	log.Print(response.Message)

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	ctx.Response().WriteHeader(response.Status)
	utils.WriteToResponseBody(ctx, response)

	return nil
}

func (handler *TodoListHandlerImpl) ReadAll(ctx echo.Context) error {
	// Check authentication
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		utils.UnauthorizedError(errors.New("missing token"), ctx)
		return errors.New("missing token")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateJWTToken(tokenString)
	if err != nil {
		utils.UnauthorizedError(err, ctx)
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	var todos requestAndresponse.TodoListResponse
	var sliceTodos []requestAndresponse.TodoListResponse

	rows, err := handler.DB.Query("SELECT todo_id, user_id, title, description, status FROM TodoList WHERE user_id = $1", int64(userID))
	if err != nil {
		utils.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&todos.TodoID, &todos.UserID, &todos.Title, &todos.Description, &todos.Status)
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
	utils.WriteToResponseBody(ctx, apiResponse)

	return nil
}

func (handler *TodoListHandlerImpl) ReadById(ctx echo.Context, todolistId int) error {
	// Check authentication
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		utils.UnauthorizedError(errors.New("missing token"), ctx)
		return errors.New("missing token")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateJWTToken(tokenString)
	if err != nil {
		utils.UnauthorizedError(err, ctx)
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))

	var count int
	if err := handler.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE todo_id=$1 AND user_id=$2", todolistId, userID).Scan(&count); err != nil {
		utils.InternalServerError(err, ctx)
		log.Error("Failed to check Todo existence in the database")
		return err
	}

	if count == 0 {
		utils.NotFound(errors.New(" id not found in db"), ctx)
		return nil
	}

	var todo requestAndresponse.TodoListResponse
	row := handler.DB.QueryRow("SELECT todo_id, user_id, title, description, status, created_at, updated_at FROM TodoList WHERE todo_id = $1 AND user_id = $2", todolistId, userID)
	err = row.Scan(&todo.TodoID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		utils.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	// Return the response
	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    todo,
	}
	log.Info("Read Id Todo successfully")

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	ctx.Response().WriteHeader(apiResponse.Status)
	utils.WriteToResponseBody(ctx, apiResponse)

	return nil
}

func (handler *TodoListHandlerImpl) UpdateTitleAndDescription(ctx echo.Context, todolistId int, request requestAndresponse.TodoListUpdateTitleDescription) error {
	// Check authentication
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		utils.UnauthorizedError(errors.New("missing token"), ctx)
		return errors.New("missing token")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateJWTToken(tokenString)
	if err != nil {
		utils.UnauthorizedError(err, ctx)
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	err = ctx.Bind(&request)
	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	var count int
	if err := handler.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE todo_id=$1 AND user_id=$2", todolistId, int64(userID)).Scan(&count); err != nil {
		utils.InternalServerError(err, ctx)
		log.Error("Failed to check Todo existence in the database")
		return err
	}

	if count == 0 {
		utils.NotFound(errors.New(" id not found in db"), ctx)
		return errors.New("id not found")
	}

	validate := validator.New()
	err = validate.Struct(requestAndresponse.TodoListUpdateTitleDescription{
		Title:       request.Title,
		Description: request.Description,
	})

	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	if request.Title != "" && request.Description == "" {
		_, err = handler.DB.Exec("UPDATE TodoList SET title=$1 WHERE todo_id=$2 AND user_id=$3", request.Title, todolistId, int64(userID))
	} else if request.Description != "" && request.Title == "" {
		_, err = handler.DB.Exec("UPDATE TodoList SET description=$1 WHERE todo_id=$2 AND user_id=$3", request.Description, todolistId, int64(userID))
	} else {
		_, err = handler.DB.Exec("UPDATE TodoList SET title=$1, description=$2 WHERE todo_id=$3 AND user_id=$4", request.Title, request.Description, todolistId, int64(userID))
	}

	if err != nil {
		utils.InternalServerError(err, ctx)
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
	utils.WriteToResponseBody(ctx, apiResponse)

	return nil
}

func (handler *TodoListHandlerImpl) UpdateStatus(ctx echo.Context, todolistId int, request requestAndresponse.TodoListUpdateStatus) error {
	// Check authentication
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		utils.UnauthorizedError(errors.New("missing token"), ctx)
		return errors.New("missing token")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateJWTToken(tokenString)
	if err != nil {
		utils.UnauthorizedError(err, ctx)
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	// Check if the user has access to update the todo list item
	var count int
	if err := handler.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE todo_id=$1 AND user_id=$2", todolistId, int(userID)).Scan(&count); err != nil {
		utils.InternalServerError(err, ctx)
		log.Error("Failed to check Todo existence in the database")
		return err
	}

	if count == 0 {
		utils.NotFound(errors.New("id not found in db"), ctx)
		return errors.New("id not found")
	}

	err = ctx.Bind(&request)
	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	validate := validator.New()
	err = validate.Struct(requestAndresponse.TodoListUpdateStatus{
		Status: request.Status,
	})

	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	_, err = handler.DB.Exec("UPDATE TodoList SET status=$1 WHERE todo_id=$2", request.Status, todolistId)

	if err != nil {
		utils.InternalServerError(err, ctx)
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
	utils.WriteToResponseBody(ctx, apiResponse)

	return nil
}

func (handler *TodoListHandlerImpl) Delete(ctx echo.Context, todolistId int) error {
	// Check authentication
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		utils.UnauthorizedError(errors.New("missing token"), ctx)
		return errors.New("missing token")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := utils.ValidateJWTToken(tokenString)
	if err != nil {
		utils.UnauthorizedError(err, ctx)
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))

	// Check if the TodoList item belongs to the user
	var count int
	if err := handler.DB.QueryRowContext(ctx.Request().Context(), "SELECT COUNT(*) FROM TodoList WHERE todo_id=$1 AND user_id=$2", todolistId, userID).Scan(&count); err != nil {
		utils.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	if count == 0 {
		utils.NotFound(errors.New(" id not found in the db"), ctx)
		return errors.New("id not found")
	}

	// Delete the TodoList item
	if _, err := handler.DB.ExecContext(ctx.Request().Context(), "DELETE FROM TodoList WHERE todo_id=$1 AND user_id=$2", todolistId, userID); err != nil {
		utils.InternalServerError(err, ctx)
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
	utils.WriteToResponseBody(ctx, apiResponse)

	return nil
}

func (handler *TodoListHandlerImpl) Login(ctx echo.Context, request domain.Users) error {
	err := ctx.Bind(&request)
	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	// Find the user by their username
	SQL := `SELECT user_id, username, password FROM users WHERE username = $1`
	row := handler.DB.QueryRowContext(ctx.Request().Context(), SQL, request.Username)

	var userID int64
	var username, hashedPassword string
	err = row.Scan(&userID, &username, &hashedPassword)
	if err != nil {
		utils.UnauthorizedError(errors.New("invalid credentials"), ctx)
		log.Error(err)
		return err
	}

	// Compare the provided password with the hashed password from the database
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(request.Password))
	if err != nil {
		utils.UnauthorizedError(errors.New("invalid credentialsSS"), ctx)
		log.Error(err)
		return err
	}

	// Generate a JWT token for the authenticated user
	token, err := utils.GenerateJWTToken(userID)
	if err != nil {
		utils.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	response := domain.Response{
		Status:  http.StatusOK,
		Message: "Welcome " + username + "! You have successfully logged in.",
		Data:    token,
	}
	log.Print(response.Message)

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	ctx.Response().WriteHeader(response.Status)
	utils.WriteToResponseBody(ctx, response)

	return nil
}

func (handler *TodoListHandlerImpl) Register(ctx echo.Context, request domain.Users) error {
	err := ctx.Bind(&request)
	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		utils.BadRequest(err, ctx)
		log.Error(err)
		return err
	}

	// Hash the user's password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	// Insert the new user into the database
	SQL := `INSERT INTO users(username, password, email) VALUES($1, $2, $3) RETURNING user_id`
	var userID int64
	err = handler.DB.QueryRowContext(ctx.Request().Context(), SQL, request.Username, hashedPassword, request.Email).Scan(&userID)
	if err != nil {
		utils.InternalServerError(err, ctx)
		log.Error(err)
		return err
	}

	response := domain.Response{
		Status:  http.StatusCreated,
		Message: "User registration successful with ID: " + strconv.FormatInt(userID, 10),
	}
	log.Print(response.Message)

	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	ctx.Response().WriteHeader(response.Status)
	utils.WriteToResponseBody(ctx, response)

	return nil
}
