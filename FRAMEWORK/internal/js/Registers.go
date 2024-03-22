package js

import (
	"reflect"
	"runtime"
	"syscall/js"
)

func getFunctionName(f interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	lastDot := len(fullName)
	for i := len(fullName) - 1; i >= 0; i-- {
		if fullName[i] == '.' {
			lastDot = i
			break
		}
	}
	return fullName[lastDot+1:]
}

func RegisterFunction(updateFunc func() string, f func(this js.Value, inputs []js.Value) interface{}) {
	funcName := getFunctionName(f)
	modifiedFunction := func(this js.Value, inputs []js.Value) interface{} {
		// Call the original function passed as an argument
		result := f(this, inputs)
		// Update the HTML content
		Update(updateFunc)
		return result
	}
	js.Global().Set(funcName, js.FuncOf(modifiedFunction))
}
