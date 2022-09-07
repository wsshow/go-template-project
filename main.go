package main

import (
	"context"
	"gtp/config"
	"gtp/log"
	"gtp/middleware"
	"gtp/router"
	"gtp/version"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer() *http.Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Recorder(), middleware.Cors())
	localPort := config.DefaultPort()
	bHttps := config.IsHttps()
	if bHttps {
		r.Use(middleware.LoadTls(localPort))
	}
	router.Init(r)
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(localPort),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}
	go func() {
		if bHttps {
			if err := srv.ListenAndServeTLS("./certificate/server.pem", "./certificate/server.key"); err != nil && err != http.ErrServerClosed {
				log.Fatal("srv.ListenAndServeTLS error:", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("srv.ListenAndServe error:", err)
			}
		}
	}()
	return srv
}

func ServerRun() {
	srv := InitServer()
	log.Info("before capture signal. the number of goroutines: ", runtime.NumGoroutine())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-c
	close(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Info("after capture signal. the remain number of goroutines: ", runtime.NumGoroutine())
}

func main() {
	log.Init("gtp.log")
	log.Info(version.Get().String())
	ServerRun()
}
