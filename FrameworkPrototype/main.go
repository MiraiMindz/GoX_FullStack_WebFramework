//go:build wasm && !amd64
// +build wasm,!amd64

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test")
	AnyFunc()
}
