package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	"annotator-backend/pkg/errors"
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
)

const (
	TableName          = "annotations"
	TextColumnName     = "text"
	MetadataColumnName = "metadata"
)

type MySqlAnnotationDal struct {
	conn *sql.DB
}

func (db MySqlAnnotationDal) AddAnnotation(annotation *dtos.Annotation) (models.Annotation, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES(?, ?)", TableName, TextColumnName, MetadataColumnName)
	stmtIns, err := db.conn.Prepare(query)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &errors.DatabaseError{CustomError: &errors.CustomError{Message: errors.InternalServererror}}
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(stmtIns)

	result, err := stmtIns.Exec(annotation.Text, annotation.Metadata)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &errors.DatabaseError{CustomError: &errors.CustomError{Message: errors.InternalServererror}}
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Warn(err.Error())
		return models.Annotation{}, nil
	}
	log.Infof("Successfully added annotation with id %d", lastInsertID)

	return models.Annotation{Id: lastInsertID, Text: annotation.Text, Metadata: annotation.Metadata}, nil
}

func (db MySqlAnnotationDal) GetAnnotations() ([]models.Annotation, error) {
	return nil, nil
}

func (db MySqlAnnotationDal) UpdateAnnotation(annotation *models.Annotation) (models.Annotation, error) {
	return models.Annotation{}, nil
}

func (db MySqlAnnotationDal) DeleteAnnotation(id int) {

}
