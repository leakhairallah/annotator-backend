package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/services/annotation"
	customErrorhandler "annotator-backend/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type annotationHandlers struct {
	annotationService annotation.AnnotationService
}

func NewAnnotationHandlers(annotationService annotation.AnnotationService) Handlers {
	return &annotationHandlers{annotationService: annotationService}
}

func (a annotationHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var annotationFromContext dtos.Annotation
		if err := c.Bind(&annotationFromContext); err != nil {
			log.Errorf(customErrorhandler.BuildRequestFailedErrorMessage(c.Request().URL.String(), http.StatusBadRequest))
			return echo.NewHTTPError(http.StatusBadRequest, customErrorhandler.BadRequest)
		}

		createdAnnotation, err := a.annotationService.CreateAnnotation(&annotationFromContext)
		if err != nil {
			handledError := customErrorhandler.HandleCustomError(err)
			log.Errorf(customErrorhandler.BuildRequestFailedErrorMessage(c.Request().URL.String(), handledError.Code))
			return handledError
		}

		return c.JSON(http.StatusCreated, createdAnnotation)
	}
}

func (a annotationHandlers) GetAll() echo.HandlerFunc {
	//TODO implement me
	return nil
}

func (a annotationHandlers) Update() echo.HandlerFunc {
	//TODO implement me
	return nil
}

func (a annotationHandlers) Delete() echo.HandlerFunc {
	//TODO implement me
	return nil
}
