//go:build wasm && !amd64
// +build wasm,!amd64

package main

import (
	"syscall/js"

	"FRAMEWORK/wasm/pages"
)

func main() {
	app := js.Global().Get("document").Call("getElementById", "_INTERNAL_GOX_APP")

	app.Set("innerHTML", pages.App())

	select {}
}
