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
	requestNotFound     = "The request resource with ID %s was not found"
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

type IdNotFound struct {
	*CustomError
}

func HandleCustomError(err error) *echo.HTTPError {
	var incorrectFieldsError *IncorrectFieldsError
	var databaseError *DatabaseError
	var idNotFoundError *IdNotFound

	if errors.As(err, &incorrectFieldsError) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if errors.As(err, &databaseError) {
		return echo.NewHTTPError(http.StatusInternalServerError, InternalServererror)
	} else if errors.As(err, &idNotFoundError) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return echo.NewHTTPError(http.StatusInternalServerError, InternalServererror)
}

func RequestNotFound(id string) string {
	return fmt.Sprintf(requestNotFound, id)
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
		message += fmt.Sprintf("Input '%s' provided for field '%s' is invalid. Field '%s' failed on the following tag '%s'.", err.Value(), err.Field(), err.Field(), err.Tag())
	}
	return message
}
