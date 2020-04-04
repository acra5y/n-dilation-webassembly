package handler

import (
	"fmt"
    "gonum.org/v1/gonum/mat"
    "math"
)

type unitaryNDilation func(*mat.Dense, int) (*mat.Dense, error)

func validatePayload(value []float64, degree int) error {
    if degree <= 0 {
        return fmt.Errorf("degree must be an integer greater than zero")
    }

    n := int(math.Sqrt(float64(len(value))))
    if len(value) == 0 || int(math.Pow(float64(n), 2)) != len(value) {
        return fmt.Errorf("value must contain a square number greater than zero of numbers")
	}

    return nil
}

func denseToSlice(u *mat.Dense) (data []float64) {
    m, _ := u.Dims()
    for i := 0; i < m; i++ {
        raw := u.RawRowView(i)
        data = append(data, raw...)
	}

    return
}

func DilationHandler(dilation unitaryNDilation, value []float64, degree int) ([]float64, error) {
    err := validatePayload(value, degree)
    if err != nil {
        return nil, err
    }

    n := int(math.Sqrt(float64(len(value))))
	t := mat.NewDense(n, n, value)
    unitary, e := dilation(t, degree)

    if e != nil {
        return nil, fmt.Errorf("value must represent a real contraction")
    }

    return denseToSlice(unitary), nil
}
