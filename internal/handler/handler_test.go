package handler

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
    "reflect"
    "testing"
)

func testUnitaryNDilation(test *testing.T, expectedT *mat.Dense, expectedN int, errorToThrow error) unitaryNDilation {
    return func(t *mat.Dense, n int) (*mat.Dense, error) {
        if !mat.Equal(t, expectedT) {
            test.Errorf("unexpected argument t in call to unitaryNDilation. Got: %v, want: %v", t, expectedT)
        }

        if n != expectedN {
            test.Errorf("unexpected argument n in call to unitaryNDilation. Got: %d, want: %d", n, expectedN)
        }

        return mat.NewDense(2, 2, nil), errorToThrow
    }
}

func TestDilationHandlerPost(t *testing.T) {
    tables := []struct {
        desc string
		value []float64
		degree int
        errorToThrow error
        expectedT *mat.Dense
        expectedN int
		expectedResult []float64
		expectedError error
    }{
        {
            desc: "returns n-dilation",
			value: []float64{0,0,0,0,},
			degree: 2,
            expectedT: mat.NewDense(2, 2, nil),
            expectedN: 2,
            expectedResult: []float64{0,0,0,0,},
			expectedError: nil,
        },
        {
            desc: "returns error if degree is not greater than 0",
			value: []float64{0,0,0,0,},
			degree: 0,
            expectedT: mat.NewDense(2, 2, nil),
            expectedN: 0,
            expectedResult: nil,
			expectedError: fmt.Errorf("degree must be an integer greater than zero"),
        },
        {
            desc: "returns error if value does not represent a square matrix",
			value: []float64{0,0,0,0,0,},
			degree: 2,
            expectedT: mat.NewDense(2, 2, nil),
            expectedN: 2,
            expectedResult: nil,
			expectedError: fmt.Errorf("value must contain a square number greater than zero of numbers"),
        },
        {
            desc: "returns error if UnitaryNDilation returns an error",
			value: []float64{0,0,0,0,},
			degree: 2,
            expectedT: mat.NewDense(2, 2, nil),
            expectedN: 2,
            errorToThrow: fmt.Errorf("test-error"),
			expectedResult: nil,
			expectedError: fmt.Errorf("value must represent a real contraction"),
        },
    }

    for _, table := range tables {
        table := table
        t.Run(table.desc, func(t *testing.T) {
			t.Parallel()

			dilation, err := DilationHandler(testUnitaryNDilation(t, table.expectedT, table.expectedN, table.errorToThrow), table.value, table.degree)

			if !reflect.DeepEqual(table.expectedError, err) {
				t.Errorf("handler returned error: got %v, want %v", err, table.expectedError)
			}

			if !reflect.DeepEqual(dilation, table.expectedResult) {
				t.Errorf("handler return wrong result: got %v, want %v", dilation, table.expectedResult)
			}
        })
    }
}
