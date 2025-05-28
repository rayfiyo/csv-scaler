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

const (
	// 有効数字を何桁にするか
	sig = 4
)

// transform: スケーリング
func transform(x, y float64) (float64, float64) {
	return (x * 5 / 4095), (y * 5 / 1023)
}

func main() {
	// 引数でファイル名を受け取る
	// 指定がなければ標準入力を使う
	var r io.Reader = os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("ファイルオープン失敗: %v", err)
		}
		defer f.Close()
		r = f
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		raw := scanner.Text()
		line := strings.TrimSpace(raw)
		if line == "" || strings.HasPrefix(line, "#") {
			// 空行やコメント行はそのままにしてスキップ
			fmt.Print(raw)
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			log.Printf("無視: フォーマット不正な行: %q\n", line)
			continue
		}

		// 0 列目を float64 にパース
		x0, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			log.Printf("警告: 0行目の数値変換失敗 (%q): %v\n", parts[0], err)
			continue
		}
		// 1 列目を float64 にパース
		y0, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Printf("警告: 1行目の数値変換失敗 (%q): %v\n", parts[1], err)
			continue
		}

		// 変換
		x1, y1 := transform(x0, y0)

		// 有効数字 sig 桁で出力
		fmt.Printf("%.*g,%.*g\n", sig, x1, sig, y1)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("読み込みエラー: %v", err)
	}
}
