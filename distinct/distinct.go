package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// 引数が2つ指定されているか確認
	if len(os.Args) != 3 {
		fmt.Println("usage: distinct <filename> <column>")
		os.Exit(1)
	}

	// ファイル名とカラム名を取得
	filename := os.Args[1]
	column := os.Args[2]

	// ファイルを開く
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(f)

	// ヘッダー行を読み込む
	header, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}

	// カラム名に対応する列番号を取得
	var columnIndex int
	for i, col := range header {
		if col == column {
			columnIndex = i
			break
		}
	}
	if columnIndex == 0 {
		fmt.Printf("column %s not found\n", column)
		os.Exit(1)
	}

	// 重複をチェックするための map を作成
	seen := make(map[string]bool)

	// ヘッダー行を出力
	fmt.Println(strings.Join(header, ","))

	// 行を1行ずつ読み込んで重複をチェックしながら出力
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// 既に現れたことがある場合はスキップ
		if seen[record[columnIndex]] {
			continue
		}
		seen[record[columnIndex]] = true

		// 重複していない場合は出力
		fmt.Println(strings.Join(record, ","))
	}
}
