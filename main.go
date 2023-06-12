package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"simple-payment/delivery"
	"simple-payment/util"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// logging to a file
	newpath := filepath.Join(".", "logs-history")
	err := os.MkdirAll(newpath, os.ModePerm)
	currentTime := time.Now().Format("20060102150405")
	f, err := os.Create("./logs-history/history-" + currentTime + ".log")
	if err != nil {
		fmt.Println("Failed create log file", err)
	}

	saveLogToDb := util.SaveLog{}

	gin.DefaultWriter = io.MultiWriter(f, &saveLogToDb, os.Stdout)

	delivery.NewServer("localhost:8080").Run()
}
