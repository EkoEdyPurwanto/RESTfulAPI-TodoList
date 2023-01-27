package helper

import (
	"LearnECHO/models/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InternalServerError(err error, ctx echo.Context) {
	ctx.Response().Header().Add("Content-Type", "application/json")
	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := domain.Response{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
		Data:    err.Error(),
	}
	ctx.Response().WriteHeader(apiResponse.Status)
	WriteToResponseBody(ctx, apiResponse)
}

func BadRequest(err error, ctx echo.Context) {
	ctx.Response().Header().Add("Content-Type", "application/json")
	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := domain.Response{
		Status:  http.StatusBadRequest,
		Message: "bad request",
		Data:    err.Error(),
	}
	ctx.Response().WriteHeader(apiResponse.Status)
	WriteToResponseBody(ctx, apiResponse)
}

func NotFound(err error, ctx echo.Context) {
	ctx.Response().Header().Add("Content-Type", "application/json")
	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := domain.Response{
		Status:  http.StatusNotFound,
		Message: "not found",
		Data:    err.Error(),
	}
	ctx.Response().WriteHeader(apiResponse.Status)
	WriteToResponseBody(ctx, apiResponse)
}
