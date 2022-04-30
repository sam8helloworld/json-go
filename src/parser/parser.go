package parser

import (
	"errors"

	"github.com/sam8helloworld/json-go/token"
	"github.com/sam8helloworld/json-go/value"
)

var (
	ErrNotStartWithLeftBrace   = errors.New("not start with '{'")
	ErrNotStartWithLeftBracket = errors.New("not start with '['")
	ErrInvalidKeyValuePair     = errors.New("invalid key value pair")
	ErrInvalidArrayValue       = errors.New("invalid array value")
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
	switch t.Type {
	case token.LeftBraceType:
		return p.parseObject()
	case token.LeftBracketType:
		return p.parseArray()
	case token.StringType:
		p.next()
		return value.String(t.Expression.(string)), nil
	case token.NumberType:
		p.next()
		return value.Number(t.Expression.(float64)), nil
	case token.FalseType, token.TrueType:
		p.next()
		return value.Bool(t.Expression.(bool)), nil
	case token.NullType:
		p.next()
		return value.Null(t.Expression), nil
	default:
		return nil, ErrParse
	}
}

func (p *Parser) parseObject() (value.Object, error) {
	t := p.peek()
	if t.Type != token.LeftBraceType {
		return nil, ErrNotStartWithLeftBrace
	}
	// { を読み飛ばす
	p.next()

	object := value.Object{}

	if t.Type == token.RightBraceType {
		return object, nil
	}

	for {
		t1 := p.next()
		t2 := p.next()

		if t1.Type == token.StringType && t2.Type == token.ColonType {
			v, err := p.parse()
			if err != nil {
				return nil, err
			}
			object[t1.Expression.(string)] = v
		} else {
			return nil, ErrInvalidKeyValuePair
		}

		t3 := p.next()
		if t3.Type == token.RightBraceType {
			return object, nil
		}
		if t3.Type == token.CommaType {
			continue
		}
		return nil, ErrParse
	}
}

func (p *Parser) parseArray() (value.Array, error) {
	// 先頭は必ず [
	t := p.peek()
	if t.Type != token.LeftBracketType {
		return nil, ErrNotStartWithLeftBracket
	}
	// [ を読み飛ばす
	p.next()

	array := value.Array{}
	t = p.peek()
	// ] なら空配列を返す
	if t.Type == token.RightBracketType {
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
		if t.Type == token.RightBracketType {
			return array, nil
		}
		// , なら次の要素(`Value`)をパースする
		if t.Type == token.CommaType {
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
