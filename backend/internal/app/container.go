package app

import (
	"pencatatan/internal/database"
	"pencatatan/internal/handler"
	"pencatatan/internal/repository"
	"pencatatan/internal/service"
)

type Container struct {
	HealthHandler *handler.HealthHandler
	SaleHandler   *handler.SaleHandler
}

func BuildContainer(db database.Service) *Container {
	healthHandler := handler.NewHealthHandler(db)

	saleRepo := repository.NewSaleRepository(db.DB())
	saleService := service.NewSaleService(saleRepo)
	saleHandler := handler.NewSaleHandler(saleService)

	return &Container{
		SaleHandler:   saleHandler,
		HealthHandler: healthHandler,
	}
}
