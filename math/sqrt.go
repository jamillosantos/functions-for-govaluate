package math

import (
	"github.com/jamillosantos/functions-for-govaluate"
	"math"
)

func Sqrt(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Sqrt(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}
