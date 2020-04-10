package wasm

import (
	"syscall/js"
)

type unitaryNDilationHandler func([]float64, int) ([]float64, error)

func typedArrayToByteSlice(arg js.Value) []float64 {
    length := arg.Length()
	floats := make([]float64, length)

    for i := 0; i < length; i++ {
        floats[i] = arg.Index(i).Float()
	}

    return floats
}

func UnitaryNDilation(dilate unitaryNDilationHandler, this js.Value, inputs []js.Value) interface{} {
    degree := inputs[1].Int()
    value := typedArrayToByteSlice(inputs[0])
	dilation, err := dilate(value, degree)

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
