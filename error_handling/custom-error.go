package main

import "fmt"

// カスタムエラーの定義方法

// 以下のようにエラーに持たせたいフィールドを設定して構造体を定義する
type HTTPError struct {
	StatusCode int
	URL        string
}

// HTTPError構造体にメソッドを実装する
func (e *HTTPError) Error() string {
	return fmt.Sprintf("status code: %d, url: %s", e.StatusCode, e.URL)
}
