package main

import (
	"fmt"
	"sync"
)

// 巨大な構造体
type BigStruct struct {
	Member string
}

// Poolは初期化関数をNewフィールドに設定して作成
var pool = &sync.Pool{
	New: func() interface{} {
		return &BigStruct{}
	},
}

func Pool() {
	// インスタンスはGetメソッドで取得する
	// なければNew()を呼び出す
	b := NewBigStruct()

	fmt.Println(b)

	// 使い終わったらputでpoolする
	pool.Put(b)

	// sync.Poolでは以下を扱うことができない
	// チャネル
	// os.FileなどClose()で閉じる必要のあるもの
}

// ファクトリー関数を用意する手もある
func NewBigStruct() *BigStruct {
	b := pool.Get().(*BigStruct)
	return b
}
