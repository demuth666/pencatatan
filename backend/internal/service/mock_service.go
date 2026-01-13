package service

import (
	"pencatatan/internal/models"
)

// MockSaleService is a mock implementation of SaleService for testing
type MockSaleService struct {
	CreateSaleFunc  func(req *models.CreateSalesRequest) (*models.Sale, error)
	GetSaleByIDFunc func(id string) (*models.Sale, error)
	GetAllSalesFunc func() ([]*models.Sale, error)
	UpdateSalesFunc func(id string, req *models.UpdateSaleRequest) (*models.Sale, error)
	DeleteSalesFunc func(id string) error
}

func (m *MockSaleService) CreateSale(req *models.CreateSalesRequest) (*models.Sale, error) {
	if m.CreateSaleFunc != nil {
		return m.CreateSaleFunc(req)
	}
	return nil, nil
}

func (m *MockSaleService) GetSaleByID(id string) (*models.Sale, error) {
	if m.GetSaleByIDFunc != nil {
		return m.GetSaleByIDFunc(id)
	}
	return nil, nil
}

func (m *MockSaleService) GetAllSales() ([]*models.Sale, error) {
	if m.GetAllSalesFunc != nil {
		return m.GetAllSalesFunc()
	}
	return nil, nil
}

func (m *MockSaleService) UpdateSales(id string, req *models.UpdateSaleRequest) (*models.Sale, error) {
	if m.UpdateSalesFunc != nil {
		return m.UpdateSalesFunc(id, req)
	}
	return nil, nil
}

func (m *MockSaleService) DeleteSales(id string) error {
	if m.DeleteSalesFunc != nil {
		return m.DeleteSalesFunc(id)
	}
	return nil
}
