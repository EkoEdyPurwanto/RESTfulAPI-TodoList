package utils

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
)

func WriteToResponseBody(ctx echo.Context, response interface{}) {

	encoder := json.NewEncoder(ctx.Response())
	err := encoder.Encode(response)
	PanicIfError(err)
}
