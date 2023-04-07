package utils

import (
	"LearnECHO/models/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UnauthorizedError(err error, ctx echo.Context) {
	response := domain.Response{
		Status:  http.StatusUnauthorized,
		Message: err.Error(),
	}
	ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	ctx.Response().WriteHeader(response.Status)
	WriteToResponseBody(ctx, response)
}
