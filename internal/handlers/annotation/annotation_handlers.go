package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/services/annotation"
	customErrorhandler "annotator-backend/pkg/errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
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
			log.Error(customErrorhandler.BuildRequestFailedMessage(c.Request().URL.String(), http.StatusBadRequest))
			return echo.NewHTTPError(http.StatusBadRequest, customErrorhandler.BadRequest)
		}

		createdAnnotation, err := a.annotationService.CreateAnnotation(&annotationFromContext)
		if err != nil {
			handledError := customErrorhandler.HandleCustomError(err)
			log.Error(customErrorhandler.BuildRequestFailedMessage(c.Request().URL.String(), handledError.Code))
			return handledError
		}

		log.Info(customErrorhandler.BuildRequestSucceededMessage(c.Request().URL.String(), http.StatusCreated))
		return c.JSON(http.StatusCreated, createdAnnotation)
	}
}

func (a annotationHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		retrievedAnnotations, err := a.annotationService.GetAnnotations()
		if err != nil {
			handledError := customErrorhandler.HandleCustomError(err)
			log.Error(customErrorhandler.BuildRequestFailedMessage(c.Request().URL.String(), handledError.Code))
			return handledError
		}

		log.Info(customErrorhandler.BuildRequestSucceededMessage(c.Request().URL.String(), http.StatusOK))
		return c.JSON(http.StatusOK, retrievedAnnotations)
	}
}

func (a annotationHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var annotationFromContext dtos.Annotation
		if err := c.Bind(&annotationFromContext); err != nil {
			log.Error(customErrorhandler.BuildRequestFailedMessage(c.Request().URL.String(), http.StatusBadRequest))
			return echo.NewHTTPError(http.StatusBadRequest, customErrorhandler.BadRequest)
		}

		annotationId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error(customErrorhandler.BuildRequestFailedMessage(c.Request().URL.String(), http.StatusNotFound))
			return echo.NewHTTPError(http.StatusNotFound, customErrorhandler.RequestNotFound(c.Param("id")))
		}

		err = a.annotationService.ModifyAnnotation(annotationId, &annotationFromContext)
		if err != nil {
			handledError := customErrorhandler.HandleCustomError(err)
			log.Error(customErrorhandler.BuildRequestFailedMessage(c.Request().URL.String(), handledError.Code))
			return handledError
		}

		log.Info(customErrorhandler.BuildRequestSucceededMessage(c.Request().URL.String(), http.StatusOK))
		return c.JSON(http.StatusOK, fmt.Sprintf("Successfully updated annotation with id %d", annotationId))
	}
}

func (a annotationHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		annotationId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error(customErrorhandler.BuildRequestFailedMessage(c.Request().URL.String(), http.StatusNotFound))
			return echo.NewHTTPError(http.StatusNotFound, customErrorhandler.RequestNotFound(c.Param("id")))
		}

		err = a.annotationService.DeleteAnnotation(annotationId)
		if err != nil {
			handledError := customErrorhandler.HandleCustomError(err)
			log.Error(customErrorhandler.BuildRequestFailedMessage(c.Request().URL.String(), handledError.Code))
			return handledError
		}

		log.Info(customErrorhandler.BuildRequestSucceededMessage(c.Request().URL.String(), http.StatusOK))
		return c.JSON(http.StatusOK, fmt.Sprintf("Successfully deleted annotation with id %d", annotationId))
	}
}
