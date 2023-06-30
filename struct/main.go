package main

import (
	"fmt"
)

// struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// 構造体のインスタンス初期化
	var p1 Person
	// もしくは...
	p2 := &Person{
		Name: "Taro",
		Age:  29,
	}
	fmt.Println(p1, p2)
	/*
		以下のようにインスタンス初期化もできるけど、やらない。
		理由は、struct側の定義が変わった際にエラーが発生するが、修正が必要であることに気づきにくいため。
		p2 := &Person{
			"Taro",
			29,
		}
	*/

	// 上記のようにインスタンスを作成することもできるが、普通はファクトリー関数を用意する。
	// メリットがいっぱいある。
	// ゼロ値以外の初期値を与えられたり、よく使うユースケースが複数ある場合にそれぞれのケースをファクトリー関数に与えられる
	// 入力値のバリデーションができるなど。
	p3 := NewPerson("Jiro", 30)
	fmt.Println(p3)

	t := &TestType{
		v: "a",
	}
	// 値型のレシーバー関数内でレシーバーのフィールド変数を直接変更しても、値は変わらない
	t.TestFunc("b")
	fmt.Println(t) // 値は変わらない

	t2 := &TestType{
		v: "a",
	}
	// ポインタ型であれば値が変わる
	t2.TestPointerFunc("b")
	fmt.Println(t2)

	r.Refrect()
}

// ファクトリー関数
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: "Taro",
		Age:  29,
	}
}

type TestType struct {
	v string
}

// レシーバー関数
// 値型のレシーバーには値を追加できない
// (コンパイル時エラーは出ないが、下記のように呼び出すと値が変わらない)
func (t TestType) TestFunc(v string) {
	t.v = v
}

// レシーバー関数
// ポインタ型レシーバー
func (t *TestType) TestPointerFunc(v string) {
	t.v = v
}
