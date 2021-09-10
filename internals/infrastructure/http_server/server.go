package httpserver

import (
	"context"
	"fmt"
	"lm-test/internals/config"
	"lm-test/internals/controller"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Route     *gin.Engine
	Config    config.Configuration
	CovidCtrl *controller.CovidController
}

func (h *HttpServer) Configuration() *gin.Engine {
	r := h.Route
	// Health check
	r.GET("/healthz", func(c *gin.Context) {
		c.String(200, "OK")
	})

	r.GET("/covid/summary", h.CovidCtrl.Summary)

	return r
}

// Start ...
func (h *HttpServer) Start() {
	router := h.Configuration()
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(h.Config.Port),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("listen: err: ", err)
		}
	}()
	fmt.Println("Listen and serve on port:", h.Config.Port)
	// Gracefully Shutdown
	// Make channel listen for signals from OS
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	<-gracefulStop
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}
	fmt.Println("Server exiting")
}

func NewHttpServer(config config.Configuration, covid *controller.CovidController) *HttpServer {
	r := gin.New()

	return &HttpServer{
		Config:    config,
		Route:     r,
		CovidCtrl: covid,
	}
}
