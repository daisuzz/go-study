package main

import (
	"flag"
	"fmt"
)

func main() {
	// コマンドライン引数の定義
	namePtr := flag.String("name", "World", "a name")

	// コマンドライン引数のパース
	flag.Parse()

	// メッセージの生成
	message := fmt.Sprintf("Hello, %s!", *namePtr)

	// メッセージの出力
	fmt.Println(message)
}
