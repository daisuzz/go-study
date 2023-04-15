package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// -pattern オプションの定義
	pattern := flag.String("pattern", "", "a regular expression")

	// フラグのパース
	flag.Parse()

	// 正規表現のコンパイル
	re, err := regexp.Compile(*pattern)
	if err != nil {
		fmt.Println("invalid regular expression")
		return
	}

	// コマンドライン引数からファイル名を取得
	filename := flag.Arg(0)

	// ファイルのオープン
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file:", filename)
		return
	}
	defer file.Close()

	// ファイルの内容を1行ずつ読み込み、正規表現にマッチした行を出力
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			fmt.Println(line)
		}
	}

	// エラー処理
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
