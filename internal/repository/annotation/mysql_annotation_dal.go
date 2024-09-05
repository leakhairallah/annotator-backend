package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	customErrorHandler "annotator-backend/pkg/errors"
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"strconv"
)

//TODO Research on a better ORM

const (
	TableName          = "annotations"
	TextColumnName     = "text"
	MetadataColumnName = "metadata"
)

type mySqlAnnotationDal struct {
	conn *sql.DB
}

func (db mySqlAnnotationDal) AddAnnotation(annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES(?, ?)", TableName, TextColumnName, MetadataColumnName)
	stmt, err := db.conn.Prepare(query)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &customErrorHandler.DatabaseError{CustomError: &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}}
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(stmt)

	result, err := stmt.Exec(annotation.Text, annotation.Metadata)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &customErrorHandler.DatabaseError{CustomError: &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}}
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Warn(err.Error())
	}

	log.Infof("Successfully added annotation with id %d", lastInsertID)
	return models.Annotation{Id: int(lastInsertID), Text: annotation.Text, Metadata: annotation.Metadata}, nil
}

func (db mySqlAnnotationDal) GetAnnotations() ([]models.Annotation, error) {
	query := fmt.Sprintf("SELECT * FROM %s", TableName)
	stmt, err := db.conn.Prepare(query)
	if err != nil {
		log.Error(err.Error())
		return nil, &customErrorHandler.DatabaseError{CustomError: &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}}
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(stmt)

	result, err := stmt.Query()
	if err != nil {
		log.Error(err.Error())
		return nil, &customErrorHandler.DatabaseError{CustomError: &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}}
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
			return nil, &customErrorHandler.DatabaseError{CustomError: &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}}
		}
		annotations = append(annotations, a)
	}

	if err := result.Err(); err != nil {
		log.Error(err.Error())
		return nil, &customErrorHandler.DatabaseError{CustomError: &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}}
	}

	log.Infof("Fetched all annotations successfully")
	return annotations, nil
}

func (db mySqlAnnotationDal) UpdateAnnotation(id int, annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	query := fmt.Sprintf("UPDATE %s SET text = ?, metadata = ? WHERE id = ?", TableName)
	stmt, err := db.conn.Prepare(query)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(stmt)

	result, err := stmt.Exec(annotation.Text, annotation.Metadata, id)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, &customErrorHandler.DatabaseError{CustomError: &customErrorHandler.CustomError{}}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Warn(err.Error())
	}

	if rowsAffected == 0 {
		log.Warnf("No rows updated, ID %d may not exist", id)
		return models.Annotation{}, &customErrorHandler.IdNotFound{CustomError: &customErrorHandler.CustomError{Message: customErrorHandler.RequestNotFound(strconv.Itoa(id))}}
	}

	log.Infof("Successfully updated annotation with id %d", id)
	updatedAnnotations, err := db.getAnnotationById(id)
	if err != nil {
		log.Warn(err.Error())
		return models.Annotation{Id: id, Text: annotation.Text, Metadata: annotation.Metadata}, nil
	}

	return updatedAnnotations, nil
}

func (db mySqlAnnotationDal) DeleteAnnotation(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TableName)
	stmt, err := db.conn.Prepare(query)
	if err != nil {
		log.Error(err.Error())
		return &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(stmt)

	result, err := stmt.Exec(id)
	if err != nil {
		log.Error(err.Error())
		return &customErrorHandler.CustomError{Message: customErrorHandler.InternalServererror}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Warn(err.Error())
	}

	if rowsAffected == 0 {
		log.Warnf("No rows updated, ID %d may not exist", id)
		return &customErrorHandler.IdNotFound{CustomError: &customErrorHandler.CustomError{Message: customErrorHandler.RequestNotFound(strconv.Itoa(id))}}
	}

	log.Infof("Successfully deleted annotation with id %d", id)
	return nil
}

func (db mySqlAnnotationDal) getAnnotationById(id int) (models.Annotation, error) {
	query := fmt.Sprintf("SELECT id, text, metadata FROM %s WHERE id = ?", TableName)
	stmt, err := db.conn.Prepare(query)
	if err != nil {
		log.Error(err.Error())
		return models.Annotation{}, err
	}
	defer func(stmtSel *sql.Stmt) {
		err := stmtSel.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}(stmt)

	var updatedAnnotation models.Annotation
	err = stmt.QueryRow(id).Scan(&updatedAnnotation.Id, &updatedAnnotation.Text, &updatedAnnotation.Metadata)
	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return models.Annotation{}, err
		}
		return models.Annotation{}, err
	}

	return updatedAnnotation, nil
}
