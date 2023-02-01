package integrationTest

import (
	"LearnECHO/internal/handlers"
	"LearnECHO/internal/router"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
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
	//connectDB, err := setupTestDB()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

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
	//
	//body, _ := io.ReadAll(response.Body)
	//var responseBody map[string]interface{}
	//json.Unmarshal(body, &responseBody)
	//
	//// Assertion
	//assert.Equal(t, 201, int(responseBody["status"].(float64)))
	//data, ok := responseBody["data"].(map[string]interface{})
	//if !ok {
	//	t.Fatal("Data field is not present in response body")
	//}
	//assert.Equal(t, "test Title", data["title"])
	//assert.Equal(t, "test Description", data["description"])

}

func TestCreateTodoListFailed(t *testing.T) {

}
