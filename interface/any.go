package main

import "fmt"

func AnyTest() {
	// interface{}型は何でも入れられる
	var a interface{}
	// interface{}型はGo1.18からany型として使える
	var b any

	a = 1
	b = "Hello World"

	fmt.Println(a, b)

	// TypeScriptと同じでリリース前だけの利用に留めたい
	// ただし、どんなデータが落ちてくるか分からないAPIから、JSONデータを取得したいときはany型を使える
	// この場合、(多くの場合はそうだが)プロパティのキー名は文字列にする必要がある
	var ieyasu = map[string]any{
		"name": "Ieyasu",
		"age":  73,
	}

	fmt.Println(ieyasu["name"], ieyasu["age"])

	// any型のままでは使いづらいので、型アサーションを使って型を変換する
	// が、おなじくGo1.18で導入されたジェネリクスを用いるとその必要もなくなる。別コードで。
}
