package token

type NumberToken struct {
	Token
	value string
}

func NewNumberToken(value string) NumberToken {
	return NumberToken{
		value: value,
	}
}

func (nt *NumberToken) Value() string {
	return nt.value
}
