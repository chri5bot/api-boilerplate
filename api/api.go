package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/chri5bot/api-boilerplate/conf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// API is the struct that contains all api functionality
type API struct {
	db      *gorm.DB
	config  *conf.Configuration
	handler *gin.Engine
}

// NewAPI mounts all routes
func NewAPI(db *gorm.DB, config *conf.Configuration) *API {
	api := &API{
		db:     db,
		config: config,
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		log.Println("request")

		c.JSON(200, gin.H{"message": "API Service OK!"})
	})

	api.handler = r

	return api
}

// ListenAndServe starts the API
func (a *API) ListenAndServe() {
	host := fmt.Sprintf(":%d", a.config.Port)

	srv := &http.Server{
		Addr:    host,
		Handler: a.handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	log.Printf("Server listening on address %s\n", host)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
