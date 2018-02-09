package math

import (
	"github.com/jamillosantos/functions-for-govaluate"
	"math"
)

func Floor(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Floor(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Ceil(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Ceil(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Round(args ...interface{}) (interface{}, error) {
	switch len(args) {
	case 1:
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		if math.IsNaN(v) {
			return math.NaN(), nil
		}
		if v > 0.5 {
			return math.Trunc(v + 0.5), nil
		} else if v < -0.5 {
			return math.Trunc(v - 0.5), nil
		}
		return 0.0, nil
	case 2:
		v1, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		v2, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(1)
		}

		var round float64
		pow := math.Pow(10, float64(v2))
		digit := pow * v1
		_, div := math.Modf(digit)
		if div >= 0.5 {
			round = math.Ceil(digit)
		} else {
			round = math.Floor(digit)
		}
		return round / pow, nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Trunc(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Trunc(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}
