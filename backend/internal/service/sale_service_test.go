package service

import (
	"errors"
	"pencatatan/internal/models"
	"pencatatan/internal/repository"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateSale_Success(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{
		CreateFunc: func(req *models.CreateSalesRequest) (*models.Sale, error) {
			return &models.Sale{
				ID:             uuid.New(),
				Product:        req.Product,
				Quantity:       req.Quantity,
				Price:          req.Price,
				Total:          float64(req.Quantity) * req.Price,
				AmountReceived: req.AmountReceived,
				ChangeAmount:   req.AmountReceived - (float64(req.Quantity) * req.Price),
				IsDebt:         req.IsDebt,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}, nil
		},
	}

	service := NewSaleService(mockRepo)

	req := &models.CreateSalesRequest{
		Product:        "Test Product",
		Quantity:       2,
		Price:          10000,
		AmountReceived: 25000,
		IsDebt:         false,
	}

	sale, err := service.CreateSale(req)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if sale == nil {
		t.Error("Expected sale to be created, got nil")
	}

	if sale.Product != req.Product {
		t.Errorf("Expected product %s, got %s", req.Product, sale.Product)
	}
}

func TestCreateSale_InsufficientAmount(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{}
	service := NewSaleService(mockRepo)

	req := &models.CreateSalesRequest{
		Product:        "Test Product",
		Quantity:       2,
		Price:          10000,
		AmountReceived: 15000, // Less than total (20000)
		IsDebt:         false,
	}

	sale, err := service.CreateSale(req)

	if err == nil {
		t.Error("Expected error for insufficient amount, got nil")
	}

	if sale != nil {
		t.Error("Expected nil sale when amount is insufficient")
	}

	expectedErr := "amount received is less than total price"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestGetSaleByID_Success(t *testing.T) {
	expectedID := uuid.New()
	mockRepo := &repository.MockSaleRepository{
		GetByIDFunc: func(id uuid.UUID) (*models.Sale, error) {
			if id == expectedID {
				return &models.Sale{
					ID:       expectedID,
					Product:  "Test Product",
					Quantity: 1,
					Price:    5000,
				}, nil
			}
			return nil, nil
		},
	}

	service := NewSaleService(mockRepo)

	sale, err := service.GetSaleByID(expectedID.String())

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if sale == nil {
		t.Error("Expected sale to be returned, got nil")
	}

	if sale.ID != expectedID {
		t.Errorf("Expected ID %s, got %s", expectedID, sale.ID)
	}
}

func TestGetSaleByID_InvalidUUID(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{}
	service := NewSaleService(mockRepo)

	sale, err := service.GetSaleByID("invalid-uuid")

	if err == nil {
		t.Error("Expected error for invalid UUID, got nil")
	}

	if sale != nil {
		t.Error("Expected nil sale for invalid UUID")
	}

	expectedErr := "invalid UUID format"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestGetSaleByID_NotFound(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{
		GetByIDFunc: func(id uuid.UUID) (*models.Sale, error) {
			return nil, nil // Not found
		},
	}

	service := NewSaleService(mockRepo)

	sale, err := service.GetSaleByID(uuid.New().String())

	if err == nil {
		t.Error("Expected error for not found sale, got nil")
	}

	if sale != nil {
		t.Error("Expected nil sale when not found")
	}

	expectedErr := "sale not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestGetAllSales_Success(t *testing.T) {
	expectedSales := []*models.Sale{
		{ID: uuid.New(), Product: "Product 1"},
		{ID: uuid.New(), Product: "Product 2"},
	}

	mockRepo := &repository.MockSaleRepository{
		GetAllFunc: func() ([]*models.Sale, error) {
			return expectedSales, nil
		},
	}

	service := NewSaleService(mockRepo)

	sales, err := service.GetAllSales()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(sales) != len(expectedSales) {
		t.Errorf("Expected %d sales, got %d", len(expectedSales), len(sales))
	}
}

func TestUpdateSales_Success(t *testing.T) {
	expectedID := uuid.New()
	mockRepo := &repository.MockSaleRepository{
		UpdateFunc: func(id uuid.UUID, req *models.UpdateSaleRequest) (*models.Sale, error) {
			return &models.Sale{
				ID:       id,
				Product:  req.Product,
				Quantity: req.Quantity,
				Price:    req.Price,
			}, nil
		},
	}

	service := NewSaleService(mockRepo)

	req := &models.UpdateSaleRequest{
		Product:  "Updated Product",
		Quantity: 5,
		Price:    15000,
	}

	sale, err := service.UpdateSales(expectedID.String(), req)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if sale == nil {
		t.Error("Expected updated sale, got nil")
	}

	if sale.Product != req.Product {
		t.Errorf("Expected product %s, got %s", req.Product, sale.Product)
	}
}

func TestUpdateSales_InvalidUUID(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{}
	service := NewSaleService(mockRepo)

	req := &models.UpdateSaleRequest{
		Product: "Updated Product",
	}

	sale, err := service.UpdateSales("invalid-uuid", req)

	if err == nil {
		t.Error("Expected error for invalid UUID, got nil")
	}

	if sale != nil {
		t.Error("Expected nil sale for invalid UUID")
	}

	expectedErr := "invalid UUID format"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestUpdateSales_InsufficientAmount(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{}
	service := NewSaleService(mockRepo)

	req := &models.UpdateSaleRequest{
		Quantity:       2,
		Price:          10000,
		AmountReceived: 15000, // Less than total (20000)
	}

	sale, err := service.UpdateSales(uuid.New().String(), req)

	if err == nil {
		t.Error("Expected error for insufficient amount, got nil")
	}

	if sale != nil {
		t.Error("Expected nil sale when amount is insufficient")
	}
}

func TestDeleteSales_Success(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{
		DeleteFunc: func(id uuid.UUID) error {
			return nil
		},
	}

	service := NewSaleService(mockRepo)

	err := service.DeleteSales(uuid.New().String())

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteSales_InvalidUUID(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{}
	service := NewSaleService(mockRepo)

	err := service.DeleteSales("invalid-uuid")

	if err == nil {
		t.Error("Expected error for invalid UUID, got nil")
	}

	expectedErr := "invalid UUID format"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestDeleteSales_NotFound(t *testing.T) {
	mockRepo := &repository.MockSaleRepository{
		DeleteFunc: func(id uuid.UUID) error {
			return errors.New("sale not found")
		},
	}

	service := NewSaleService(mockRepo)

	err := service.DeleteSales(uuid.New().String())

	if err == nil {
		t.Error("Expected error for not found sale, got nil")
	}
}
