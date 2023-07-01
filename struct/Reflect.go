package main

import (
	"log"
	"reflect"
	"strconv"
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
		"str":   "shoichiro",
		"bool":  "true",
		"int":   "1",
		"int64": "32478170570532", // 変換先の構造体に存在しないフィールドは無視されるようになる(無視しないような処理も書ける)
	}

	// タグを使った構造体へのデータの書き込み処理
	var ms MapStruct
	err := Decode(&ms, src) // srcのmap[string]stringをデコードしてMapStruct型のmsに格納する
	if err != nil {
		// TODO: エラー処理
	}
	log.Println(ms)
}

func Decode(target interface{}, src map[string]string) error {
	v := reflect.ValueOf(target)
	e := v.Elem()
	return decode(e, src)
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

		// 子要素が構造体であれば再帰処理
		if f.Type.Kind() == reflect.Struct {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// タグがなければフィールド名を使用する
		key := f.Tag.Get("map")
		if key == "" {
			key = f.Name
		}

		// 元データになければスキップ
		sv, ok := src[key]
		if !ok {
			continue
		}

		// フィールドの型を取得
		var k reflect.Kind
		var isP bool
		// 対象の型がポインターではない場合
		if f.Type.Kind() != reflect.Ptr {
			k = f.Type.Kind()
		} else { // 対象の型がポインターの場合
			k = f.Type.Elem().Kind() // 値を取得してkに格納する

			// もしkがポインターのポインターであれば無視する
			if k == reflect.Ptr {
				continue
			}

			// 対象の値はポインターかどうかを管理するフラグを立てる
			isP = true
		}

		// 型に応じて値をセットする
		switch k {
		case reflect.String:
			if isP {
				e.Field(i).Set(reflect.ValueOf(&sv))
			} else {
				e.Field(i).SetString(sv)
			}
		case reflect.Bool:
			b, err := strconv.ParseBool(sv)
			if err == nil {
				if isP {
					e.Field(i).Set(reflect.ValueOf(&b))
				} else {
					e.Field(i).SetBool(b)
				}
			}
		case reflect.Int:
			n64, err := strconv.ParseInt(sv, 10, 64)
			if err == nil {
				if isP {
					n := int(n64)
					e.Field(i).Set(reflect.ValueOf(&n))
				} else {
					e.Field(i).SetInt(n64)
				}
			}
		}
	}
	return nil
}
