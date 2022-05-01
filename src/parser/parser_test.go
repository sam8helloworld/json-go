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
				token.LeftBraceToken{},
				token.NewStringToken("key"),
				token.ColonToken{},
				token.NewStringToken("value"),
				token.RightBraceToken{},
			},
			want: value.Object{
				"key": value.String("value"),
			},
		},
		{
			name: "数値のみ",
			input: []token.Token{
				token.LeftBraceToken{},
				token.NewStringToken("key"),
				token.ColonToken{},
				token.NewNumberToken("100"),
				token.RightBraceToken{},
			},
			want: value.Object{
				"key": value.NumberInt(int64(100)),
			},
		},
		{
			name: "boolのみ",
			input: []token.Token{
				token.LeftBraceToken{},
				token.NewStringToken("key"),
				token.ColonToken{},
				token.TrueToken{},
				token.RightBraceToken{},
			},
			want: value.Object{
				"key": value.Bool(true),
			},
		},
		{
			name: "nullのみ",
			input: []token.Token{
				token.LeftBraceToken{},
				token.NewStringToken("key"),
				token.ColonToken{},
				token.NullToken{},
				token.RightBraceToken{},
			},
			want: value.Object{
				"key": value.Null,
			},
		},
		{
			name: "objectがネスト",
			input: []token.Token{
				token.LeftBraceToken{},
				token.NewStringToken("key"),
				token.ColonToken{},
				token.LeftBraceToken{},
				token.NewStringToken("key"),
				token.ColonToken{},
				token.NewStringToken("value"),
				token.RightBraceToken{},
				token.RightBraceToken{},
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
				token.LeftBraceToken{},
				token.NewStringToken("key"),
				token.ColonToken{},
				token.LeftBracketToken{},
				token.NewStringToken("value1"),
				token.CommaToken{},
				token.NewStringToken("value2"),
				token.RightBracketToken{},
				token.RightBraceToken{},
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
				token.LeftBracketToken{},
				token.NewStringToken("value1"),
				token.CommaToken{},
				token.NewStringToken("value2"),
				token.RightBracketToken{},
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
