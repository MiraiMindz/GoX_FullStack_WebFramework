package js

import (
	"syscall/js"
)

func Update(updateFunc interface{}) {
	res := updateFunc.(func() string)()
	js.Global().Get("document").Call("getElementById", "_INTERNAL_GOX_APP").Set("innerHTML", res)
}
