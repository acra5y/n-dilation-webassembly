package main

import (
	"fmt"
	"github.com/acra5y/go-dilation"
	"github.com/acra5y/n-dilation-webassembly/internal/handler"
	"github.com/acra5y/n-dilation-webassembly/internal/wasm"
	"syscall/js"
)

func dilation(this js.Value, inputs []js.Value) interface{} {
    return wasm.UnitaryNDilation(handler.UnitaryNDilation(godilation.UnitaryNDilation), this, inputs)
}

func main() {
	fmt.Printf("Running wasm\n")
	js.Global().Set("UnitaryNDilation", js.FuncOf(dilation))

    c := make(chan bool)
	<-c // make sure go program does not exit and dilation function is always callable
}
