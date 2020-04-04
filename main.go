package main

import (
	"fmt"
	"github.com/acra5y/n-dilation-webassembly/internal/handler"
	"github.com/acra5y/go-dilation"
	"syscall/js"
)

func typedArrayToByteSlice(arg js.Value) []float64 {
    length := arg.Length()
	floats := make([]float64, length)

    for i := 0; i < length; i++ {
        floats[i] = arg.Index(i).Float()
	}

    return floats
}

func dilation(this js.Value, inputs []js.Value) interface{} {
    degree := inputs[1].Int()
    value := typedArrayToByteSlice(inputs[0])
	dilation, err := handler.DilationHandler(godilation.UnitaryNDilation, value, degree)

	if (err != nil) {
		return js.ValueOf(map[string]interface{}{
			"error": js.ValueOf(err.Error()),
		})
	}

	ret := make([]interface{}, len(dilation))
	for i, entry := range dilation {
		ret[i] = js.ValueOf(entry)
	}

	return js.ValueOf(map[string]interface{}{
		"value": js.ValueOf(ret),
	})
}

func main() {
	fmt.Printf("Running wasm\n")
	js.Global().Set("dilation", js.FuncOf(dilation))

    c := make(chan bool)
	<-c // make sure go program does not exit and dilation function is always callable
}
