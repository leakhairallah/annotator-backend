package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	"annotator-backend/internal/repository/annotation"
	"annotator-backend/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type defaultAnnotationService struct {
	annotationDal annotation.AnnotationDal
	validator     *validator.Validate
}

func (annotationService defaultAnnotationService) CreateAnnotation(annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	err := annotationService.validator.Struct(annotation)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &errors.IncorrectFieldsError{CustomError: &errors.CustomError{Message: errors.BuildIncorrectFieldsMessage(err)}}
	}
	return annotationService.annotationDal.AddAnnotation(annotation)
}

func (annotationService defaultAnnotationService) GetAnnotations() ([]models.Annotation, error) {
	return annotationService.annotationDal.GetAnnotations()
}

func (annotationService defaultAnnotationService) ModifyAnnotation(id int, annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	err := annotationService.validator.Struct(annotation)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &errors.IncorrectFieldsError{CustomError: &errors.CustomError{Message: errors.BuildIncorrectFieldsMessage(err)}}
	}
	return annotationService.annotationDal.UpdateAnnotation(id, annotation)
}

func (annotationService defaultAnnotationService) DeleteAnnotation(id int) (dtos.DeleteAnnotationResponse, error) {
	err := annotationService.annotationDal.DeleteAnnotation(id)
	if err != nil {
		return dtos.DeleteAnnotationResponse{}, err
	}
	return dtos.DeleteAnnotationResponse{Success: true}, nil
}
