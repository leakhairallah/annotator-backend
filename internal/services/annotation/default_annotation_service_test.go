package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	customErroHandler "annotator-backend/pkg/errors"
	"annotator-backend/pkg/utils"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockAnnotationDal struct {
	mock.Mock
}

func (m *mockAnnotationDal) AddAnnotation(annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	args := m.Called(annotation)
	return args.Get(0).(models.Annotation), args.Error(1)
}

func (m *mockAnnotationDal) GetAnnotations() ([]models.Annotation, error) {
	args := m.Called()
	return args.Get(0).([]models.Annotation), args.Error(1)
}

func (m *mockAnnotationDal) UpdateAnnotation(id int, annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	args := m.Called(id, annotation)
	return args.Get(0).(models.Annotation), args.Error(1)
}

func (m *mockAnnotationDal) DeleteAnnotation(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateAnnotation_should_return_created_object(t *testing.T) {
	jsonData := utils.CreateDummyMetadata()

	mockDal := new(mockAnnotationDal)
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
	request := &dtos.AnnotationRequest{Metadata: utils.CreateDummyMetadata()}

	mockDal := new(mockAnnotationDal)
	annotationService := NewDefaultAnnotationService(mockDal)

	result, err := annotationService.CreateAnnotation(request)

	var incorrectFieldsError *customErroHandler.IncorrectFieldsError

	assert.True(t, errors.As(err, &incorrectFieldsError))
	assert.Equal(t, models.Annotation{}, result)
	mockDal.AssertNotCalled(t, "AddAnnotation")
}

func TestGetAnnotations_should_return_all_annotations(t *testing.T) {
	expectedAnnotations := []models.Annotation{
		{1, "بي", utils.CreateDummyMetadata()},
		{2, "بي", utils.CreateDummyMetadata()},
		{3, "بي", utils.CreateDummyMetadata()},
	}

	mockDal := new(mockAnnotationDal)
	annotationService := NewDefaultAnnotationService(mockDal)

	mockDal.On("GetAnnotations").Return(expectedAnnotations, nil)

	result, err := annotationService.GetAnnotations()

	assert.NoError(t, err)
	assert.Equal(t, expectedAnnotations, result)
	mockDal.AssertExpectations(t)
}

func TestModifyAnnotation_should_returned_updated_object(t *testing.T) {
	mockDal := new(mockAnnotationDal)
	annotationService := NewDefaultAnnotationService(mockDal)

	jsonData := utils.CreateDummyMetadata()
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
	mockDal := new(mockAnnotationDal)
	annotationService := NewDefaultAnnotationService(mockDal)

	request := &dtos.AnnotationRequest{Metadata: utils.CreateDummyMetadata()}

	result, err := annotationService.ModifyAnnotation(1, request)

	var incorrectFieldsError *customErroHandler.IncorrectFieldsError

	assert.True(t, errors.As(err, &incorrectFieldsError))
	assert.Equal(t, models.Annotation{}, result)
	mockDal.AssertNotCalled(t, "UpdateAnnotation")
}

func TestDeleteAnnotation_should_return_successful_response(t *testing.T) {
	mockDal := new(mockAnnotationDal)
	annotationService := NewDefaultAnnotationService(mockDal)

	expectedResponse := &dtos.DeleteAnnotationResponse{Success: true}
	id := 1
	mockDal.On("DeleteAnnotation", id).Return(nil)

	response, err := annotationService.DeleteAnnotation(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, &response)
	mockDal.AssertCalled(t, "DeleteAnnotation", id)
}

func TestDeleteAnnotation_should_return_unsuccessful_response(t *testing.T) {
	mockDal := new(mockAnnotationDal)
	annotationService := NewDefaultAnnotationService(mockDal)

	expectedResponse := &dtos.DeleteAnnotationResponse{Success: false}
	id := 1
	mockDal.On("DeleteAnnotation", id).Return(errors.New("error"))

	response, err := annotationService.DeleteAnnotation(id)

	assert.Error(t, err)
	assert.Equal(t, expectedResponse, &response)
	mockDal.AssertCalled(t, "DeleteAnnotation", id)
}
