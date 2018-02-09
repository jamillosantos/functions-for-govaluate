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
