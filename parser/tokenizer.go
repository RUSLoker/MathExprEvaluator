package parser

import (
	"fmt"
	"regexp"
)

type Named interface {
	Name() string
}

type Token interface {
	Named
	Pos() (loc []int)
	Children() []Token
}

type Parser interface {
	Named
	Parse(str string, pos int) (token Token, next int, err error)
}

type RegexParser struct {
	regex regexp.Regexp
	name  string
}

type RegexToken struct {
	name  string
	value string
	loc   []int
}

type AndGroupParser struct {
	Parsers []Parser
	name    string
}

type AndTokenGroup struct {
	tokens []Token
	name   string
}

type OrGroupParser struct {
	Parsers []Parser
	name    string
}

type OrTokenGroup struct {
	token Token
	name  string
}

func (parser *RegexParser) Parse(str string, pos int) (token Token, next int, err error) {
	regexToken := RegexToken{name: parser.name}
	regexToken.loc = parser.regex.FindStringIndex(str[pos:])

	if regexToken.loc == nil {
		return nil, pos, fmt.Errorf("can't find %s", parser.name)
	}

	regexToken.loc[0] += pos
	regexToken.loc[1] += pos

	regexToken.value = str[regexToken.loc[0]:regexToken.loc[1]]

	return &regexToken, regexToken.loc[1], nil
}

func (parser *RegexParser) Name() string {
	return parser.name
}

func (token *RegexToken) Pos() (loc []int) {
	return token.loc
}

func (token *RegexToken) Name() string {
	return token.name
}

func (token *RegexToken) Value() string {
	return token.value
}

func (token *RegexToken) Children() []Token {
	return []Token{}
}

func (parser *AndGroupParser) Parse(str string, pos int) (token Token, next int, err error) {
	andTokenGroup := AndTokenGroup{tokens: []Token{}, name: parser.name}
	next = pos
	for _, subParser := range parser.Parsers {
		token, next, err = subParser.Parse(str, next)
		if err != nil {
			//fmt.Printf("%s not success!\n", subParser.Name())
			return nil, pos, err
		}

		//fmt.Printf("%s success!\n", subParser.Name())
		andTokenGroup.tokens = append(andTokenGroup.tokens, token)
	}

	return &andTokenGroup, next, nil
}

func (parser *AndGroupParser) Name() string {
	return parser.name
}

func (token *AndTokenGroup) Pos() (loc []int) {
	if len(token.tokens) == 0 {
		return nil
	}

	begin := token.tokens[0].Pos()
	end := token.tokens[len(token.tokens)-1].Pos()
	return []int{begin[0], end[1]}
}

func (token *AndTokenGroup) Name() string {
	return token.name
}

func (token *AndTokenGroup) Children() []Token {
	return token.tokens
}

func (parser *OrGroupParser) Parse(str string, pos int) (token Token, next int, err error) {
	orTokenGroup := OrTokenGroup{name: parser.name}

	for _, subParser := range parser.Parsers {
		token, next, err = subParser.Parse(str, pos)
		if err == nil {
			//fmt.Printf("%s success!\n", subParser.Name())
			orTokenGroup.token = token
			return &orTokenGroup, next, err
		}

		//fmt.Printf("%s not success!\n", subParser.Name())
	}

	return nil, pos, fmt.Errorf("can't find %s", parser.name)
}

func (parser *OrGroupParser) Name() string {
	return parser.name
}

func (token *OrTokenGroup) Pos() (loc []int) {
	if token.token == nil {
		return nil
	}

	return token.token.Pos()
}

func (token *OrTokenGroup) Name() string {
	return token.name
}

func (token *OrTokenGroup) Children() []Token {
	return []Token{token.token}
}

func Iterate(t Token) <-chan Token {
	ch := make(chan Token)

	go func() {
		defer close(ch)
		iterate(t, ch)
	}()

	return ch
}

func iterate(t Token, ch chan Token) {
	// Send the current token
	ch <- t

	// Recursively send children
	for _, child := range t.Children() {
		iterate(child, ch)
	}
}
