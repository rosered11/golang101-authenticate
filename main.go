package main

import (
	"github.com/rosered11/golang101-authenticate/app"
	"github.com/rosered11/golang101-authenticate/logger"
)

func main() {
	//log.Println("Starting application.....")
	logger.Info("Starting application.....")
	app.Start()
}
