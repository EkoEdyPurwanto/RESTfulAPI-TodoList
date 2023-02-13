package handlers

import (
	"LearnECHO/models/requestAndresponse"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func toJSON(t interface{}) string {
	bytes, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func TestTodoListHandlerImpl_Create(t *testing.T) {
	db, err := sql.Open("mysql", "eep:1903@/RESTfulAPI_todos_test")
	if err != nil {
		log.Fatal(err.Error())
	}
	handler := NewTodoListHandlerImpl(db)

	type args struct {
		request requestAndresponse.TodoListCreateRequest
	}

	tests := []struct {
		name    string
		handler *TodoListHandlerImpl
		args    args
		wantErr bool
		want    int
	}{
		{
			name:    "Success - Create Title And Description Todo",
			handler: handler,
			args: args{
				request: requestAndresponse.TodoListCreateRequest{
					Title:       "Create Todo Title",
					Description: "Create Todo Description",
				},
			},
			wantErr: false,
			want:    http.StatusCreated,
		},
		{
			name:    "Error - Bad request due to invalid title",
			handler: handler,
			args: args{
				request: requestAndresponse.TodoListCreateRequest{
					Title:       "",
					Description: "Updated Todo Description",
				},
			},
			wantErr: true,
			want:    http.StatusBadRequest,
		},
		{
			name:    "Error - Bad request due to invalid description",
			handler: handler,
			args: args{
				request: requestAndresponse.TodoListCreateRequest{
					Title:       "Updated Todo Title",
					Description: "",
				},
			},
			wantErr: true,
			want:    http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPatch, "http://localhost:1234/api.todolist.com/todolist/managed-todolist/", strings.NewReader(toJSON(tt.args.request)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			c.SetPath("/api.todolist.com/todolist/managed-todolist")

			err := tt.handler.Create(c, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTitleAndDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if rec.Code != tt.want {
				t.Errorf("UpdateTitleAndDescription() got = %v, want %v", rec.Code, tt.want)
			}
		})
	}
}

func TestTodoListHandlerImpl_ReadAll(t *testing.T) {
}

func TestTodoListHandlerImpl_ReadById(t *testing.T) {
}

func TestTodoListHandlerImpl_UpdateTitleAndDescription(t *testing.T) {
	db, err := sql.Open("mysql", "eep:1903@/RESTfulAPI_todos_test")
	if err != nil {
		log.Fatal(err.Error())
	}
	handler := NewTodoListHandlerImpl(db)

	type args struct {
		todolistId int
		request    requestAndresponse.TodoListUpdateTitleDescription
	}

	tests := []struct {
		name    string
		handler *TodoListHandlerImpl
		args    args
		wantErr bool
		want    int
	}{
		{
			name:    "Success - Update Title And Description Todo",
			handler: handler,
			args: args{
				todolistId: 1,
				request: requestAndresponse.TodoListUpdateTitleDescription{
					Title:       "Updated Todo Title",
					Description: "Updated Todo Description",
				},
			},
			wantErr: false,
			want:    http.StatusOK,
		},
		{
			name:    "Error - Missing Todo id in the database",
			handler: handler,
			args: args{
				todolistId: 99,
				request: requestAndresponse.TodoListUpdateTitleDescription{
					Title:       "Updated Todo Title",
					Description: "Updated Todo Description",
				},
			},
			wantErr: true,
			want:    http.StatusNotFound,
		},
		{
			name:    "Error - Bad request due to invalid title",
			handler: handler,
			args: args{
				todolistId: 1,
				request: requestAndresponse.TodoListUpdateTitleDescription{
					Title:       "",
					Description: "Updated Todo Description",
				},
			},
			wantErr: true,
			want:    http.StatusBadRequest,
		},
		{
			name:    "Error - Bad request due to invalid description",
			handler: handler,
			args: args{
				todolistId: 1,
				request: requestAndresponse.TodoListUpdateTitleDescription{
					Title:       "Updated Todo Title",
					Description: "",
				},
			},
			wantErr: true,
			want:    http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPatch, "http://localhost:1234/api.todolist.com/todolists/managed-todolists/", strings.NewReader(toJSON(tt.args.request)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			c.SetPath("/api.todolist.com/todolists/managed-todolists/:todolistId")
			c.SetParamNames("todolistId")
			c.SetParamValues(strconv.Itoa(tt.args.todolistId))

			err := tt.handler.UpdateTitleAndDescription(c, tt.args.todolistId, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTitleAndDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if rec.Code != tt.want {
				t.Errorf("UpdateTitleAndDescription() got = %v, want %v", rec.Code, tt.want)
			}
		})
	}
}

func TestTodoListHandlerImpl_UpdateStatus(t *testing.T) {
	db, err := sql.Open("mysql", "eep:1903@/RESTfulAPI_todos_test")
	if err != nil {
		log.Fatal(err.Error())
	}
	handler := NewTodoListHandlerImpl(db)

	type args struct {
		todolistId int
		request    requestAndresponse.TodoListUpdateStatus
	}

	tests := []struct {
		name    string
		handler *TodoListHandlerImpl
		args    args
		wantErr bool
		want    int
	}{
		{
			name:    "Success - Update Title And Description Todo",
			handler: handler,
			args: args{
				todolistId: 1,
				request: requestAndresponse.TodoListUpdateStatus{
					Status: "DONE",
				},
			},
			wantErr: false,
			want:    http.StatusOK,
		},
		{
			name:    "Error - Missing Todo id in the database",
			handler: handler,
			args: args{
				todolistId: 99,
				request: requestAndresponse.TodoListUpdateStatus{
					Status: "DONE",
				},
			},
			wantErr: true,
			want:    http.StatusNotFound,
		},
		{
			name:    "Error - Bad request due to invalid title",
			handler: handler,
			args: args{
				todolistId: 1,
				request: requestAndresponse.TodoListUpdateStatus{
					Status: "",
				},
			},
			wantErr: true,
			want:    http.StatusBadRequest,
		},
		{
			name:    "Error - Bad request due to invalid description",
			handler: handler,
			args: args{
				todolistId: 1,
				request: requestAndresponse.TodoListUpdateStatus{
					Status: "GK BISA NGASAL KARNA INI TYPE DATA ENUM('PENDING', 'DONE')",
				},
			},
			wantErr: true,
			want:    http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPatch, "http://localhost:1234/api.todolist.com/todolist/managed-todolist/", strings.NewReader(toJSON(tt.args.request)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			c.SetPath("/api.todolist.com/todolist/managed-todolist/:todolistId")
			c.SetParamNames("todolistId")
			c.SetParamValues(strconv.Itoa(tt.args.todolistId))

			err := tt.handler.UpdateStatus(c, tt.args.todolistId, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if rec.Code != tt.want {
				t.Errorf("UpdateStatus() got = %v, want %v", rec.Code, tt.want)
			}
		})
	}
}

func TestTodoListHandlerImpl_Delete(t *testing.T) {
	db, err := sql.Open("mysql", "eep:1903@/RESTfulAPI_todos_test")
	if err != nil {
		log.Fatal(err.Error())
	}
	handler := NewTodoListHandlerImpl(db)

	type args struct {
		todolistId int
	}

	tests := []struct {
		name    string
		handler *TodoListHandlerImpl
		args    args
		wantErr bool
		want    int
	}{
		{
			name:    "Success - Delete Title And Description Todo",
			handler: handler,
			args: args{
				todolistId: 1,
			},
			wantErr: false,
			want:    http.StatusOK,
		},
		{
			name:    "Error - Missing Todo id in the database",
			handler: handler,
			args: args{
				todolistId: 99,
			},
			wantErr: true,
			want:    http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPatch, "http://localhost:1234/api.todolist.com/todolist/manage-todolist/", strings.NewReader(toJSON(tt.args.todolistId)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			c.SetPath("/api.todolist.com/todolist/manage-todolist/:todolistId")
			c.SetParamNames("todolistId")
			c.SetParamValues(strconv.Itoa(tt.args.todolistId))

			err := tt.handler.Delete(c, tt.args.todolistId)

			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if rec.Code != tt.want {
				t.Errorf("DeleteStatus() got = %v, want %v", rec.Code, tt.want)
			}
		})
	}
}
