package lexer

import (
	"errors"

	"github.com/sam8helloworld/json-go/token"
)

var (
	ErrStringTokenize = errors.New("failed to string tokenize")
)

const (
	QuoteSymbol        = byte('"')
	LeftBraceSymbol    = byte('{')
	RightBraceSymbol   = byte('}')
	LeftBracketSymbol  = byte('[')
	RightBracketSymbol = byte(']')
	CommaSymbol        = byte(',')
	ColonSymbol        = byte(':')
	TrueSymbol         = byte('t')
	FalseSymbol        = byte('f')
	NullSymbol         = byte('n')
	WhiteSpace         = byte(' ')
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
	for t := l.readChar(); l.ReadPosition <= len(l.Input); t = l.readChar() {
		switch t {
		case LeftBraceSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.LeftBraceType,
				Literal: string(t),
			})
		case RightBraceSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.RightBraceType,
				Literal: string(t),
			})
		case LeftBracketSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.LeftBracketType,
				Literal: string(t),
			})
		case RightBracketSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.RightBracketType,
				Literal: string(t),
			})
		case ColonSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.ColonType,
				Literal: string(t),
			})
		case CommaSymbol:
			tokens = append(tokens, token.Token{
				Type:    token.CommaType,
				Literal: string(t),
			})
		case WhiteSpace:
			continue
		case QuoteSymbol:
			token, err := l.stringTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		default:
			continue
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
