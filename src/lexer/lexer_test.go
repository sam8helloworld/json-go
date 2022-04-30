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
	f, err := os.Open("../testdata/string_only.json")
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
		{
			Type:       token.LeftBraceType,
			Expression: "{",
		},
		{
			Type:       token.StringType,
			Expression: "string",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.StringType,
			Expression: "hogehoge",
		},
		{
			Type:       token.RightBraceType,
			Expression: "}",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestFailedStringTokenize(t *testing.T) {
	f, err := os.Open("../testdata/string_only_fragile.json")
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
	f, err := os.Open("../testdata/bool_only.json")
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
		{
			Type:       token.LeftBraceType,
			Expression: "{",
		},
		{
			Type:       token.StringType,
			Expression: "boolTrue",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.TrueType,
			Expression: true,
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		{
			Type:       token.StringType,
			Expression: "boolFalse",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.FalseType,
			Expression: false,
		},
		{
			Type:       token.RightBraceType,
			Expression: "}",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestFailedBoolTokenize(t *testing.T) {
	f, err := os.Open("../testdata/bool_only_fragile.json")
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
	f, err := os.Open("../testdata/null_only.json")
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
		{
			Type:       token.LeftBraceType,
			Expression: "{",
		},
		{
			Type:       token.StringType,
			Expression: "null",
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
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestFailedNullTokenize(t *testing.T) {
	f, err := os.Open("../testdata/null_only_fragile.json")
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
	f, err := os.Open("../testdata/number_only.json")
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
		{
			Type:       token.LeftBraceType,
			Expression: "{",
		},
		// "int": 100
		{
			Type:       token.StringType,
			Expression: "int",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(100),
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		// "float": 1.234
		{
			Type:       token.StringType,
			Expression: "float",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(1.234),
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		// "floatDotStart": .1234
		{
			Type:       token.StringType,
			Expression: "floatDotStart",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(.1234),
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		// "exponentialSmall": 1e10
		{
			Type:       token.StringType,
			Expression: "exponentialSmall",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(1e10),
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		// "exponentialLarge": 1E10
		{
			Type:       token.StringType,
			Expression: "exponentialLarge",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(1e10),
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		// "exponentialPlus": 1e+10
		{
			Type:       token.StringType,
			Expression: "exponentialPlus",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(1e+10),
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		// "exponentialMinus": 1e-10
		{
			Type:       token.StringType,
			Expression: "exponentialMinus",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(1e-10),
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		// "plus": +10
		{
			Type:       token.StringType,
			Expression: "plus",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(10),
		},
		{
			Type:       token.CommaType,
			Expression: ",",
		},
		// "minus": -10
		{
			Type:       token.StringType,
			Expression: "minus",
		},
		{
			Type:       token.ColonType,
			Expression: ":",
		},
		{
			Type:       token.NumberType,
			Expression: float64(-10),
		},
		{
			Type:       token.RightBraceType,
			Expression: "}",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}

func TestFailedNumberTokenize(t *testing.T) {
	f, err := os.Open("../testdata/number_only_fragile.json")
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
		t.Errorf("want error %v, but got result %v", ErrNumberTokenize, got)
	}
	if !errors.Is(err, ErrNumberTokenize) {
		t.Fatalf("want ErrNumberTokenize, but got %v", err)
	}
}

func TestFailedLexer(t *testing.T) {
	f, err := os.Open("../testdata/lexer_structure_fragile.json")
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
