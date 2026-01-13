package server

import (
	"net/http"
	"pencatatan/internal/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, c *app.Container) http.Handler {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/health", c.HealthHandler.Check)

	api := r.Group("/api")

	sales := api.Group("/sales")
	{
		sales.POST("", c.SaleHandler.CreateSale)
		sales.GET("/:id", c.SaleHandler.GetSaleByID)
		sales.GET("", c.SaleHandler.GetAllSales)
		sales.PUT("/:id", c.SaleHandler.UpdateSale)
		sales.DELETE("/:id", c.SaleHandler.DeleteSale)
	}

	return r
}
