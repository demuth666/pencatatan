package handler

import (
	"net/http"
	"pencatatan/internal/database"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	db database.Service
}

func NewHealthHandler(db database.Service) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, h.db.Health())
}
