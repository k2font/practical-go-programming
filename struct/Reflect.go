package main

import (
	"fmt"
	"reflect"
)

type MapStruct struct {
	Str     string  `map:"str"`
	StrPtr  *string `map:"str"`
	Bool    bool    `map:"bool"`
	BoolPtr *bool   `map:"bool"`
	Int     int     `map:"int"`
	IntPtr  *int    `map:"int"`
}

func Reflect() {
	src := map[string]string{
		"str":  "string data",
		"bool": "true",
		"int":  "1",
	}

	var ms MapStruct
	Decode(&ms, src) // srcのmap[string]stringをデコードしてMapStruct型のmsに格納する
	fmt.Printf("%+x\n", ms)
}

func Decode(target interface{}, src map[string]string) error {
	v := reflect.ValueOf(target)
	e := v.Elem()
	decode(e, src)
}

// structの型が不明なので、reflectを使ってデコードする
// これぞメタプログラミングみたいなコードになりそう
func decode(e reflect.Value, src map[string]string) error {
	t := e.Type()
	// MapStruct構造体の型を1つずつ取得して処理する
	// Str, StrPtr, Bool, BoolPtr, Int, IntPtrの順に処理
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 内部に埋め込まれた構造体は再帰処理する
		if f.Anonymous {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}
	}
	return nil
}
