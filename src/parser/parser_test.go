package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/json-go/token"
	"github.com/sam8helloworld/json-go/value"
)

func TestSuccess(t *testing.T) {
	tests := []struct {
		name  string
		input []token.Token
		want  interface{}
	}{
		{
			name: "文字列のみ",
			input: []token.Token{
				{
					Type:       token.LeftBraceType,
					Expression: "{",
				},
				{
					Type:       token.StringType,
					Expression: "key",
				},
				{
					Type:       token.ColonType,
					Expression: ":",
				},
				{
					Type:       token.StringType,
					Expression: "value",
				},
				{
					Type:       token.RightBraceType,
					Expression: "}",
				},
			},
			want: value.Object{
				"key": value.String("value"),
			},
		},
		{
			name: "数値のみ",
			input: []token.Token{
				{
					Type:       token.LeftBraceType,
					Expression: "{",
				},
				{
					Type:       token.StringType,
					Expression: "key",
				},
				{
					Type:       token.ColonType,
					Expression: ":",
				},
				{
					Type:       token.NumberType,
					Expression: "100",
				},
				{
					Type:       token.RightBraceType,
					Expression: "}",
				},
			},
			want: value.Object{
				"key": value.Number(int64(100)),
			},
		},
		{
			name: "boolのみ",
			input: []token.Token{
				{
					Type:       token.LeftBraceType,
					Expression: "{",
				},
				{
					Type:       token.StringType,
					Expression: "key",
				},
				{
					Type:       token.ColonType,
					Expression: ":",
				},
				{
					Type:       token.TrueType,
					Expression: "true",
				},
				{
					Type:       token.RightBraceType,
					Expression: "}",
				},
			},
			want: value.Object{
				"key": value.Bool(true),
			},
		},
		{
			name: "nullのみ",
			input: []token.Token{
				{
					Type:       token.LeftBraceType,
					Expression: "{",
				},
				{
					Type:       token.StringType,
					Expression: "key",
				},
				{
					Type:       token.ColonType,
					Expression: ":",
				},
				{
					Type:       token.NullType,
					Expression: "null",
				},
				{
					Type:       token.RightBraceType,
					Expression: "}",
				},
			},
			want: value.Object{
				"key": value.Null("null"),
			},
		},
		{
			name: "objectがネスト",
			input: []token.Token{
				{
					Type:       token.LeftBraceType,
					Expression: "{",
				},
				{
					Type:       token.StringType,
					Expression: "key",
				},
				{
					Type:       token.ColonType,
					Expression: ":",
				},
				{
					Type:       token.LeftBraceType,
					Expression: "{",
				},
				{
					Type:       token.StringType,
					Expression: "key",
				},
				{
					Type:       token.ColonType,
					Expression: ":",
				},
				{
					Type:       token.StringType,
					Expression: "value",
				},
				{
					Type:       token.RightBraceType,
					Expression: "}",
				},
				{
					Type:       token.RightBraceType,
					Expression: "}",
				},
			},
			want: value.Object{
				"key": value.Object{
					"key": value.String("value"),
				},
			},
		},
		{
			name: "arrayのみ",
			input: []token.Token{
				{
					Type:       token.LeftBraceType,
					Expression: "{",
				},
				{
					Type:       token.StringType,
					Expression: "key",
				},
				{
					Type:       token.ColonType,
					Expression: ":",
				},
				{
					Type:       token.LeftBracketType,
					Expression: "[",
				},
				{
					Type:       token.StringType,
					Expression: "value1",
				},
				{
					Type:       token.CommaType,
					Expression: ",",
				},
				{
					Type:       token.StringType,
					Expression: "value2",
				},
				{
					Type:       token.RightBracketType,
					Expression: "]",
				},
				{
					Type:       token.RightBraceType,
					Expression: "}",
				},
			},
			want: value.Object{
				"key": value.Array{
					value.String("value1"),
					value.String("value2"),
				},
			},
		},
		{
			name: "トップレベルがarray",
			input: []token.Token{
				{
					Type:       token.LeftBracketType,
					Expression: "[",
				},
				{
					Type:       token.StringType,
					Expression: "value1",
				},
				{
					Type:       token.CommaType,
					Expression: ",",
				},
				{
					Type:       token.StringType,
					Expression: "value2",
				},
				{
					Type:       token.RightBracketType,
					Expression: "]",
				},
			},
			want: value.Array{
				value.String("value1"),
				value.String("value2"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := NewParser(tt.input)
			got, err := sut.Execute()
			if err != nil {
				t.Fatalf("failed to execute parser %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}
