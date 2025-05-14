package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// transform: 元の出力 を受け取って、新しい出力を返す関数
func transform(y float64) float64 {
	return y * 5 / 1024
}

func main() {
	// 引数でファイル名を受け取る
	// 指定がなければ標準入力を使う
	var r io.Reader = os.Stdin
	if len(os.Args) == 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("ファイルオープン失敗: %v", err)
		}
		defer f.Close()
		r = f
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			// 空行やコメント行はスキップ
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			log.Printf("無視: フォーマット不正な行: %q\n", line)
			continue
		}

		// 第一カラムはそのまま文字列として出力
		in0 := parts[0]

		// 第二カラムを float64 にパース
		y0, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Printf("警告: 数値変換失敗 (%q): %v\n", parts[1], err)
			continue
		}

		// 変換
		y1 := transform(y0)

		// 出力（小数点の揃え に注意）
		fmt.Printf("%s,%.1f\n", in0, y1)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("読み込みエラー: %v", err)
	}
}
