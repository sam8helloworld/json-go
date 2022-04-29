package token

type Type string

const (
	StringType       = Type("String")
	NumberType       = Type("Number")
	TrueType         = Type("True")
	FalseType        = Type("False")
	NullType         = Type("Null")
	WhiteSpaceType   = Type("WhiteSpace")
	LeftBraceType    = Type("LeftBrace")
	RightBraceType   = Type("RightBrace")
	LeftBracketType  = Type("LeftBracket")
	RightBracketType = Type("RightBracket")
	CommaType        = Type("Comma")
	ColonType        = Type("Colon")
)

type Token struct {
	Type    Type
	Literal string
}
