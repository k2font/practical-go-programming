package main

import (
	"log"
	"reflect"
)

type MapStruct struct {
	Str    string  `map:"str"`
	StrPtr *string `map:"str"`
}

func Reflect() {
	src := map[string]string{
		"str": "string data",
	}

	var ms MapStruct
	Decode(&ms, src)
	log.Println(ms)
}

func Decode(target interface{}, src map[string]string) {
	v := reflect.ValueOf(target)
	e := v.Elem()
	decode(e, src)
}

func decode(e reflect.Value, src map[string]string) {

}
