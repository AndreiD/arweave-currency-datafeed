package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

// shows the last error...
var lastError string

// shows if the bot is running
func healthHandler(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.JSON(http.StatusOK, gin.H{"status": "alive", "last_error": lastError, "alloc": bToMb(m.Alloc), "total_alloc": bToMb(m.TotalAlloc),
		"sys": bToMb(m.Sys), "num_gc": m.NumGC})
}

// converts bytes to Mb
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
