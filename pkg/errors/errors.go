package errors

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	BadRequest          = "Invalid request payload"
	InternalServererror = "Internal server error"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf(e.Message)
}

type IncorrectFieldsError struct {
	*CustomError
}

type DatabaseError struct {
	*CustomError
}

func HandleCustomError(err error) *echo.HTTPError {
	var incorrectFieldsError *IncorrectFieldsError
	var databaseError *DatabaseError

	if errors.As(err, &incorrectFieldsError) {
		return echo.NewHTTPError(http.StatusBadRequest, BadRequest)
	} else if errors.As(err, &databaseError) {
		return echo.NewHTTPError(http.StatusInternalServerError, InternalServererror)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, InternalServererror)
}

func BuildRequestFailedMessage(request string, statusCode int) string {
	return fmt.Sprintf("Request %s failed with status code: %d", request, statusCode)
}

func BuildRequestSucceededMessage(request string, statusCode int) string {
	return fmt.Sprintf("Request %s succeeded with status code: %d", request, statusCode)
}

func BuildIncorrectFieldsMessage(error error) string {
	var message string
	for _, err := range error.(validator.ValidationErrors) {
		message += fmt.Sprintf("Field '%s'='%s' failed validation with tag '%s'.", err.Field(), err.Value(), err.Tag())
	}
	return message
}
