package handler

import (
	"net/http"
	"pencatatan/internal/models"
	"pencatatan/internal/service"

	"github.com/gin-gonic/gin"
)

type SaleHandler struct {
	service service.SaleService
}

func NewSaleHandler(service service.SaleService) *SaleHandler {
	return &SaleHandler{
		service: service,
	}
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func (h *SaleHandler) CreateSale(c *gin.Context) {
	var req models.CreateSalesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	sale, err := h.service.CreateSale(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Sale created successfully",
		Data:    sale,
	})
}

func (h *SaleHandler) GetSaleByID(c *gin.Context) {
	id := c.Param("id")

	sale, err := h.service.GetSaleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    sale,
	})
}

func (h *SaleHandler) GetAllSales(c *gin.Context) {
	sales, err := h.service.GetAllSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    sales,
	})
}

func (h *SaleHandler) UpdateSale(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateSaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	sale, err := h.service.UpdateSales(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Sale updated successfully",
		Data:    sale,
	})
}

func (h *SaleHandler) DeleteSale(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteSales(id)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Sale deleted successfully",
	})
}
