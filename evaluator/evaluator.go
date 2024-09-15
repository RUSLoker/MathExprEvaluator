package evaluator

import (
	"fmt"
	"parser"
	"reflect"
)

type Associativity int8

const (
	LeftAssociativity Associativity = iota
	RightAssociativity
)

type Number interface {
	int64 | float64
}

type operator struct {
	Precedence    int8
	Associativity Associativity
	Arity         uint8
	Apply         func([]any) (any, error)
}

type lParen struct{}

type evaluator struct {
	operators Stack[any]
	operands  Stack[any]
}

func (eval *evaluator) EvaluateAst(tree parser.Token) (any, error) {
	eval.operators = Stack[any]{}
	eval.operands = Stack[any]{}

	eval.operators.Push(lParen{})

	for cur := range parser.Iterate(tree) {
		switch cur.Name() {
		case "Float":
			fallthrough
		case "Int":
			result, err := parseNumber(cur)
			if err != nil {
				return nil, err
			}
			eval.operands.Push(result)

		case "LParen":
			eval.operators.Push(lParen{})
		case "RParen":
			err := eval.applyTillParen()
			if err != nil {
				return nil, err
			}

		default:
			operA, exists := operators[cur.Name()]
			if !exists {
				break
			}
			if operA.Arity == 1 {
				eval.operators.Push(operA)
			} else if operA.Arity == 2 {
				operStackTop, hasTopOper := eval.operators.Peek()
				operB, isBOper := (*operStackTop).(operator)

				if hasTopOper && isBOper && (operA.Associativity == LeftAssociativity &&
					operB.Precedence <= operA.Precedence ||
					operA.Associativity == RightAssociativity &&
						operB.Precedence < operA.Precedence) {
					err := eval.applyOne()
					if err != nil {
						return nil, err
					}
				}
				eval.operators.Push(operA)
			} else {
				return nil, fmt.Errorf("unknown operator")
			}
		}
	}

	err := eval.applyTillParen()
	if err != nil {
		return nil, err
	}

	if !eval.operators.IsEmpty() {
		return nil, fmt.Errorf("unevaluated operators remained")
	}

	if eval.operands.Size() > 1 {
		return nil, fmt.Errorf("unevaluated operands remained")
	}

	result, exists := eval.operands.Pop()

	if !exists {
		return nil, fmt.Errorf("no result")
	}

	return *result, nil
}

func (eval *evaluator) applyOne() error {
	stackTop, exists := eval.operators.Pop()
	if !exists {
		return fmt.Errorf("no operators in the queue")
	}

	oper, ok := (*stackTop).(operator)
	if !ok {
		return fmt.Errorf("expected operator, but got %s", reflect.TypeOf(stackTop))
	}

	operands := eval.operands.PopN(int(oper.Arity))
	if len(operands) != int(oper.Arity) {
		return fmt.Errorf("expected %s operands, but got %s", oper.Arity, len(operands))
	}

	result, err := oper.Apply(operands)
	if err != nil {
		return err
	}
	eval.operands.Push(result)

	return nil
}

func (eval *evaluator) applyTillParen() error {
	top, exists := eval.operators.Peek()
	for ; exists; top, exists = eval.operators.Peek() {
		if _, ok := (*top).(lParen); ok {
			break
		}
		err := eval.applyOne()
		if err != nil {
			return err
		}
	}
	if _, ok := (*top).(lParen); !ok {
		return fmt.Errorf("opening bracket expected")
	}
	eval.operators.Pop()
	return nil
}

// Evaluate receives string with a math expression as the input and evaluates it to int64 or float64
func Evaluate(input string) (any, error) {
	result, next, err := parser.Expr.Parse(input, 0)
	if err != nil {
		return nil, err
	}
	if next != len(input) {
		return nil, fmt.Errorf("Next was %d, expected %d\n", next, len(input))
	}
	result.Pos()
	e := evaluator{}
	res, err := e.EvaluateAst(result)
	if err != nil {
		return nil, err
	}
	return res, nil
}
