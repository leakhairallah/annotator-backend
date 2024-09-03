package app

import (
	handlers "annotator-backend/internal/handlers/annotation"
	repo "annotator-backend/internal/repository/annotation"
	service "annotator-backend/internal/services/annotation"
	"github.com/labstack/echo/v4"
)

func (annotatorApp *AnnotatorApp) MapHandlers(e *echo.Echo) error {
	annotationRepo := repo.NewMySqlAnnotationDal(annotatorApp.db)
	annotationService := service.NewDefaultAnnotationService(&annotationRepo)
	annotationHandler := handlers.NewAnnotationHandlers(annotationService)

	v1 := e.Group("/api/v1")
	annotationGroup := v1.Group("/annotations")
	handlers.MapAnnotationRoutes(annotationGroup, annotationHandler)
	return nil
}
