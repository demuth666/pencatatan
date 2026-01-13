package repository

import (
	"pencatatan/internal/models"

	"github.com/google/uuid"
)

// MockSaleRepository is a mock implementation of SaleRepository for testing
type MockSaleRepository struct {
	CreateFunc  func(sale *models.CreateSalesRequest) (*models.Sale, error)
	GetByIDFunc func(id uuid.UUID) (*models.Sale, error)
	GetAllFunc  func() ([]*models.Sale, error)
	UpdateFunc  func(id uuid.UUID, sale *models.UpdateSaleRequest) (*models.Sale, error)
	DeleteFunc  func(id uuid.UUID) error
}

func (m *MockSaleRepository) Create(sale *models.CreateSalesRequest) (*models.Sale, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(sale)
	}
	return nil, nil
}

func (m *MockSaleRepository) GetByID(id uuid.UUID) (*models.Sale, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, nil
}

func (m *MockSaleRepository) GetAll() ([]*models.Sale, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, nil
}

func (m *MockSaleRepository) Update(id uuid.UUID, sale *models.UpdateSaleRequest) (*models.Sale, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, sale)
	}
	return nil, nil
}

func (m *MockSaleRepository) Delete(id uuid.UUID) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}
