package lexer

import (
	"errors"
	"strconv"

	"github.com/sam8helloworld/json-go/token"
)

var (
	ErrStringTokenize = errors.New("failed to string tokenize")
	ErrBoolTokenize   = errors.New("failed to bool tokenize")
	ErrNullTokenize   = errors.New("failed to null tokenize")
	ErrNumberTokenize = errors.New("failed to number tokenize")
	ErrLexer          = errors.New("failed to lexer")
)

const (
	QuoteSymbol         = byte('"')
	LeftBraceSymbol     = byte('{')
	RightBraceSymbol    = byte('}')
	LeftBracketSymbol   = byte('[')
	RightBracketSymbol  = byte(']')
	CommaSymbol         = byte(',')
	ColonSymbol         = byte(':')
	TrueSymbol          = byte('t')
	FalseSymbol         = byte('f')
	NullSymbol          = byte('n')
	WhiteSpaceSymbol    = byte(' ')
	WhiteSpaceTabSymbol = byte('\t')
	WhiteSpaceCRSymbol  = byte('\r')
	WhiteSpaceLFSymbol  = byte('\n')
	NumberPlusSymbol    = byte('+')
	NumberMinusSymbol   = byte('-')
	NumberDotSymbol     = byte('.')
)

type Lexer struct {
	Input        string
	Position     int  // 読み込んでる文字のインデックス
	ReadPosition int  // 次に読み込む文字のインデックス
	Ch           byte // 検査中の文字
}

func NewLexer(input string) *Lexer {
	// Lexerに引数inputをセットしreturn
	return &Lexer{Input: input}
}

func (l *Lexer) Execute() (*[]token.Token, error) {
	// 1文字ずつ読み取ってその文字によってどのパースを行うか分岐
	// パースしてトークンを返す
	tokens := []token.Token{}
	for ch := l.readChar(); l.ReadPosition <= len(l.Input); ch = l.readChar() {
		switch {
		case ch == LeftBraceSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.LeftBraceType,
				Literal: string(ch),
			})
		case ch == RightBraceSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.RightBraceType,
				Literal: string(ch),
			})
		case ch == LeftBracketSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.LeftBracketType,
				Literal: string(ch),
			})
		case ch == RightBracketSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.RightBracketType,
				Literal: string(ch),
			})
		case ch == ColonSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.ColonType,
				Literal: string(ch),
			})
		case ch == CommaSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.CommaType,
				Literal: string(ch),
			})
		case ch == TrueSymbol:
			token, err := l.boolTokenize(true)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		case ch == FalseSymbol:
			token, err := l.boolTokenize(false)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		case ch == NullSymbol:
			token, err := l.nullTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		case ch == WhiteSpaceSymbol, ch == WhiteSpaceTabSymbol, ch == WhiteSpaceTabSymbol, ch == WhiteSpaceCRSymbol, ch == WhiteSpaceLFSymbol:
			continue
		case ch == QuoteSymbol:
			token, err := l.stringTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		case '0' <= ch && ch <= '9', ch == NumberPlusSymbol, ch == NumberMinusSymbol, ch == NumberDotSymbol:
			// Numberは開始文字が[0-9]もしくは('+', '-', '.')
			// e.g.
			//     -1235
			//     +10
			//     .00001
			token, err := l.numberTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		default:
			return nil, ErrLexer
		}
	}
	return &tokens, nil
}

func (l *Lexer) readChar() byte {
	// 入力が終わったらchを0に
	if l.ReadPosition >= len(l.Input) {
		l.Ch = 0
	} else {
		// まだ終わっていない場合readPositionをchにセット
		l.Ch = l.Input[l.ReadPosition]
	}
	// positionを次に進める
	l.Position = l.ReadPosition
	// readpositonを次に進める
	l.ReadPosition += 1
	return l.Ch
}

func (l *Lexer) peakChar() byte {
	// 入力が終わったらchを0に
	if l.ReadPosition >= len(l.Input) {
		return 0
	} else {
		return l.Input[l.ReadPosition]
	}
}

func (l *Lexer) stringTokenize() (*token.Token, error) {
	str := ""
	for ch := l.readChar(); ch != 0; ch = l.readChar() {
		if ch == QuoteSymbol {
			return &token.Token{
				Type:    token.StringType,
				Literal: str,
			}, nil
		}
		str += string(ch)
	}
	return nil, ErrStringTokenize
}

func (l *Lexer) boolTokenize(b bool) (*token.Token, error) {
	s := string(l.Ch)
	if b {
		for i := 0; i < 3; i++ {
			s += string(l.readChar())
		}
		if s == "true" {
			return &token.Token{
				Type:    token.TrueType,
				Literal: s,
			}, nil
		}
		return nil, ErrBoolTokenize
	}
	for i := 0; i < 4; i++ {
		s += string(l.readChar())
	}
	if s == "false" {
		return &token.Token{
			Type:    token.FalseType,
			Literal: s,
		}, nil
	}
	return nil, ErrBoolTokenize
}

func (l *Lexer) nullTokenize() (*token.Token, error) {
	s := string(l.Ch)
	for i := 0; i < 3; i++ {
		s += string(l.readChar())
	}
	if s == "null" {
		return &token.Token{
			Type:    token.NullType,
			Literal: s,
		}, nil
	}
	return nil, ErrNullTokenize
}

func (l *Lexer) numberTokenize() (*token.Token, error) {
	num := string(l.Ch)
	for {
		ch := l.peakChar()
		if isNumberSymbol(ch) {
			num += string(ch)
			l.readChar()
		} else {
			break
		}
	}
	_, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return nil, ErrNumberTokenize
	}
	return &token.Token{
		Type:    token.NumberType,
		Literal: num,
	}, nil
}

func isNumberSymbol(s byte) bool {
	// 数字に使いそうな文字は全て読み込む
	// 1e10, 1E10, 1.0000
	if ('0' <= s && s <= '9') || s == NumberPlusSymbol || s == NumberMinusSymbol || s == NumberDotSymbol || s == 'e' || s == 'E' {
		return true
	}
	return false
}
