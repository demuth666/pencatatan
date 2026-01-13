package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pencatatan/internal/models"
	"pencatatan/internal/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter(handler *SaleHandler) *gin.Engine {
	router := gin.New()
	router.POST("/sales", handler.CreateSale)
	router.GET("/sales", handler.GetAllSales)
	router.GET("/sales/:id", handler.GetSaleByID)
	router.PUT("/sales/:id", handler.UpdateSale)
	router.DELETE("/sales/:id", handler.DeleteSale)
	return router
}

func TestCreateSale_Success(t *testing.T) {
	mockService := &service.MockSaleService{
		CreateSaleFunc: func(req *models.CreateSalesRequest) (*models.Sale, error) {
			return &models.Sale{
				ID:             uuid.New(),
				Product:        req.Product,
				Quantity:       req.Quantity,
				Price:          req.Price,
				Total:          float64(req.Quantity) * req.Price,
				AmountReceived: req.AmountReceived,
				IsDebt:         req.IsDebt,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}, nil
		},
	}

	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	reqBody := models.CreateSalesRequest{
		Product:        "Test Product",
		Quantity:       2,
		Price:          10000,
		AmountReceived: 25000,
		IsDebt:         false,
	}

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/sales", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)

	if !response.Success {
		t.Error("Expected success to be true")
	}
}

func TestCreateSale_BadRequest(t *testing.T) {
	mockService := &service.MockSaleService{}
	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("POST", "/sales", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Success {
		t.Error("Expected success to be false for bad request")
	}
}

func TestCreateSale_ServiceError(t *testing.T) {
	mockService := &service.MockSaleService{
		CreateSaleFunc: func(req *models.CreateSalesRequest) (*models.Sale, error) {
			return nil, errors.New("service error")
		},
	}

	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	reqBody := models.CreateSalesRequest{
		Product:        "Test Product",
		Quantity:       2,
		Price:          10000,
		AmountReceived: 5000, // Insufficient amount
		IsDebt:         false,
	}

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/sales", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestGetSaleByID_Success(t *testing.T) {
	expectedID := uuid.New()
	mockService := &service.MockSaleService{
		GetSaleByIDFunc: func(id string) (*models.Sale, error) {
			return &models.Sale{
				ID:       expectedID,
				Product:  "Test Product",
				Quantity: 1,
				Price:    5000,
			}, nil
		},
	}

	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/sales/"+expectedID.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)

	if !response.Success {
		t.Error("Expected success to be true")
	}
}

func TestGetSaleByID_NotFound(t *testing.T) {
	mockService := &service.MockSaleService{
		GetSaleByIDFunc: func(id string) (*models.Sale, error) {
			return nil, errors.New("sale not found")
		},
	}

	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/sales/"+uuid.New().String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestGetAllSales_Success(t *testing.T) {
	mockService := &service.MockSaleService{
		GetAllSalesFunc: func() ([]*models.Sale, error) {
			return []*models.Sale{
				{ID: uuid.New(), Product: "Product 1"},
				{ID: uuid.New(), Product: "Product 2"},
			}, nil
		},
	}

	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/sales", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)

	if !response.Success {
		t.Error("Expected success to be true")
	}
}

func TestUpdateSale_Success(t *testing.T) {
	expectedID := uuid.New()
	mockService := &service.MockSaleService{
		UpdateSalesFunc: func(id string, req *models.UpdateSaleRequest) (*models.Sale, error) {
			return &models.Sale{
				ID:       expectedID,
				Product:  req.Product,
				Quantity: req.Quantity,
				Price:    req.Price,
			}, nil
		},
	}

	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	reqBody := models.UpdateSaleRequest{
		Product:  "Updated Product",
		Quantity: 5,
		Price:    15000,
	}

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/sales/"+expectedID.String(), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)

	if !response.Success {
		t.Error("Expected success to be true")
	}
}

func TestUpdateSale_BadRequest(t *testing.T) {
	mockService := &service.MockSaleService{}
	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("PUT", "/sales/"+uuid.New().String(), bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestDeleteSale_Success(t *testing.T) {
	mockService := &service.MockSaleService{
		DeleteSalesFunc: func(id string) error {
			return nil
		},
	}

	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("DELETE", "/sales/"+uuid.New().String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response Response
	json.Unmarshal(w.Body.Bytes(), &response)

	if !response.Success {
		t.Error("Expected success to be true")
	}
}

func TestDeleteSale_NotFound(t *testing.T) {
	mockService := &service.MockSaleService{
		DeleteSalesFunc: func(id string) error {
			return errors.New("sale not found")
		},
	}

	handler := NewSaleHandler(mockService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("DELETE", "/sales/"+uuid.New().String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}
