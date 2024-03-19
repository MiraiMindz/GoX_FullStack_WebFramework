//go:build wasm && !amd64
// +build wasm,!amd64

package main

import (
	"fmt"
)

func AnyFunc() {
	fmt.Println("Test 2")
}
