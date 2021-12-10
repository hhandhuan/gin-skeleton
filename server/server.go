package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/hhandhuan/gin-skeleton/internal/route"
	waggeries "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/hhandhuan/gin-skeleton/docs"
)

// Run server run
// @title gin-skeleton
// @version 1.0
// @description gin-skeleton 示例项目
// @host 127.0.0.1:8080
func Run() {
	engine := gin.New()
	if configs.Conf.Server.Pprof {
		pprof.Register(engine)
	}
	engine.Use(gin.Logger(), gin.Recovery())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(waggeries.Handler))
	server := &http.Server{
		Addr:           configs.Conf.Server.Addr,
		Handler:        route.Register(engine),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1m
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// monitor service termination signal.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Blocked
	<-quit

	// delay in launching services
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("timeout of %d seconds.", 1))
	}
}
