package handlers

import (
	"LearnECHO/models/requestAndresponse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMockTodoListHandler_Create(t *testing.T) {
	// Create a new instance of MockTodoListHandler
	mockHandler := NewMockTodoListHandler(t)

	// Set the expected return value
	mockHandler.On("Create", mock.Anything, mock.Anything).Return(nil)

	// Call the function being tested
	err := mockHandler.Create(nil, requestAndresponse.TodoListCreateRequest{})

	// Check the result
	assert.NoError(t, err)

	// Assert that the expected function was called with the expected parameters
	mockHandler.AssertCalled(t, "Create", mock.Anything, mock.Anything)
}

func TestMockTodoListHandler_Delete(t *testing.T) {
	// Create a new instance of MockTodoListHandler
	mockHandler := NewMockTodoListHandler(t)

	// Set the expected return value
	mockHandler.On("Delete", mock.Anything, mock.Anything).Return(nil)

	// Call the function being tested
	err := mockHandler.Delete(nil, 1)

	// Check the result
	assert.NoError(t, err)

	// Assert that the expected function was called with the expected parameters
	mockHandler.AssertCalled(t, "Delete", mock.Anything, 1)
}

func TestMockTodoListHandler_ReadAll(t *testing.T) {
	// Create a new instance of MockTodoListHandler
	mockHandler := NewMockTodoListHandler(t)

	// Set the expected return value
	mockHandler.On("ReadAll", mock.Anything).Return(nil)

	// Call the function being tested
	err := mockHandler.ReadAll(nil)

	// Check the result
	assert.NoError(t, err)

	// Assert that the expected function was called with the expected parameters
	mockHandler.AssertCalled(t, "ReadAll", mock.Anything)
}

func TestMockTodoListHandler_ReadById(t *testing.T) {
	// Create a new instance of MockTodoListHandler
	mockHandler := NewMockTodoListHandler(t)

	// Set the expected return value
	mockHandler.On("ReadById", mock.Anything, mock.Anything).Return(nil)

	// Call the function being tested
	err := mockHandler.ReadById(nil, 1)

	// Check the result
	assert.NoError(t, err)

	// Assert that the expected function was called with the expected parameters
	mockHandler.AssertCalled(t, "ReadById", mock.Anything, 1)
}

func TestMockTodoListHandler_UpdateStatus(t *testing.T) {
	// Create a new instance of MockTodoListHandler
	mockHandler := NewMockTodoListHandler(t)

	// Set the expected return value
	mockHandler.On("UpdateStatus", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Call the function being tested
	err := mockHandler.UpdateStatus(nil, 1, requestAndresponse.TodoListUpdateStatus{})

	// Check the result
	assert.NoError(t, err)

	// Assert that the expected function was called with the expected parameters
	mockHandler.AssertCalled(t, "UpdateStatus", mock.Anything, 1, mock.Anything)
}

func TestMockTodoListHandler_UpdateTitleAndDescription(t *testing.T) {
	// Create a new instance of MockTodoListHandler
	mockHandler := NewMockTodoListHandler(t)

	// Set the expected return value
	mockHandler.On("UpdateTitleAndDescription", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Call the function being tested
	err := mockHandler.UpdateTitleAndDescription(nil, 1, requestAndresponse.TodoListUpdateTitleDescription{})

	// Check the result
	assert.NoError(t, err)

	// Assert that the expected function was called with the expected parameters
	mockHandler.AssertCalled(t, "UpdateTitleAndDescription", mock.Anything, 1, mock.Anything)
}
