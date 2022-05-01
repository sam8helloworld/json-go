package parser

import (
	"errors"
	"strconv"

	"github.com/sam8helloworld/json-go/token"
	"github.com/sam8helloworld/json-go/value"
)

var (
	ErrNotStartWithLeftBrace   = errors.New("not start with '{'")
	ErrNotStartWithLeftBracket = errors.New("not start with '['")
	ErrInvalidKeyValuePair     = errors.New("invalid key value pair")
	ErrInvalidArrayValue       = errors.New("invalid array value")
	ErrInvalidNumberValue      = errors.New("invalid number value")
	ErrInvalidBoolValue        = errors.New("invalid bool value")
	ErrParse                   = errors.New("failed to parse")
)

type Parser struct {
	Tokens []token.Token
	index  int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		Tokens: tokens,
	}
}

func (p *Parser) Execute() (interface{}, error) {
	values, err := p.parse()
	return values, err
}

func (p *Parser) parse() (interface{}, error) {
	t := p.peek()
	switch t := t.(type) {
	case token.LeftBraceToken:
		return p.parseObject()
	case token.LeftBracketToken:
		return p.parseArray()
	case token.StringToken:
		p.next()
		return value.String(t.Value()), nil
	case token.NumberToken:
		p.next()
		i, err := strconv.ParseInt(t.Value(), 10, 64)
		if err == nil {
			return value.NumberInt(i), nil
		}
		f, err := strconv.ParseFloat(t.Value(), 64)
		if err == nil {
			return value.NumberFloat(f), nil
		}
		return nil, ErrInvalidNumberValue
	case token.TrueToken:
		p.next()
		return value.Bool(true), nil
	case token.FalseToken:
		p.next()
		return value.Bool(false), nil
	case token.NullToken:
		p.next()
		return value.Null, nil
	default:
		return nil, ErrParse
	}
}

func (p *Parser) parseObject() (value.Object, error) {
	t := p.peek()

	switch t.(type) {
	case token.LeftBraceToken:
	default:
		return nil, ErrNotStartWithLeftBrace
	}
	// { を読み飛ばす
	p.next()

	object := value.Object{}

	switch t.(type) {
	case token.RightBraceToken:
		return object, nil
	}

	for {
		t1 := p.next()
		t2 := p.next()

		t1t, t1Ok := t1.(token.StringToken)
		_, t2Ok := t2.(token.ColonToken)

		if t1Ok && t2Ok {
			v, err := p.parse()
			if err != nil {
				return nil, err
			}
			object[t1t.Value()] = v
		} else {
			return nil, ErrInvalidKeyValuePair
		}

		t3 := p.next()
		switch t3.(type) {
		case token.RightBraceToken:
			return object, nil
		case token.CommaToken:
			continue
		}

		return nil, ErrParse
	}
}

func (p *Parser) parseArray() (value.Array, error) {
	// 先頭は必ず [
	t := p.peek()

	_, ok := t.(token.LeftBracketToken)
	if !ok {
		return nil, ErrNotStartWithLeftBracket
	}

	// [ を読み飛ばす
	p.next()

	array := value.Array{}
	t = p.peek()
	// ] なら空配列を返す
	switch t.(type) {
	case token.RightBracketToken:
		return array, nil
	}

	for {
		// 残りの`Value`をパースする
		value, err := p.parse()
		if err != nil {
			return nil, ErrInvalidArrayValue
		}
		array = append(array, value)

		t = p.next()
		// `Array`が終端もしくは次の要素(`Value`)があるかを確認
		switch t.(type) {
		case token.RightBracketToken:
			return array, nil
			// , なら次の要素(`Value`)をパースする
		case token.CommaToken:
			continue
		}
		return nil, ErrParse
	}
}

func (p *Parser) peek() token.Token {
	return p.Tokens[p.index]
}

func (p *Parser) next() token.Token {
	p.index += 1
	return p.Tokens[p.index-1]
}
