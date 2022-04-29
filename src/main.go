package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("ファイル読み取り処理を開始します")
	// ファイルをOpenする
	f, err := os.Open("testdata/testdata01.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
	}
	// 出力
	fmt.Println(string(b))
}
