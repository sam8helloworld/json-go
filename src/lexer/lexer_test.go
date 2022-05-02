package lexer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/json-go/token"
)

func TestSuccessStringTokenize(t *testing.T) {
	f, err := os.Open("./testdata/string_only.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if err != nil {
		t.Fatalf("failed to execute lexer %#v", err)
	}
	want := &[]token.Token{
		token.LeftBraceToken{},
		token.NewStringToken("string"),
		token.ColonToken{},
		token.NewStringToken("hogehoge"),
		token.RightBraceToken{},
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(token.StringToken{})); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestSuccessStringTokenizeEscape(t *testing.T) {
	f, err := os.Open("./testdata/escape_string_only.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if err != nil {
		t.Fatalf("failed to execute lexer %#v", err)
	}
	want := &[]token.Token{
		token.LeftBraceToken{},
		token.NewStringToken("escape_double_quote"),
		token.ColonToken{},
		token.NewStringToken("\"\""),
		token.RightBraceToken{},
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(token.StringToken{})); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestFailedStringTokenize(t *testing.T) {
	f, err := os.Open("./testdata/string_only_fragile.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if got != nil {
		t.Errorf("want error %v, but got result %v", ErrStringTokenize, got)
	}
	if !errors.Is(err, ErrStringTokenize) {
		t.Fatalf("want ErrStringTokenize, but got %v", err)
	}
}

func TestSuccessBoolTokenize(t *testing.T) {
	f, err := os.Open("./testdata/bool_only.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if err != nil {
		t.Fatalf("failed to execute lexer %#v", err)
	}
	want := &[]token.Token{
		token.LeftBraceToken{},
		token.NewStringToken("boolTrue"),
		token.ColonToken{},
		token.TrueToken{},
		token.CommaToken{},
		token.NewStringToken("boolFalse"),
		token.ColonToken{},
		token.FalseToken{},
		token.RightBraceToken{},
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(token.StringToken{})); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestFailedBoolTokenize(t *testing.T) {
	f, err := os.Open("./testdata/bool_only_fragile.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if got != nil {
		t.Errorf("want error %v, but got result %v", ErrBoolTokenize, got)
	}
	if !errors.Is(err, ErrBoolTokenize) {
		t.Fatalf("want ErrBoolTokenize, but got %v", err)
	}
}

func TestSuccessNullTokenize(t *testing.T) {
	f, err := os.Open("./testdata/null_only.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if err != nil {
		t.Fatalf("failed to execute lexer %#v", err)
	}
	want := &[]token.Token{
		token.LeftBraceToken{},
		token.NewStringToken("null"),
		token.ColonToken{},
		token.NullToken{},
		token.RightBraceToken{},
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(token.StringToken{})); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestFailedNullTokenize(t *testing.T) {
	f, err := os.Open("./testdata/null_only_fragile.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if got != nil {
		t.Errorf("want error %v, but got result %v", ErrNullTokenize, got)
	}
	if !errors.Is(err, ErrNullTokenize) {
		t.Fatalf("want ErrNullTokenize, but got %v", err)
	}
}

func TestSuccessNumberTokenize(t *testing.T) {
	f, err := os.Open("./testdata/number_only.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if err != nil {
		t.Fatalf("failed to execute lexer %#v", err)
	}
	want := &[]token.Token{
		token.LeftBraceToken{},
		// "int": 100
		token.NewStringToken("int"),
		token.ColonToken{},
		token.NewNumberToken("100"),
		token.CommaToken{},
		// "float": 1.234
		token.NewStringToken("float"),
		token.ColonToken{},
		token.NewNumberToken("1.234"),
		token.CommaToken{},
		// "floatDotStart": .1234
		token.NewStringToken("floatDotStart"),
		token.ColonToken{},
		token.NewNumberToken(".1234"),
		token.CommaToken{},
		// "exponentialSmall": 1e10
		token.NewStringToken("exponentialSmall"),
		token.ColonToken{},
		token.NewNumberToken("1e10"),
		token.CommaToken{},
		// "exponentialLarge": 1E10
		token.NewStringToken("exponentialLarge"),
		token.ColonToken{},
		token.NewNumberToken("1E10"),
		token.CommaToken{},
		// "exponentialPlus": 1e+10
		token.NewStringToken("exponentialPlus"),
		token.ColonToken{},
		token.NewNumberToken("1e+10"),
		token.CommaToken{},
		// "exponentialMinus": 1e-10
		token.NewStringToken("exponentialMinus"),
		token.ColonToken{},
		token.NewNumberToken("1e-10"),
		token.CommaToken{},
		// "plus": +10
		token.NewStringToken("plus"),
		token.ColonToken{},
		token.NewNumberToken("+10"),
		token.CommaToken{},
		// "minus": -10
		token.NewStringToken("minus"),
		token.ColonToken{},
		token.NewNumberToken("-10"),
		token.RightBraceToken{},
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(token.StringToken{}, token.NumberToken{})); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestFailedLexer(t *testing.T) {
	f, err := os.Open("./testdata/lexer_structure_fragile.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	sut := NewLexer(string(b))
	got, err := sut.Execute()
	if got != nil {
		t.Errorf("want error %v, but got result %v", ErrLexer, got)
	}
	if !errors.Is(err, ErrLexer) {
		t.Fatalf("want ErrLexer, but got %v", err)
	}
}
