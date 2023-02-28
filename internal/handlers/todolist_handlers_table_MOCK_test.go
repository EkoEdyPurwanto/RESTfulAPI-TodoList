package handlers

import (
	"LearnECHO/models/requestAndresponse"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockTodoListHandlerImpl struct {
	mock.Mock
}

func (h *MockTodoListHandlerImpl) Create(c echo.Context, req requestAndresponse.TodoListCreateRequest) error {
	if req.Title != "" && req.Description == "" {
		return c.String(http.StatusBadRequest, "Invalid description")
	} else if req.Description != "" && req.Title == "" {
		return c.String(http.StatusBadRequest, "Invalid title")
	}
	return c.String(http.StatusCreated, "Todo created")
}

func toJSON2(t interface{}) string {
	bytes, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func TestMockTodoListHandlerImpl_Create(t *testing.T) {
	handler := &MockTodoListHandlerImpl{}

	type args struct {
		request requestAndresponse.TodoListCreateRequest
	}

	tests := []struct {
		name    string
		handler *MockTodoListHandlerImpl
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
			req := httptest.NewRequest(http.MethodPatch, "http://localhost:1234/api.todolist.com/todolist/managed-todolist/", strings.NewReader(toJSON2(tt.args.request)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			c.SetPath("/api.todolist.com/todolist/managed-todolist")

			err := tt.handler.Create(c, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if rec.Code != tt.want {
				t.Errorf("Create() got = %v, want %v", rec.Code, tt.want)
			}
		})
	}
}
