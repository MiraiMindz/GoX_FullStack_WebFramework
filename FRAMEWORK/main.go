package main

import (
	"FRAMEWORK/internal"
	"FRAMEWORK/server"
	"FRAMEWORK/utils"
)

func main() {
	server.New()
	utils.Helper()
	internal.Serve()
}
