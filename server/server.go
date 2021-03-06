package server

import (
	"context"
	"github.com/hhandhuan/gin-skeleton/logger"
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

func Run() {
	gin.SetMode(configs.Conf.Server.Mode)
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
			logger.I.Fatalf("listen: %s\n", err)
		}
	}()

	// 监听终端信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 延时终止服务避免未处理完之前的请求
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.I.Printf("Server Shutdown: %v \n", err)
	}
	select {
	case <-ctx.Done():
		logger.I.Printf("timeout of %d seconds. \n", 1)
	}
}
