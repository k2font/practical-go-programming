package main

// ジェネリクスを使う場合
// https://go.dev/doc/tutorial/generics

// comparableは比較可能な型を表す
func ReturnValue[T comparable](x T) T {
	return x
}

// 特定の型のみを受け取ることもできる
func ReturnValueSandI[T string | int](x T) T {
	return x
}
