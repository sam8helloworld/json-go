package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sam8helloworld/json-go/lexer"
	"github.com/sam8helloworld/json-go/parser"
	"github.com/sam8helloworld/json-go/printer"
)

func main() {
	fmt.Println("ファイル読み取り処理を開始します")
	// ファイルをOpenする
	f, err := os.Open("sample/sample.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	lexer := lexer.NewLexer(string(b))
	tokens, err := lexer.Execute()
	if err != nil {
		fmt.Println("error")
	}
	parser := parser.NewParser(*tokens)
	json, err := parser.Execute()
	if err != nil {
		fmt.Println("error")
	}
	p := printer.NewPrinter(json)
	p.Execute()
}
