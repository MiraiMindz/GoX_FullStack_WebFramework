//go:build wasm && !amd64
// +build wasm,!amd64

package main

import (
	"fmt"
	"syscall/js"
)

var counter int

func main() {
	doc := js.Global().Get("document")
	fmt.Println("Got document element", doc)
	app := doc.Call("getElementById", "_INTERNAL_GOX_APP")
	fmt.Println("Got app element", app)
	header := doc.Call("createElement", "h1")
	fmt.Println("Created H1 element", header)
	header.Set("innerText", "Element Created from WebAssembly")
	fmt.Println("Set H1 text to 'Element Created from WebAssembly'", header)
	app.Call("appendChild", header)
	fmt.Println("Appended H1 to app", app)

	paragraph := doc.Call("createElement", "p")
	paragraph.Set("innerText", fmt.Sprintf("Counter %d", counter))
	app.Call("appendChild", paragraph)

	button := doc.Call("createElement", "button")
	button.Set("innerText", "Increment Counter")
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Increment the counter directly
		counter++
		paragraph.Set("innerHTML", fmt.Sprintf("Counter: %d", counter))

		// Print the counter to console (for demonstration)
		fmt.Println("Counter:\t", counter)

		return nil
	}))
	app.Call("appendChild", button)

	AnyFunc()

	select {} // Halts the main go routine
}
