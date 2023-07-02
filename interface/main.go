package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
)

func main() {
	// Show()メソッドを持つインスタンスは何でも入れられる変数
	var warn Warning
	// インターフェースを満たす構造体を代入できる

	// コンソールに出力
	warn = &ConsoleWarning{}
	warn.Show("Hello World to Console")

	// デスクトップにアラートを出力
	warn = &DesktopWarning{}
	warn.Show("Hello World to Desktop")

	AnyTest()
}

// すべての構造体の共通インターフェース
type Warning interface {
	Show(message string)
}

// コンソールにメッセージを通知する構造体
type ConsoleWarning struct{}

// コンソールにメッセージを表示するShow関数
func (c ConsoleWarning) Show(message string) {
	fmt.Fprintf(os.Stderr, "[%s]: %s\n", os.Args[0], message)
}

// デスクトップにメッセージを通知する構造体
type DesktopWarning struct{}

// デスクトップにアラートを表示するShow関数
func (d DesktopWarning) Show(message string) {
	// beeepおもろ。あとでコード読む。
	beeep.Alert(os.Args[0], message, "")
}

// Slackにメッセージを通知する構造体
// TODO: 実装は省略
type SlackWarning struct{}
