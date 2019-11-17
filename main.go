package main

import (
	"arweave-datafeed/configs"
	"arweave-datafeed/utils/log"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const version = "1.0 Alpha"

var configuration *configs.ViperConfiguration
var router *gin.Engine

func init() {

	configuration = configs.NewConfiguration()
	configuration.Init()

	debug := configuration.GetBool("debug")
	log.Init(debug)

	log.Println("==================================================")
	log.Println("Starting Arweave Exchanges Bridge version: " + version)
	log.Println("==================================================")
	log.Println()
}

func main() {

	router = gin.New()

	if configuration.GetBool("debug") {
		log.Warn("application runs in debug mode.")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	api := router.Group("/api")
	api.GET("/health", healthHandler)

	server := &http.Server{
		Addr:           configuration.Get("server.host") + ":" + strconv.Itoa(configuration.GetInt("server.port")),
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 10, // 1Mb
	}
	server.SetKeepAlivesEnabled(true)

	// out main logic
	startCronService()

	// Serve'em
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	log.Printf("Running on %s:%s", configuration.Get("server.host"), strconv.Itoa(configuration.GetInt("server.port")))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("initiated server shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting. bye!")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
