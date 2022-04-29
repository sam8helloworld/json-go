package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sam8helloworld/json-go/lexer"
)

func main() {
	fmt.Println("ファイル読み取り処理を開始します")
	// ファイルをOpenする
	f, err := os.Open("testdata/string_only.json")
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
	lexer.Execute()
	// 出力
	fmt.Println(string(b))
}
