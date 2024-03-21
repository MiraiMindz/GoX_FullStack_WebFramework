package wasm

import (
	"syscall/js"

	"FRAMEWORK/internal/html"
)

func App() {
	app := js.Global().Get("document").Call("getElementById", "_INTERNAL_GOX_APP")

	s := html.CreateBareHTMLTemplate("App", `
<main>
	<h1>Hello from: {{.First}} and {{.Second}}</h1>
</main>
		`, struct {
		First  string
		Second string
	}{First: "Hello", Second: "World"})

	app.Set("innerHTML", s)

}
