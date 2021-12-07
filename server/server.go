package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/hhandhuan/gin-skeleton/docs"
	"github.com/hhandhuan/gin-skeleton/internal/routes"
	waggeries "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	timeout = time.Second * 1
)

// Run RunWebServer
func Run() {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	// swagger docs url(:8080/swagger/index.html)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(waggeries.Handler))

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes.Register(engine),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// monitor service termination signal.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// delay in launching services
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("timeout of %d seconds.", timeout))
	}
}
