package main

import (
	"banking-auth/app"
	"banking-auth/logger"
)

func main() {
	logger.Info("Starting the auth server...")
	app.Start()
}
