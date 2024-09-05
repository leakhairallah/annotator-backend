package annotation

import (
	"annotator-backend/internal/dtos"
	"annotator-backend/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockAnnotationService struct {
	mock.Mock
}

func (m *MockAnnotationService) CreateAnnotation(annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	args := m.Called(annotation)
	return args.Get(0).(models.Annotation), args.Error(1)
}

func (m *MockAnnotationService) GetAnnotations() ([]models.Annotation, error) {
	args := m.Called()
	return args.Get(0).([]models.Annotation), args.Error(1)
}

func (m *MockAnnotationService) ModifyAnnotation(id int, annotation *dtos.AnnotationRequest) (models.Annotation, error) {
	args := m.Called(id, annotation)
	return args.Get(0).(models.Annotation), args.Error(1)
}

func (m *MockAnnotationService) DeleteAnnotation(id int) (dtos.DeleteAnnotationResponse, error) {
	args := m.Called(id)
	return args.Get(0).(dtos.DeleteAnnotationResponse), args.Error(1)
}

//func setupEchoContext(t *testing.T, method, url string, body interface{}) (echo.Context, *httptest.ResponseRecorder) {
//	buf, err := converter.AnyToBytesBuffer(body)
//	require.NoError(t, err)
//	require.NotNil(t, buf)
//	require.Nil(t, err)
//
//	req := httptest.NewRequest(method, url, strings.NewReader(buf.String()))
//	rec := httptest.NewRecorder()
//	e := echo.New()
//	c := e.NewContext(req, rec)
//	return c, rec
//}
//
//var annotationH Handlers
//var mockAnnotationService = new(MockAnnotationService)
//
//func init() {
//	annotationH = NewAnnotationHandlers(mockAnnotationService)
//}
//
//func TestCreateAnnotation_should_return_200(t *testing.T) {
//	annotationRequest := &dtos.AnnotationRequest{Text: "هاها", Metadata: utils.CreateDummyMetadata()}
//	annotationResponse := &models.Annotation{Id: 1, Text: "هاها", Metadata: utils.CreateDummyMetadata()}
//
//	mockAnnotationService.On("CreateAnnotation", annotationRequest).Return(annotationResponse, nil)
//
//	c, rec := setupEchoContext(t, http.MethodPost, "/annotations", annotationRequest)
//
//	if assert.NoError(t, annotationH.Create()(c)) {
//		assert.Equal(t, http.StatusCreated, rec.Code)
//		assert.JSONEq(t, "", rec.Body.String())
//	}
//
//	mockAnnotationService.AssertExpectations(t)
//
//}

//func TestGetAllAnnotations(t *testing.T) {
//	mockService := new(MockAnnotationService)
//	handlers := NewAnnotationHandlers(mockService)
//	annotations := []*dtos.AnnotationResponse{ /* fill with test data */ }
//
//	mockService.On("GetAnnotations").Return(annotations, nil)
//
//	c, rec := setupEchoContext(http.MethodGet, "/annotations", nil)
//
//	if assert.NoError(t, handlers.GetAll()(c)) {
//		assert.Equal(t, http.StatusOK, rec.Code)
//		assert.JSONEq(t, `/* expected JSON response */`, rec.Body.String())
//	}
//
//	mockService.AssertExpectations(t)
//}
//
//func TestUpdateAnnotation(t *testing.T) {
//	mockService := new(MockAnnotationService)
//	handlers := NewAnnotationHandlers(mockService)
//	annotationRequest := &dtos.AnnotationRequest{ /* fill with test data */ }
//	annotationResponse := &dtos.AnnotationResponse{ /* fill with expected response data */ }
//	id := 1
//
//	mockService.On("ModifyAnnotation", id, annotationRequest).Return(annotationResponse, nil)
//
//	c, rec := setupEchoContext(http.MethodPut, "/annotations/"+strconv.Itoa(id), annotationRequest)
//	c.SetParamNames("id")
//	c.SetParamValues(strconv.Itoa(id))
//
//	if assert.NoError(t, handlers.Update()(c)) {
//		assert.Equal(t, http.StatusOK, rec.Code)
//		assert.JSONEq(t, `/* expected JSON response */`, rec.Body.String())
//	}
//
//	mockService.AssertExpectations(t)
//}
//
//func TestDeleteAnnotation(t *testing.T) {
//	mockService := new(MockAnnotationService)
//	handlers := NewAnnotationHandlers(mockService)
//	id := 1
//	deleteResponse := &dtos.DeleteResponse{ /* fill with expected response data */ }
//
//	mockService.On("DeleteAnnotation", id).Return(deleteResponse, nil)
//
//	c, rec := setupEchoContext(http.MethodDelete, "/annotations/"+strconv.Itoa(id), nil)
//	c.SetParamNames("id")
//	c.SetParamValues(strconv.Itoa(id))
//
//	if assert.NoError(t, handlers.Delete()(c)) {
//		assert.Equal(t, http.StatusOK, rec.Code)
//		assert.JSONEq(t, `/* expected JSON response */`, rec.Body.String())
//	}
//
//	mockService.AssertExpectations(t)
//}
//
//func TestCreateAnnotation_BindError(t *testing.T) {
//	mockService := new(MockAnnotationService)
//	handlers := NewAnnotationHandlers(mockService)
//
//	c, rec := setupEchoContext(http.MethodPost, "/annotations", "invalid json")
//
//	if assert.Error(t, handlers.Create()(c)) {
//		assert.Equal(t, http.StatusBadRequest, rec.Code)
//	}
//
//	mockService.AssertNotCalled(t, "CreateAnnotation")
//}
//
//func TestUpdateAnnotation_IdConversionError(t *testing.T) {
//	mockService := new(MockAnnotationService)
//	handlers := NewAnnotationHandlers(mockService)
//
//	c, rec := setupEchoContext(http.MethodPut, "/annotations/abc", nil)
//	c.SetParamNames("id")
//	c.SetParamValues("abc")
//
//	if assert.Error(t, handlers.Update()(c)) {
//		assert.Equal(t, http.StatusNotFound, rec.Code)
//	}
//
//	mockService.AssertNotCalled(t, "ModifyAnnotation")
//}
//
//func TestDeleteAnnotation_IdConversionError(t *testing.T) {
//	mockService := new(MockAnnotationService)
//	handlers := NewAnnotationHandlers(mockService)
//
//	c, rec := setupEchoContext(http.MethodDelete, "/annotations/abc", nil)
//	c.SetParamNames("id")
//	c.SetParamValues("abc")
//
//	if assert.Error(t, handlers.Delete()(c)) {
//		assert.Equal(t, http.StatusNotFound, rec.Code)
//	}
//
//	mockService.AssertNotCalled(t, "DeleteAnnotation")
//}
