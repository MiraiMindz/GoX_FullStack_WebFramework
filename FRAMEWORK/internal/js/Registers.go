package js

import (
	"fmt"
	"reflect"
	"syscall/js"
)

// func CallFunction(f interface{}, args ...interface{}) []reflect.Value {
// 	// Convert f to a reflect.Value
// 	fv := reflect.ValueOf(f)
//
// 	// Check if f is a function
// 	if fv.Kind() != reflect.Func {
// 		return nil
// 	}
//
// 	// Convert the arguments to reflect.Values
// 	var reflectArgs []reflect.Value
// 	for _, arg := range args {
// 		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
// 	}
//
// 	// Call the function with the arguments
// 	result := fv.Call(reflectArgs)
//
// 	return result
// }

type UpdateFunc func(...interface{}) []interface{}

func addUpdateMethod(upFunc, fn interface{}) UpdateFunc {
	// Convert the function to a reflect.Value
	originalFn := reflect.ValueOf(fn)

	// Define the Update function
	updateFunc := func(args ...interface{}) []interface{} {
		// Call the original function with the provided arguments
		originalArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			originalArgs[i] = reflect.ValueOf(arg)
		}
		results := originalFn.Call(originalArgs)

		// Add your update logic here if needed
		Update(upFunc)

		// Convert results back to interface{} and return
		updatedResults := make([]interface{}, len(results))
		for i, result := range results {
			updatedResults[i] = result.Interface()
		}
		return updatedResults
	}

	return updateFunc
}
func RegisterFunc(updateFunc interface{}, funcName string, originalFunction interface{}) {
	fmt.Println("Registering function: ", funcName)
	fnValue := reflect.ValueOf(originalFunction)

	// Wrap the original function
	wrappedFunction := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Call the update function before executing the original function
		Update(updateFunc)

		// If the original function has arguments, convert the Go function arguments to reflect.Values
		if len(args) > 0 {
			var reflectArgs []reflect.Value
			for _, arg := range args {
				reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
			}

			// Call the original function with the provided arguments
			results := fnValue.Call(reflectArgs)

			// Convert the results to []interface{}
			var interfaceResults []interface{}
			for _, result := range results {
				interfaceResults = append(interfaceResults, result.Interface())
			}

			return interfaceResults
		}

		// If the original function doesn't have arguments, simply return nil
		fnValue.Call([]reflect.Value{})
		return nil
	})

	// Set the wrapped function to the global JavaScript object
	js.Global().Set(funcName, wrappedFunction)
}

// func RegisterFunc(updateFunc interface{}, fnName string, fn interface{}) {
// 	x := func(fnValue reflect.Value) reflect.Value {
// 		return reflect.MakeFunc(fnValue.Type(), func(args []reflect.Value) (results []reflect.Value) {
// 			Update(updateFunc)
//
// 			return fnValue.Call(args)
// 		})
// 	}(reflect.ValueOf(fn)).Interface()
// 	js.Global().Set(fnName, x)
// }

// func RegisterFunc(updateFunc interface{}, fnName string, fn interface{}, fnArgs ...interface{}) {
// 	fFunc := func(this js.Value, inputs []js.Value) interface{} {
// 		Update(updateFunc)
// 		// Convert fnArgs to a slice of empty interfaces
// 		args := make([]interface{}, len(fnArgs))
// 		for i, arg := range fnArgs {
// 			args[i] = arg
// 		}
//
// 		// Call the function with the unpacked arguments
// 		CallFunction(fn, args)
// 		return nil
// 	}
//
