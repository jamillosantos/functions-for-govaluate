package functions_for_govaluate

import (
	"gopkg.in/Knetic/govaluate.v2"
	"github.com/jamillosantos/functions-for-govaluate/math"
	"github.com/jamillosantos/functions-for-govaluate/flow"
)

var Functions map[string]govaluate.ExpressionFunction

func init() {
	Functions = make(map[string]govaluate.ExpressionFunction)
	for k, v := range math.Functions {
		Functions[k] = v
	}
	for k, v := range flow.Functions {
		Functions[k] = v
	}
}
