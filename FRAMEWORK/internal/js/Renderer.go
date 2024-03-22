package js

import (
	"syscall/js"
)

func Update(updateFunc func() string) {
	js.Global().Get("document").Call("getElementById", "_INTERNAL_GOX_APP").Set("innerHTML", updateFunc())
}
