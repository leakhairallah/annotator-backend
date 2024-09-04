package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	"annotator-backend/internal/repository/annotation"
	"github.com/go-playground/validator/v10"
)

type AnnotationService interface {
	CreateAnnotation(annotation *dtos.AnnotationRequest) (models.Annotation, error)
	GetAnnotations() ([]models.Annotation, error)
	ModifyAnnotation(id int, annotation *dtos.AnnotationRequest) (models.Annotation, error)
	DeleteAnnotation(id int) error
}

func NewDefaultAnnotationService(annotationDal annotation.AnnotationDal) AnnotationService {
	return &defaultAnnotationService{annotationDal: annotationDal, validator: validator.New()}
}
