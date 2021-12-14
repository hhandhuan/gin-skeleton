package server

import (
	"context"
	"github.com/hhandhuan/gin-skeleton/internal/middware"
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

func Run() {
	gin.SetMode(configs.Conf.Server.Mode)
	engine := gin.New()
	if configs.Conf.Server.Pprof {
		pprof.Register(engine)
	}
	engine.Use(middware.Logger(), middware.Recovery(true))
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

	// 监听信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 延时终止服务避免未处理完之前的请求
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown: %v \n", err)
	}
	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds. \n", 1)
	}
}
