package integrationTest

import (
	"LearnECHO/internal/handlers"
	"LearnECHO/internal/router"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func setupTestDB() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "eep"
	dbPass := "1903"
	dbName := "RESTfulAPI_todos_test"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func truncateTodoList(db *sql.DB) {
	db.Exec("TRUNCATE TABLE TodoList ")
}

func setupRouter(db *sql.DB) *echo.Echo {
	todoListHandler := handlers.NewTodoListHandlerImpl(db)

	e := router.NewRouter(todoListHandler)

	return e
}

func TestCreateTodolistSuccess(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)
	router := setupRouter(conn)

	requestBody := strings.NewReader(`{"title" : "test Title", "description": "test Description"}`)

	httpRequest := httptest.NewRequest(echo.POST, "http://localhost:1234/api.todolist.com/todolist/managed-todolist", requestBody)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	assert.Equal(t, 201, int(responseBody["status"].(float64)))
	//assert.Equal(t, "you have successfully created todo list with ID: "+strconv.FormatInt(id, 10), responseBody["message"])

	//data, ok := responseBody["data"].(map[string]interface{})
	//if !ok {
	//	t.Fatal("Data field is not present in response body")
	//}
	//assert.Equal(t, "test Title", data["title"])
	//assert.Equal(t, "test Description", data["description"])

}

func TestCreateTodoListFailed(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)
	router := setupRouter(conn)

	requestBody := strings.NewReader(`{"title" : "", "description": ""}`)
	httpRequest := httptest.NewRequest(echo.POST, "http://localhost:1234/api.todolist.com/todolist/managed-todolist", requestBody)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	assert.Equal(t, 400, int(responseBody["status"].(float64)))
	assert.Equal(t, "bad request", responseBody["message"])
}

func TestUpdateTitleAndDescriptionSuccess(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)

	tx, _ := conn.Begin()
	result, err := tx.Exec("INSERT INTO TodoList (title, description) VALUES (?, ?)", "Buy Milk", "Buy 1 liter of milk")
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	tx.Commit()

	router := setupRouter(conn)

	requestBody := strings.NewReader(`{"title" : "Update Title", "description": "Update Description"}`)
	httpRequest := httptest.NewRequest(echo.PATCH, "http://localhost:1234/api.todolist.com/todolists/managed-todolists/"+strconv.FormatInt(id, 10), requestBody)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	//body, _ := io.ReadAll(response.Body)
	//var responseBody map[string]interface{}
	//json.Unmarshal(body, &responseBody)
	//
	//// Assertion
	//assert.Equal(t, 200, int(responseBody["status"].(float64)))
	//data, ok := responseBody["data"].(map[string]interface{})
	//if !ok {
	//	t.Fatal("Data field is not present in response body")
	//}
	//assert.Equal(t, "Update Title", data["title"])
	//assert.Equal(t, "Update Description", data["description"])
}

func TestUpdateTitleAndDescriptionFailed(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)

	tx, _ := conn.Begin()
	result, err := tx.Exec("INSERT INTO TodoList (title, description) VALUES (?, ?)", "Buy Milk", "Buy 1 liter of milk")
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	tx.Commit()

	router := setupRouter(conn)

	requestBody := strings.NewReader(`{"title" : "", "description": ""}`)
	httpRequest := httptest.NewRequest(echo.PATCH, "http://localhost:1234/api.todolist.com/todolists/managed-todolists/"+strconv.FormatInt(id, 10), requestBody)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	assert.Equal(t, 400, int(responseBody["status"].(float64)))
	assert.Equal(t, "bad request", responseBody["message"])
}

func TestUpdateStatusSuccess(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)

	tx, _ := conn.Begin()
	result, err := tx.Exec("INSERT INTO TodoList (title, description) VALUES (?, ?)", "Buy Milk", "Buy 1 liter of milk")
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	tx.Commit()

	router := setupRouter(conn)

	requestBody := strings.NewReader(`{"status" : "DONE"}`)
	httpRequest := httptest.NewRequest(echo.PUT, "http://localhost:1234/api.todolist.com/todolist/managed-todolist/"+strconv.FormatInt(id, 10), requestBody)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	//body, _ := io.ReadAll(response.Body)
	//var responseBody map[string]interface{}
	//json.Unmarshal(body, &responseBody)
	//
	//// Assertion
	//assert.Equal(t, 200, int(responseBody["status"].(float64)))
	//
	//data, ok := responseBody["data"].(map[string]interface{})
	//if !ok {
	//	t.Fatal("Data field is not present in response body")
	//}
	//assert.Equal(t, "DONE", data["status"])
}

func TestUpdateStatusFailed(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)

	tx, _ := conn.Begin()
	result, err := tx.Exec("INSERT INTO TodoList (title, description) VALUES (?, ?)", "Buy Milk", "Buy 1 liter of milk")
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	tx.Commit()

	router := setupRouter(conn)

	requestBody := strings.NewReader(`{"status" : ""}`)
	httpRequest := httptest.NewRequest(echo.PATCH, "http://localhost:1234/api.todolist.com/todolists/managed-todolists/"+strconv.FormatInt(id, 10), requestBody)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	assert.Equal(t, 400, int(responseBody["status"].(float64)))
	assert.Equal(t, "bad request", responseBody["message"])
}

func TestGetByIdFailed(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)
	router := setupRouter(conn)

	httpRequest := httptest.NewRequest(echo.GET, "http://localhost:1234/api.todolist.com/todolists/managed-todolists/404", nil)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	assert.Equal(t, 404, int(responseBody["status"].(float64)))
	assert.Equal(t, "not found", responseBody["message"])

}

func TestDeleteSuccess(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)

	tx, _ := conn.Begin()
	result, err := tx.Exec("INSERT INTO TodoList (title, description) VALUES (?, ?)", "Buy Milk", "Buy 1 liter of milk")
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	tx.Commit()

	router := setupRouter(conn)

	httpRequest := httptest.NewRequest(echo.DELETE, "http://localhost:1234/api.todolist.com/todolist/manage-todolist/"+strconv.FormatInt(id, 10), nil)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	assert.Equal(t, 200, int(responseBody["status"].(float64)))
	assert.Equal(t, "Todo with id "+strconv.FormatInt(id, 10)+" has been deleted", responseBody["message"])
}

func TestDeleteFailed(t *testing.T) {
	// setup
	conn, err := setupTestDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	truncateTodoList(conn)
	router := setupRouter(conn)

	httpRequest := httptest.NewRequest(echo.DELETE, "http://localhost:1234/api.todolist.com/todolist/manage-todolist/404", nil)
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpRequest.Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httpRequest)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	assert.Equal(t, 404, int(responseBody["status"].(float64)))
	assert.Equal(t, "not found", responseBody["message"])
}
