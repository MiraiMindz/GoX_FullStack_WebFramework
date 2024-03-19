//go:build amd64 && !wasm
// +build amd64,!wasm

package main

import (
	"fmt"
)

func AnyF() {
	fmt.Println("Test 4")
}
