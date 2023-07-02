package main

import "fmt"

func main() {
	// ジェネリクスを使わない場合
	fmt.Println(ReturnString("Hello World"))
	fmt.Println(ReturnBoolean(true))
	fmt.Println(ReturnInteger(1))

	// ジェネリクスを使う場合
	fmt.Println(ReturnValue("Hello World"))
	fmt.Println(ReturnValue(true))
	fmt.Println(ReturnValue(1))
}
