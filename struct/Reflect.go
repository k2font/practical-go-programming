package main

import (
	"fmt"
	"log"
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
	log.Println(ms)
}

func Decode(target interface{}, src map[string]string) {
	v := reflect.ValueOf(target)
	e := v.Elem()
	decode(e, src)
}

// structの型が不明なので、reflectを使ってデコードする
// これぞメタプログラミングみたいなコードになりそう
func decode(e reflect.Value, src map[string]string) {
	t := e.Type()
	fmt.Println(t)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f)
	}
}
