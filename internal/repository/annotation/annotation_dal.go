package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	"database/sql"
)

type AnnotationDal interface {
	AddAnnotation(annotation *dtos.Annotation) (models.Annotation, error)
	GetAnnotations() ([]models.Annotation, error)
	UpdateAnnotation(annotation *models.Annotation) (models.Annotation, error)
	DeleteAnnotation(id int)
}

func NewMySqlAnnotationDal(db *sql.DB) AnnotationDal {
	return &MySqlAnnotationDal{conn: db}
}
