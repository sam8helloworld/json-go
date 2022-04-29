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
			Type:    token.LeftBraceType,
			Literal: "{",
		},
		{
			Type:    token.StringType,
			Literal: "string",
		},
		{
			Type:    token.ColonType,
			Literal: ":",
		},
		{
			Type:    token.StringType,
			Literal: "hogehoge",
		},
		{
			Type:    token.RightBraceType,
			Literal: "}",
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
			Type:    token.LeftBraceType,
			Literal: "{",
		},
		{
			Type:    token.StringType,
			Literal: "boolTrue",
		},
		{
			Type:    token.ColonType,
			Literal: ":",
		},
		{
			Type:    token.TrueType,
			Literal: "true",
		},
		{
			Type:    token.CommaType,
			Literal: ",",
		},
		{
			Type:    token.StringType,
			Literal: "boolFalse",
		},
		{
			Type:    token.ColonType,
			Literal: ":",
		},
		{
			Type:    token.FalseType,
			Literal: "false",
		},
		{
			Type:    token.RightBraceType,
			Literal: "}",
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
			Type:    token.LeftBraceType,
			Literal: "{",
		},
		{
			Type:    token.StringType,
			Literal: "null",
		},
		{
			Type:    token.ColonType,
			Literal: ":",
		},
		{
			Type:    token.NullType,
			Literal: "null",
		},
		{
			Type:    token.RightBraceType,
			Literal: "}",
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
