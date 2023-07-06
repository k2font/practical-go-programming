package main

import "fmt"

var version string

// go buildのオプションを学ぶ
func main() {
	fmt.Printf("Hello world! version=%s\n", version)
}

// 本番環境向けビルド
// $ CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X main.version=1.0.0" main.go
// -ldflags "-s -w" はビルド時にシンボルテーブルとデバッグ情報を削除するオプション
// ビルドしたバイナリサイズが小さくなるので積極的につける
// -trimpathはビルド時に絶対パスを削除するオプション
// -X はビルド時の変数を初期化できるオプション
