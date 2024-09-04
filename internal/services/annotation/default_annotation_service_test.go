package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	customErroHandler "annotator-backend/pkg/errors"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockAnnotationDal struct {
	mock.Mock
}

func (m *MockAnnotationDal) AddAnnotation(annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	args := m.Called(annotation)
	return args.Get(0).(models.Annotation), args.Error(1)
}

func (m *MockAnnotationDal) GetAnnotations() ([]models.Annotation, error) {
	args := m.Called()
	return args.Get(0).([]models.Annotation), args.Error(1)
}

func (m *MockAnnotationDal) UpdateAnnotation(id int, annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	args := m.Called(id, annotation)
	return args.Get(0).(models.Annotation), args.Error(1)
}

func (m *MockAnnotationDal) DeleteAnnotation(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

type DummyMetadata struct {
	Comment string `json:"comment"`
}

func TestCreateAnnotation_should_return_created_object(t *testing.T) {
	mockDal := new(MockAnnotationDal)

	jsonData := createDummyMetadata()

	annotationService := NewDefaultAnnotationService(mockDal)
	request := &dtos.AnnotationRequest{Text: "العما", Metadata: jsonData}
	expectedAnnotation := models.Annotation{Id: 10, Text: "العما", Metadata: jsonData}

	mockDal.On("AddAnnotation", request).Return(expectedAnnotation, nil)

	result, err := annotationService.CreateAnnotation(request)

	assert.NoError(t, err)
	assert.Equal(t, expectedAnnotation, result)
	mockDal.AssertCalled(t, "AddAnnotation", request)
}

func TestCreateAnnotation_with_invalid_fields_should_return_incorrect_fields_error(t *testing.T) {
	mockDal := new(MockAnnotationDal)

	annotationService := NewDefaultAnnotationService(mockDal)
	request := &dtos.AnnotationRequest{Metadata: createDummyMetadata()}

	result, err := annotationService.CreateAnnotation(request)

	var incorrectFieldsError *customErroHandler.IncorrectFieldsError

	assert.True(t, errors.As(err, &incorrectFieldsError))
	assert.Equal(t, models.Annotation{}, result)
	mockDal.AssertNotCalled(t, "AddAnnotation")
}

func TestGetAnnotations_should_return_all_annotations(t *testing.T) {
	mockDal := new(MockAnnotationDal)

	annotationService := NewDefaultAnnotationService(mockDal)
	expectedAnnotations := []models.Annotation{
		{1, "بي", createDummyMetadata()},
		{2, "بي", createDummyMetadata()},
		{3, "بي", createDummyMetadata()},
	}

	mockDal.On("GetAnnotations").Return(expectedAnnotations, nil)

	result, err := annotationService.GetAnnotations()

	assert.NoError(t, err)
	assert.Equal(t, expectedAnnotations, result)
	mockDal.AssertExpectations(t)
}

func TestModifyAnnotation_should_returned_updated_object(t *testing.T) {
	mockDal := new(MockAnnotationDal)

	jsonData := createDummyMetadata()

	annotationService := NewDefaultAnnotationService(mockDal)
	request := &dtos.AnnotationRequest{Text: "تن", Metadata: jsonData}
	id := 1
	expectedAnnotation := models.Annotation{Id: id, Text: "تن", Metadata: jsonData}

	mockDal.On("UpdateAnnotation", id, request).Return(expectedAnnotation, nil)

	result, err := annotationService.ModifyAnnotation(id, request)

	assert.NoError(t, err)
	assert.Equal(t, expectedAnnotation, result)
	mockDal.AssertCalled(t, "UpdateAnnotation", 1, request)
}

func TestModifyAnnotation_with_invalid_fields_should_return_incorrect_fields_error(t *testing.T) {
	mockDal := new(MockAnnotationDal)

	annotationService := NewDefaultAnnotationService(mockDal)
	request := &dtos.AnnotationRequest{Metadata: createDummyMetadata()}

	result, err := annotationService.ModifyAnnotation(1, request)

	var incorrectFieldsError *customErroHandler.IncorrectFieldsError

	assert.True(t, errors.As(err, &incorrectFieldsError))
	assert.Equal(t, models.Annotation{}, result)
	mockDal.AssertNotCalled(t, "UpdateAnnotation")
}

func TestDeleteAnnotation_should_return_successful_response(t *testing.T) {
	mockDal := new(MockAnnotationDal)

	annotationService := NewDefaultAnnotationService(mockDal)

	id := 1
	mockDal.On("DeleteAnnotation", id).Return(nil)

	err := annotationService.DeleteAnnotation(id)

	assert.NoError(t, err)
	mockDal.AssertExpectations(t)
}

func createDummyMetadata() json.RawMessage {
	dummyMetadata := &DummyMetadata{Comment: "comment"}
	jsonData, _ := json.Marshal(dummyMetadata)
	return jsonData
}
