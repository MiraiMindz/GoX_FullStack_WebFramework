package main

import (
	"FRAMEWORK/internal"
	"FRAMEWORK/server"
	"FRAMEWORK/utils"
	"FRAMEWORK/wasm"
	"FRAMEWORK/wasm/components"
	"FRAMEWORK/wasm/pages"
)

func main() {
	server.New()
	utils.Helper()
	internal.Serve()
	wasm.App()
	components.Bar()
	pages.Index()
}
