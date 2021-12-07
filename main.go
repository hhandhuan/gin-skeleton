package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/hhandhuan/gin-skeleton/db"
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

// @title gin-skeleton
// @version 1.0
// @description gin-skeleton 示例项目
// @host 127.0.0.1:8080
func main() {
	configs.Initialize()
	db.MysqlInitialize()
	RunServer()
}

// RunServer RunServer
func RunServer() {
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())
	// swagger docs 默认: 127.0.0.1:8080/swagger/index.html
	e.GET("/swagger/*any", ginSwagger.WrapHandler(waggeries.Handler))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", "8080"),
		Handler:        routes.Register(e),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("timeout of %d seconds.", timeout))
	}
}
