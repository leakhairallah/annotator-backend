package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	"database/sql"
)

type AnnotationDal interface {
	AddAnnotation(annotation *dtos.AnnotationRequest) (models.Annotation, error)
	GetAnnotations() ([]models.Annotation, error)
	UpdateAnnotation(id int, annotation *dtos.AnnotationRequest) (models.Annotation, error)
	DeleteAnnotation(id int) error
}

func NewMySqlAnnotationDal(db *sql.DB) AnnotationDal {
	return &MySqlAnnotationDal{conn: db}
}
