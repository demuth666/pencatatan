package server

import (
	"fmt"
	"net/http"
	"pencatatan/internal/app"
	"pencatatan/internal/config"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer(cfg *config.Config, c *app.Container) *http.Server {
	r := gin.Default()
	Register(r, c)

	return &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerPort),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
