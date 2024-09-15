package evaluator

import (
	"fmt"
	"math"
	"reflect"
)

var operators = map[string]operator{
	"Pow": {
		Precedence:    1,
		Associativity: RightAssociativity,
		Arity:         2,
		Apply: func(ops []any) (any, error) {
			lhsInt, isLhsInt := ops[0].(int64)
			lhsFloat, isLhsFloat := ops[0].(float64)
			rhsInt, isRhsInt := ops[1].(int64)
			rhsFloat, isRhsFloat := ops[1].(float64)
			if isLhsFloat && isRhsFloat {
				return math.Pow(lhsFloat, rhsFloat), nil
			}
			if isLhsFloat && isRhsInt {
				return math.Pow(lhsFloat, float64(rhsInt)), nil
			}
			if isLhsInt && isRhsFloat {
				return math.Pow(float64(lhsInt), rhsFloat), nil
			}
			if isLhsInt && isRhsInt {
				return math.Pow(float64(lhsInt), float64(rhsInt)), nil
			}
			return nil, fmt.Errorf(
				"unsupported operation '^' for %s and %s",
				reflect.TypeOf(ops[0]), reflect.TypeOf(ops[1]),
			)
		},
	},
	"UnPlus": {
		Precedence:    2,
		Associativity: RightAssociativity,
		Arity:         1,
		Apply: func(ops []any) (any, error) {
			return ops[0], nil
		},
	},
	"UnMinus": {
		Precedence:    2,
		Associativity: RightAssociativity,
		Arity:         1,
		Apply: func(ops []any) (any, error) {
			switch ops[0].(type) {
			case int64:
				return -ops[0].(int64), nil
			case float64:
				return -ops[0].(float64), nil
			default:
				return nil, fmt.Errorf("unsupported operation '-' for %s", reflect.TypeOf(ops[0]))
			}
		},
	},
	"Mul": {
		Precedence:    3,
		Associativity: LeftAssociativity,
		Arity:         2,
		Apply: func(ops []any) (any, error) {
			lhsInt, isLhsInt := ops[0].(int64)
			lhsFloat, isLhsFloat := ops[0].(float64)
			rhsInt, isRhsInt := ops[1].(int64)
			rhsFloat, isRhsFloat := ops[1].(float64)
			if isLhsFloat && isRhsFloat {
				return lhsFloat * rhsFloat, nil
			}
			if isLhsFloat && isRhsInt {
				return lhsFloat * float64(rhsInt), nil
			}
			if isLhsInt && isRhsFloat {
				return float64(lhsInt) * rhsFloat, nil
			}
			if isLhsInt && isRhsInt {
				return lhsInt * rhsInt, nil
			}
			return nil, fmt.Errorf(
				"unsupported operation '*' for %s and %s",
				reflect.TypeOf(ops[0]), reflect.TypeOf(ops[1]),
			)
		},
	},
	"Div": {
		Precedence:    3,
		Associativity: LeftAssociativity,
		Arity:         2,
		Apply: func(ops []any) (any, error) {
			lhsInt, isLhsInt := ops[0].(int64)
			lhsFloat, isLhsFloat := ops[0].(float64)
			rhsInt, isRhsInt := ops[1].(int64)
			rhsFloat, isRhsFloat := ops[1].(float64)
			if isLhsFloat && isRhsFloat {
				return lhsFloat / rhsFloat, nil
			}
			if isLhsFloat && isRhsInt {
				return lhsFloat / float64(rhsInt), nil
			}
			if isLhsInt && isRhsFloat {
				return float64(lhsInt) / rhsFloat, nil
			}
			if isLhsInt && isRhsInt {
				return lhsInt / rhsInt, nil
			}
			return nil, fmt.Errorf(
				"unsupported operation '/' for %s and %s",
				reflect.TypeOf(ops[0]), reflect.TypeOf(ops[1]),
			)
		},
	},
	"BiPlus": {
		Precedence:    4,
		Associativity: LeftAssociativity,
		Arity:         2,
		Apply: func(ops []any) (any, error) {
			lhsInt, isLhsInt := ops[0].(int64)
			lhsFloat, isLhsFloat := ops[0].(float64)
			rhsInt, isRhsInt := ops[1].(int64)
			rhsFloat, isRhsFloat := ops[1].(float64)
			if isLhsFloat && isRhsFloat {
				return lhsFloat + rhsFloat, nil
			}
			if isLhsFloat && isRhsInt {
				return lhsFloat + float64(rhsInt), nil
			}
			if isLhsInt && isRhsFloat {
				return float64(lhsInt) + rhsFloat, nil
			}
			if isLhsInt && isRhsInt {
				return lhsInt + rhsInt, nil
			}
			return nil, fmt.Errorf(
				"unsupported operation '+' for %s and %s",
				reflect.TypeOf(ops[0]), reflect.TypeOf(ops[1]),
			)
		},
	},
	"BiMinus": {
		Precedence:    4,
		Associativity: LeftAssociativity,
		Arity:         2,
		Apply: func(ops []any) (any, error) {
			lhsInt, isLhsInt := ops[0].(int64)
			lhsFloat, isLhsFloat := ops[0].(float64)
			rhsInt, isRhsInt := ops[1].(int64)
			rhsFloat, isRhsFloat := ops[1].(float64)
			if isLhsFloat && isRhsFloat {
				return lhsFloat - rhsFloat, nil
			}
			if isLhsFloat && isRhsInt {
				return lhsFloat - float64(rhsInt), nil
			}
			if isLhsInt && isRhsFloat {
				return float64(lhsInt) - rhsFloat, nil
			}
			if isLhsInt && isRhsInt {
				return lhsInt - rhsInt, nil
			}
			return nil, fmt.Errorf(
				"unsupported operation '-' for %s and %s",
				reflect.TypeOf(ops[0]), reflect.TypeOf(ops[1]),
			)
		},
	},
}
