package utils

import (
	"math"

	"process-logs/types"

	"github.com/labstack/echo/v4"
)

func FormatTo2Decimal(value float64) float64 {
	return math.Ceil((value)*100) / 100
}

func SendResponse[T any](c echo.Context, statusCode int, response types.Response[T]) error {
	return c.JSON(statusCode, response)
}

// move to utils
func SuccessResponse[T any](data T) types.Response[T] {
	return types.Response[T]{
		Success: true,
		Data:    data,
		Message: "Operation successful",
	}
}

func ErrorResponse(message string) types.Response[interface{}] {
	return types.Response[interface{}]{
		Data:    nil,
		Success: false,
		Message: message,
	}
}

func FailureResponse[T any](c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, types.Response[interface{}]{
		Data:    nil,
		Success: false,
		Message: message,
	})
}
