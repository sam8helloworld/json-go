package value

type String string
type NumberInt int64
type NumberFloat float64
type Bool bool

const Null = iota

type Array []interface{}
type Object map[string]interface{}
