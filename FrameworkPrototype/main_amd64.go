//go:build amd64 && !wasm
// +build amd64,!wasm

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test 3")
	AnyF()
}
