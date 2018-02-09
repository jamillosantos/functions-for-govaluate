package math

import "gopkg.in/Knetic/govaluate.v2"

var Functions map[string]govaluate.ExpressionFunction = map[string]govaluate.ExpressionFunction{
	"cos":   Cos,
	"acos":  Acos,
	"cosh":  Cosh,
	"acosh": Acosh,

	"exp":  Exp,
	"exp2": Exp2,

	"log":   Log,
	"log10": Log10,

	"round": Round,
	"floor": Floor,
	"ceil":  Ceil,
	"trunc": Trunc,

	"sin":   Sin,
	"asin":  Asin,
	"sinh":  Sinh,
	"asinh": Asinh,

	"sqrt": Sqrt,

	"tan":   Tan,
	"atan":  Atan,
	"atan2": Atan2,
	"tanh":  Tanh,
	"atanh": Atanh,
}

func ToFloat64(v interface{}) (float64, bool) {
	switch vv := v.(type) {
	case float64:
		return vv, true
	case float32:
		return float64(vv), true
	case int8:
		return float64(vv), true
	case int32:
		return float64(vv), true
	case int64:
		return float64(vv), true
	case uint8:
		return float64(vv), true
	case uint32:
		return float64(vv), true
	case uint64:
		return float64(vv), true
	case int:
		return float64(vv), true
	case uint:
		return float64(vv), true
	default:
		return 0, false
	}
}
