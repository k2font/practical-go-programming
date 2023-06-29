package main

import "fmt"

// structの埋め込み
type Book struct {
	Title string
	ISBN  string
}

func (b Book) GetAmazonURL() string {
	return "https://www.amazon.co.jp/dp/" + b.ISBN
}

type OreillyBook struct {
	Book
	ISBN13 string
}

func (o OreillyBook) GetAmazonURL() string {
	return "https://www.amazon.co.jp/dp/" + o.ISBN13
}

func getBook() {
	// インスタンス作成時、埋め込んだstructの名前と同名のフィールドを指定する必要がある(この場合はBook)
	ob := OreillyBook{
		ISBN13: "978-4873119048",
		Book: Book{
			Title: "Go言語によるWebアプリケーション開発",
		},
	}

	fmt.Println(ob.GetAmazonURL())
}
