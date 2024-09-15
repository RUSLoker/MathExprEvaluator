package evaluator

import (
	"fmt"
	"parser"
	"reflect"
	"strconv"
)

func parseNumber(number parser.Token) (any, error) {
	value, ok := number.(*parser.RegexToken)

	if !ok {
		return nil, fmt.Errorf(`expected parser.RegexToken got %s`, reflect.TypeOf(number))
	}

	switch value.Name() {
	case "Float":
		return strconv.ParseFloat(value.Value(), 64)
	case "Int":
		return strconv.ParseInt(value.Value(), 10, 64)
	default:
		return nil, fmt.Errorf(`expected "Float" or "Int" got %s`, value.Name())
	}
}
