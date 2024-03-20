//go:build wasm && !amd64
// +build wasm,!amd64

package main

import (
	"fmt"
	"io"
	"net/http"
	"syscall/js"

	"third"
)

var counter int

func main() {
	var buttonsList []string = []string{
		"bg-neutral-950", "border-2", "rounded-md", "p-2", "border-neutral-950", "text-neutral-50",
		"hover:bg-neutral-50", "hover:border-neutral-950", "hover:text-neutral-950", "transition-all",
	}
	doc := js.Global().Get("document")

	// Get the body element
	body := doc.Get("body")

	// Add classes to the body element
	body.Get("classList").Call("add", "w-screen", "min-h-screen")

	app := doc.Call("getElementById", "_INTERNAL_GOX_APP")
	app.Get("classList").Call("add", "w-full", "min-h-full", "bg-neutral-50", "text-neutral-950")

	mainWrapper := doc.Call("createElement", "main")
	mainWrapper.Get("classList").Call("add", "w-screen", "min-h-screen", "flex", "flex-col", "items-center", "p-2")
	app.Call("appendChild", mainWrapper)

	headerWrapper := doc.Call("createElement", "div")
	header := doc.Call("createElement", "h1")
	header.Set("innerText", "Go Full-Stack Framework GoX")
	header.Get("classList").Call("add", "text-2xl", "font-bold")

	headerWrapper.Get("classList").Call("add", "flex", "flex-col", "items-center", "justify-start")

	headerWrapper.Call("appendChild", header)

	subHeader := doc.Call("createElement", "h1")
	subHeader.Set("innerText", "Powered by WebAssembly")
	subHeader.Get("classList").Call("add", "text-lg", "text-neutral-600", "font-bold")
	headerWrapper.Call("appendChild", subHeader)

	mainWrapper.Call("appendChild", headerWrapper)

	wasmSvg := doc.Call("createElementNS", "http://www.w3.org/2000/svg", "svg")

	wasmSvg.Get("classList").Call("add", "w-32", "aspect-square", "mx-2")

	// Set attributes for the SVG element
	//wasmSvg.Call("setAttributeNS", nil, "height", 2500)
	wasmSvg.Call("setAttributeNS", nil, "preserveAspectRatio", "xMidYMid")
	//wasmSvg.Call("setAttributeNS", nil, "width", 2500)
	wasmSvg.Call("setAttributeNS", nil, "viewBox", "0 0 256 255.997")

	// Create the path element
	path := doc.Call("createElementNS", "http://www.w3.org/2000/svg", "path")

	// Set attributes for the path element
	path.Call("setAttributeNS", nil, "d", "M157.29 0H256v255.997H0V0h98.71c-.02.458-.068.911-.068 1.375 0 16.215 13.144 29.36 29.358 29.36s29.36-13.145 29.36-29.36c0-.464-.047-.917-.07-1.375zm58.327 228.6h17.598l-26.657-90.632h-26.753L157.853 228.6h17.058l4.444-20.18h30.485zm-93.4 0h16.842l21.609-90.632H144.04l-13.236 62.453h-.216l-12.57-62.453h-15.871l-13.984 61.69h-.217l-11.604-61.69H59.39L78.92 228.6h17.166l13.447-61.69h.216zm67.784-68.291h7.13l8.502 33.258h-23.03z")
	path.Call("setAttributeNS", nil, "fill", "#654ff0")

	// Append the path element to the SVG element
	wasmSvg.Call("appendChild", path)

	tmpSvg := doc.Call("createElement", "div")
	tmpSvg.Set("innerHTML", `<svg viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg">
<path fill-rule="evenodd" clip-rule="evenodd" d="M18.1177 14.0442C17.7408 14.1497 17.3586 14.2566 16.9162 14.3768C16.7001 14.438 16.6509 14.4519 16.4498 14.2074C16.2086 13.9194 16.0317 13.7331 15.6939 13.5636C14.6807 13.0384 13.6996 13.1909 12.7829 13.8178C11.6893 14.5632 11.1264 15.6644 11.1425 17.0367C11.1585 18.3921 12.0431 19.5103 13.3137 19.6966C14.4073 19.8491 15.324 19.4425 16.0477 18.5785C16.1924 18.3922 16.3212 18.1887 16.482 17.9516H13.378C13.0402 17.9516 12.9598 17.7314 13.0724 17.4433C13.2815 16.9181 13.6675 16.0372 13.8926 15.5967C13.9409 15.495 14.0535 15.3256 14.2947 15.3256H19.4702C19.7027 14.5496 20.0799 13.8164 20.5831 13.1226C21.7572 11.4961 23.1725 10.649 25.0863 10.2933C26.7268 9.9883 28.2707 10.1577 29.6699 11.1573C30.9405 12.0722 31.7285 13.3089 31.9376 14.9354C32.211 17.2225 31.5838 19.0862 30.0881 20.6787C29.0266 21.8138 27.7239 22.5254 26.2282 22.8473C25.9429 22.9029 25.6576 22.9293 25.3768 22.9553C25.2303 22.9689 25.085 22.9823 24.9416 22.9998C23.478 22.9659 22.1432 22.5254 21.0173 21.5089C20.2256 20.7879 19.6803 19.9019 19.4092 18.8705C19.2211 19.2707 18.9962 19.6539 18.7336 20.0185C17.5756 21.628 16.0638 22.6276 14.15 22.8987C12.5738 23.1189 11.1103 22.797 9.82366 21.7805C8.63353 20.8317 7.95805 19.578 7.78114 18.0194C7.57206 16.1727 8.08671 14.5124 9.14818 13.0554C10.2901 11.4798 11.8019 10.4802 13.6514 10.1244C15.1632 9.8364 16.6106 10.0228 17.9134 10.9546C18.7657 11.5475 19.3769 12.3608 19.779 13.3434C19.8755 13.4959 19.8111 13.5806 19.6181 13.6314C19.0545 13.7822 18.5903 13.9121 18.1177 14.0442ZM28.7581 15.974C28.7613 16.0309 28.7646 16.0909 28.7693 16.1552C28.6889 17.6122 27.9973 18.6965 26.7268 19.3911C25.8744 19.8485 24.9898 19.8994 24.1053 19.4928C22.9473 18.9506 22.3361 17.6122 22.6256 16.2907C22.9795 14.6982 23.9444 13.6986 25.4401 13.3428C26.968 12.9701 28.4316 13.9188 28.7211 15.5961C28.7438 15.7161 28.7505 15.836 28.7581 15.974Z" fill="#00ACD7"/>
<path d="M2.44461 13.8517C2.41244 13.9025 2.42852 13.9364 2.49285 13.9364L7.2826 13.9534C7.33085 13.9534 7.41126 13.9025 7.44343 13.8517L7.71684 13.4112C7.749 13.3604 7.73292 13.3096 7.66859 13.3096H2.95926C2.89493 13.3096 2.81451 13.3435 2.78235 13.3943L2.44461 13.8517Z" fill="#00ACD7"/>
<path d="M0.0160829 15.4103C-0.0160829 15.4611 7.45058e-09 15.495 0.0643316 15.495L6.63928 15.4781C6.70361 15.4781 6.76794 15.4442 6.78402 15.3764L6.91269 14.9698C6.92877 14.919 6.8966 14.8682 6.83227 14.8682H0.530735C0.466404 14.8682 0.385989 14.902 0.353823 14.9529L0.0160829 15.4103Z" fill="#00ACD7"/>
<path d="M3.90813 16.9521C3.87596 17.0029 3.89204 17.0537 3.95638 17.0537L6.43019 17.0707C6.47843 17.0707 6.54277 17.0199 6.54277 16.9521L6.57493 16.5455C6.57493 16.4777 6.54277 16.4269 6.47843 16.4269H4.29412C4.22978 16.4269 4.16545 16.4777 4.13329 16.5285L3.90813 16.9521Z" fill="#00ACD7"/>
</svg>`)

	goSvg := tmpSvg.Get("childNodes").Index(0)
	goSvg.Get("classList").Call("add", "w-64", "aspect-square", "mx-2")

	/*
			tmpSvg.Set("innerHTML", `<svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
		<path d="M4 12H20M12 4V20" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
		</svg>`)

			plusSvg := tmpSvg.Get("childNodes").Index(0)
			plusSvg.Get("classList").Call("add", "stroke-2", "stroke-neutral-950", "aspect-square", "w-24", "mx-2")

	*/
	leftWrapper := doc.Call("createElement", "div")
	//logosWrapper.Get("classList").Call("add", "flex", "items-center", "justify-between")

	rightWrapper := doc.Call("createElement", "div")
	leftWrapper.Call("appendChild", goSvg)
	//logosWrapper.Call("appendChild", plusSvg)
	rightWrapper.Call("appendChild", wasmSvg)

	paragraph := doc.Call("createElement", "p")
	paragraph.Set("innerText", fmt.Sprintf("[CLIENT SIDE] Counter %d", counter))
	rightWrapper.Call("appendChild", paragraph)

	button := doc.Call("createElement", "button")
	button.Set("innerText", "[CLIENT SIDE] Increment Counter")
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Increment the counter directly
		counter++
		paragraph.Set("innerHTML", fmt.Sprintf("[CLIENT SIDE] %s %d", third.Third("CLIENT"), counter))

		// Print the counter to console (for demonstration)
		fmt.Println("38: Counter:\t", counter)

		return nil
	}))

	for _, c := range buttonsList {
		button.Get("classList").Call("add", c)
	}
	rightWrapper.Call("appendChild", button)

	// Create a div element for displaying the response
	serverParagraph := doc.Call("createElement", "p")
	serverParagraph.Set("innerText", "[SERVER SIDE] this is being rendered on the server 0")
	leftWrapper.Call("appendChild", serverParagraph)

	fmt.Println(third.Third("CLIENT"))

	serverButton := doc.Call("createElement", "button")
	serverButton.Set("innerHTML", "[SERVER SIDE] Get data")
	serverButton.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Make a GET request
		go func() {
			resp, err := http.Get("http://localhost:8080/portal")
			if err != nil {
				fmt.Println("56: Error fetching data:", err)
			}
			defer resp.Body.Close()

			// Read the response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("64: Error reading response body:", err)
			}

			// Update the paragraph content with the response
			serverParagraph.Set("innerHTML", string(body))

			// Append the response to the existing content
			//icurrentContent := responseParagraph.Get("innerHTML").String()
			//updatedContent := currentContent + string(body)
			//responseParagraph.Set("innerHTML", updatedContent)
		}()

		return nil
	}))

	for _, c := range buttonsList {
		serverButton.Get("classList").Call("add", c)
	}
	leftWrapper.Call("appendChild", serverButton)

	wrappers := doc.Call("createElement", "div")
	wrappers.Get("classList").Call("add", "flex", "items-center", "justify-between")

	ws := doc.Call("createElement", "div")
	ws.Get("classList").Call("add", "w-full", "h-full", "flex", "items-center", "justify-center", "grow")

	wrappers.Call("appendChild", leftWrapper)
	wrappers.Call("appendChild", rightWrapper)
	ws.Call("appendChild", wrappers)
	mainWrapper.Call("appendChild", ws)

	AnyFunc()

	select {} // Halts the main go routine
}
