package functions_for_govaluate

func IsDefined(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		if args[0] != nil {
			switch cond := args[0].(type) {
			case string:
				return cond != "", nil
			case float64:
				return cond != 0.0, nil
			default:
				return nil, NewWrongParamType(0)
			}
		}
		return false, nil
	}
	return nil, WrongParamsCount
}
