package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	"annotator-backend/pkg/errors"
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	"strconv"
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
	}

	log.Infof("Successfully added annotation with id %d", lastInsertID)
	return models.Annotation{Id: lastInsertID, Text: annotation.Text, Metadata: annotation.Metadata}, nil
}

func (db MySqlAnnotationDal) GetAnnotations() ([]models.Annotation, error) {
	query := fmt.Sprintf("SELECT * FROM %s", TableName)
	stmtIns, err := db.conn.Prepare(query)
	if err != nil {
		log.Error(err.Error())
		return nil, &errors.DatabaseError{CustomError: &errors.CustomError{Message: errors.InternalServererror}}
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(stmtIns)

	result, err := stmtIns.Query()
	if err != nil {
		log.Error(err.Error())
		return nil, &errors.DatabaseError{CustomError: &errors.CustomError{Message: errors.InternalServererror}}
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(result)

	var annotations []models.Annotation
	for result.Next() {
		var a models.Annotation
		err := result.Scan(&a.Id, &a.Text, &a.Metadata)
		if err != nil {
			log.Error(err.Error())
			return nil, &errors.DatabaseError{CustomError: &errors.CustomError{Message: errors.InternalServererror}}
		}
		annotations = append(annotations, a)
	}

	if err := result.Err(); err != nil {
		log.Error(err.Error())
		return nil, &errors.DatabaseError{CustomError: &errors.CustomError{Message: errors.InternalServererror}}
	}

	log.Infof("Fetched all annotations successfully")
	return annotations, nil
}

func (db MySqlAnnotationDal) UpdateAnnotation(id int, annotation *dtos.Annotation) error {
	query := fmt.Sprintf("UPDATE %s SET text = ?, metadata = ? WHERE id = ?", TableName)
	stmtIns, err := db.conn.Prepare(query)
	if err != nil {
		log.Error(err.Error())
		return &errors.CustomError{Message: errors.InternalServererror}
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(stmtIns)

	result, err := stmtIns.Exec(annotation.Text, annotation.Metadata, id)
	if err != nil {
		log.Error(err.Error())
		return &errors.DatabaseError{CustomError: &errors.CustomError{}}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Warn(err.Error())
	}

	if rowsAffected == 0 {
		log.Warn("No rows updated, ID may not exist")
		return &errors.IdNotFound{CustomError: &errors.CustomError{Message: errors.RequestNotFound(strconv.Itoa(id))}}
	}

	log.Infof("Successfully updated annotation with id %d", id)
	return nil
}

func (db MySqlAnnotationDal) DeleteAnnotation(id int) {

}
