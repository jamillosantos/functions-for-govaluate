package functions_for_govaluate

func IfNull(args ...interface{}) (interface{}, error) {
	for i, v := range args {
		condition := false
		switch cond := v.(type) {
		case string:
			condition = cond != ""
		case float64:
			condition = cond != 0
		default:
			return nil, NewWrongParamType(i)
		}
		if condition {
			return v, nil
		}
	}
	return nil, nil
}
