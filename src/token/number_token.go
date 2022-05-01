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

// func (nt *NumberToken) getValue() (interface{}, error) {
// 	i, err := strconv.ParseInt(nt.value, 10, 64)
// 	if err == nil {
// 		return i, nil
// 	}
// 	f, err := strconv.ParseFloat(nt.value, 64)
// 	if err == nil {
// 		return f, nil
// 	}
// 	return nil, ErrInvalidNumber
// }
