package flow

import "gopkg.in/Knetic/govaluate.v2"

var Functions map[string]govaluate.ExpressionFunction = map[string]govaluate.ExpressionFunction{
	"isDefined": IsDefined,
	"ifnull":    IfNull,
	"if":        If,
}
