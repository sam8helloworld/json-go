package token

type StringToken struct {
	Token
	value string
}

func NewStringToken(value string) StringToken {
	return StringToken{
		value: value,
	}
}

func (st *StringToken) Value() string {
	return st.value
}
