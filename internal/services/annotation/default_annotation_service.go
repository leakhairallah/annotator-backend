package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	"annotator-backend/internal/repository/annotation"
	"annotator-backend/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type DefaultAnnotationService struct {
	annotationDal annotation.AnnotationDal
	validator     *validator.Validate
}

func (annotationService DefaultAnnotationService) CreateAnnotation(annotation *dtos.Annotation) (models.Annotation, error) {
	err := annotationService.validator.Struct(annotation)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &errors.IncorrectFieldsError{CustomError: &errors.CustomError{Message: errors.BuildIncorrectFieldsMessage(err)}}
	}
	return annotationService.annotationDal.AddAnnotation(annotation)
}

func (annotationService DefaultAnnotationService) GetAnnotations() ([]models.Annotation, error) {
	return annotationService.annotationDal.GetAnnotations()
}

func (annotationService DefaultAnnotationService) ModifyAnnotation(id int, annotation *dtos.Annotation) error {
	err := annotationService.validator.Struct(annotation)
	if err != nil {
		log.Error(err.Error())
		return &errors.IncorrectFieldsError{CustomError: &errors.CustomError{Message: errors.BuildIncorrectFieldsMessage(err)}}
	}
	return annotationService.annotationDal.UpdateAnnotation(id, annotation)
}

func (annotationService DefaultAnnotationService) DeleteAnnotation(id int) error {
	return annotationService.annotationDal.DeleteAnnotation(id)
}
