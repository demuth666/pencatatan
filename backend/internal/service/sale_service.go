package service

import (
	"errors"
	"pencatatan/internal/models"
	"pencatatan/internal/repository"

	"github.com/google/uuid"
)

type SaleService interface {
	CreateSale(req *models.CreateSalesRequest) (*models.Sale, error)
	GetSaleByID(id string) (*models.Sale, error)
	GetAllSales() ([]*models.Sale, error)
	UpdateSales(id string, req *models.UpdateSaleRequest) (*models.Sale, error)
	DeleteSales(id string) error
}

type saleService struct {
	repo repository.SaleRepository
}

func NewSaleService(repo repository.SaleRepository) SaleService {
	return &saleService{
		repo: repo,
	}
}

func (s *saleService) CreateSale(req *models.CreateSalesRequest) (*models.Sale, error) {
	return s.repo.Create(req)
}

func (s *saleService) GetSaleByID(id string) (*models.Sale, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	sale, err := s.repo.GetByID(uid)
	if err != nil {
		return nil, err
	}

	if sale == nil {
		return nil, errors.New("sale not found")
	}

	return sale, nil
}

func (s *saleService) GetAllSales() ([]*models.Sale, error) {
	return s.repo.GetAll()
}

func (s *saleService) UpdateSales(id string, req *models.UpdateSaleRequest) (*models.Sale, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	if req.Quantity > 0 && req.Price > 0 && req.AmountReceived != 0 {
		total := float64(req.Quantity) * req.Price
		if req.AmountReceived < total {
			return nil, errors.New("amount received is less than total price")
		}
	}

	sale, err := s.repo.Update(uid, req)
	if err != nil {
		return nil, err
	}

	if sale == nil {
		return nil, errors.New("sale not found")
	}

	return sale, nil
}

func (s *saleService) DeleteSales(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid UUID format")
	}

	err = s.repo.Delete(uid)
	if err != nil {
		return errors.New("sale not found")
	}

	return nil
}
