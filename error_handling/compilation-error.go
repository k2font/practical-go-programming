package main

import "io"

/* 複数のエラーハンドラをまとめる方法
// 例えば以下のようにファイルディスクリプタに複数回データを書き込む場合を考える
// _, err = file.Write(data1)
// if err != nil {
// 	return err
// }
// _, err = file.Write(data2)
// if err != nil {
// 	return err
// }
// _, err = file.Write(data3)
// if err != nil {
// 	return err
// }
*/

// まずはファイルディスクリプタに書き込む時に発生したエラーをまとめる構造体を定義する
type errWrite struct {
	w   io.Writer
	err error
}

// 上記構造体をラップするメソッドを用意する
func (e *errWrite) write(b []byte) {
	if e.err != nil {
		return
	}
	_, e.err = e.w.Write(b)
}

// あとは前述のエラーハンドラをまとめる
func SummarizingTest() {
	ew := &errWrite{w: fd}
	ew.write(data1)
	ew.write(data2)
	ew.write(data3)

	if ew.err != nil {
		return ew.err
	}
}
