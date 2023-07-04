package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// errorsパッケージを用いてerrorを生成する
var EOF = errors.New("EOF")

func main() {
	validate(0)

	res, _ := ReadContents("https://example.com")
	fmt.Println(string(res))
}

func validate(length int) error {
	if length <= 0 {
		// fmt.Errorf()を用いてフォーマットされたエラーを表示できる
		return fmt.Errorf("length must be greater than 0, but length = %d", length)
	}
	return nil
}

// カスタムエラーをレスポンスする場合は以下のようにする
func ReadContents(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// ステータスコードが200以外の場合はエラーを返す
	if res.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: res.StatusCode,
			URL:        url,
		}
	}

	return io.ReadAll(res.Body)
}
