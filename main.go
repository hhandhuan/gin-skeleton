package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/app/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	timeout = time.Second * 5
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", "8080"),
		Handler:        routes.Register(e),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
}
