package handlers

import (
	"LearnECHO/models/requestAndresponse"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockTodoListHandler struct {
	mock.Mock
	TodoListHandler
}

func (m *MockTodoListHandler) Create(c echo.Context, request requestAndresponse.TodoListCreateRequest) error {
	args := m.Called(c, request)
	return args.Error(0)
}

func TestMockTodoListHandlerImpl_Create(t *testing.T) {
	handler := new(MockTodoListHandler)

	type args struct {
		request requestAndresponse.TodoListCreateRequest
	}

	tests := []struct {
		name    string
		handler *MockTodoListHandler
		args    args
		wantErr bool
		want    int
	}{
		// test cases here
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
			// set up the request and context
			req := httptest.NewRequest(http.MethodPatch, "http://localhost:1234/api.todolist.com/todolist/managed-todolist/", strings.NewReader(toJSON(tt.args.request)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			c.SetPath("/api.todolist.com/todolist/managed-todolist")

			// set up the expected behavior of the mock object
			if tt.wantErr {
				handler.On("Create", c, tt.args.request).Return(errors.New("error occurred"))
			} else {
				handler.On("Create", c, tt.args.request).Return(nil)
			}

			// call the method being tested
			err := handler.Create(c, tt.args.request)

			// verify the results
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.want, rec.Code)
			handler.AssertExpectations(t)
		})
	}
}
