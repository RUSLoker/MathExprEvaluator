package parser

import "regexp"

var WS = RegexParser{}
var ParExpr = AndGroupParser{}
var LParen = RegexParser{}
var RParen = RegexParser{}
var Expr = OrGroupParser{}
var NulOpGroup = OrGroupParser{}
var UnOpGroup = AndGroupParser{}
var BiOpsGroup = AndGroupParser{}
var Number = OrGroupParser{}
var Float = RegexParser{}
var Int = RegexParser{}
var UnOp = OrGroupParser{}
var UnPlus = RegexParser{}
var UnMinus = RegexParser{}
var BiOp = OrGroupParser{}
var BiPlus = RegexParser{}
var BiMinus = RegexParser{}
var Mul = RegexParser{}
var Div = RegexParser{}
var Pow = RegexParser{}

func init() {
	WS.name = "WS"
	WS.regex = *regexp.MustCompile(`^\s*`)

	ParExpr.name = "ParExpr"
	ParExpr.Parsers = []Parser{
		&LParen,
		&WS,
		&Expr,
		&WS,
		&RParen,
	}

	LParen.name = "LParen"
	LParen.regex = *regexp.MustCompile(`^\(`)

	RParen.name = "RParen"
	RParen.regex = *regexp.MustCompile(`^\)`)

	Expr.name = "Expr"
	Expr.Parsers = []Parser{
		&BiOpsGroup,
		&UnOpGroup,
		&NulOpGroup,
	}

	NulOpGroup.name = "NulOpGroup"
	NulOpGroup.Parsers = []Parser{
		&ParExpr,
		&Number,
	}

	UnOpGroup.name = "UnOpGroup"
	UnOpGroup.Parsers = []Parser{
		&UnOp,
		&WS,
		&NulOpGroup,
	}

	BiOpsGroup.name = "BiOpsGroup"
	BiOpsGroup.Parsers = []Parser{
		&OrGroupParser{
			name: "",
			Parsers: []Parser{
				&ParExpr,
				&UnOpGroup,
				&NulOpGroup,
			},
		},
		&WS,
		&BiOp,
		&WS,
		&Expr,
	}

	Number.name = "Number"
	Number.Parsers = []Parser{
		&Float,
		&Int,
	}

	Float.name = "Float"
	Float.regex = *regexp.MustCompile(`^[0-9]+\.[0-9]*`)

	Int.name = "Int"
	Int.regex = *regexp.MustCompile(`^[0-9]+`)

	UnOp.name = "UnOp"
	UnOp.Parsers = []Parser{
		&UnPlus,
		&UnMinus,
	}

	UnPlus.name = "UnPlus"
	UnPlus.regex = *regexp.MustCompile(`^\+`)

	UnMinus.name = "UnMinus"
	UnMinus.regex = *regexp.MustCompile(`^-`)

	BiOp.name = "BiOp"
	BiOp.Parsers = []Parser{
		&BiPlus,
		&BiMinus,
		&Mul,
		&Div,
		&Pow,
	}

	BiPlus.name = "BiPlus"
	BiPlus.regex = *regexp.MustCompile(`^\+`)

	BiMinus.name = "BiMinus"
	BiMinus.regex = *regexp.MustCompile(`^-`)

	Mul.name = "Mul"
	Mul.regex = *regexp.MustCompile(`^\*`)

	Div.name = "Div"
	Div.regex = *regexp.MustCompile(`^/`)

	Pow.name = "Pow"
	Pow.regex = *regexp.MustCompile(`^\^`)
}
