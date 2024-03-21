//go:build wasm && !amd64
// +build wasm,!amd64

package main

import (
	"syscall/js"
)

func main() {
	app := js.Global().Get("document").Call("getElementById", "_INTERNAL_GOX_APP")
	s := "Teste"
	app.Set("innerHTML", s)
	select {}
}
