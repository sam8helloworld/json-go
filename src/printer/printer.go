package printer

import (
	"fmt"

	"github.com/sam8helloworld/json-go/value"
)

type Printer struct {
	value interface{}
}

func NewPrinter(value interface{}) *Printer {
	return &Printer{
		value: value,
	}
}

func (p *Printer) Execute() {
	print(p.value)
}

func print(val interface{}) {
	switch v := val.(type) {
	case value.NumberInt:
		fmt.Printf("%d", v)
	case value.NumberFloat:
		fmt.Printf("%f", v)
	case value.Bool:
		fmt.Printf("%t", v)
	case value.String:
		fmt.Printf("%s", v)
	case value.Array:
		fmt.Printf("[")
		for i, vi := range v {
			print(vi)
			if i != len(v)-1 {
				fmt.Printf(",")
			}
		}
		fmt.Printf("]")
	case value.Object:
		fmt.Printf("{")
		cnt := 0
		for k, vi := range v {
			fmt.Printf("\"%s\":", k)
			print(vi)
			if cnt != len(v)-1 {
				fmt.Printf(",")
			}
			cnt++
		}
		fmt.Printf("}")
	default:

	}

}
