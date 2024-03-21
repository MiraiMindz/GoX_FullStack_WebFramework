//go:build wasm && !amd64
// +build wasm,!amd64

package main

import (
	"syscall/js"

	"FRAMEWORK/internal/html"
)

func main() {
	app := js.Global().Get("document").Call("getElementById", "_INTERNAL_GOX_APP")

	s := html.CreateBareHTMLTemplate("App", `
<main>
	<h1>Hello from: {{.First}} and {{.Second}}</h1>
</main>
		`, map[string]any{
		"First":  "Hello",
		"Second": 12,
	})

	app.Set("innerHTML", s)

}
