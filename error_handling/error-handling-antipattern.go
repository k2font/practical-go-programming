package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/songmu/retry"
)

func getInvitedUserWithEmail(ctx, email string) (string, error) {
	return "", errors.New("not implemented")
}

// パターン1: 呼び出し元にエラーを返却するパターン
func ErrHndlAntiPattern() (string, error) {
	ctx := "context"
	email := "email"
	user, err := getInvitedUserWithEmail(ctx, email)
	return user, err // このようにerrだけ返却すると、エラーの詳細が分からない
}

func ErrHndl() (string, error) {
	ctx := "context"
	email := "email"
	user, err := getInvitedUserWithEmail(ctx, email)
	return user, fmt.Errorf("fail to get invited user w/email: %w", err)
	// このようにエラーの詳細を返却すると、エラーの詳細が分かるようになる
}

// パターン2: エラーをcatchするが処理は継続するパターン
func fetchCapacity() (int, error) {
	err := errors.New("not implemented") // エラーの発生！
	if err != nil {
		if errors.Is(err, http.ErrNotSupported) { // 特定のエラーの場合はerrを返却しない
			log.Printf("fetch capacity not found!", err) // ログを吐いて処理を継続する
			return -1, nil
		}
		return 0, err // それ以外ならエラーを返却する
	}
	return 0, nil
}

// パターン3: リトライを実施するパターン
// 例えばネットワーク越しにAPIをキックする場合、ネットワークの問題で処理が正常に終了しない可能性がある
// そのため、一定回数リトライして復帰を試みることで堅牢なアプリケーションになる
func Retry() {
	err := retry.Retry(2, 0, func() error {
		_, err := http.Get("https://example.com")
		return err
	})
	if err != nil {
		// エラーハンドリングなど
	}
}

// パターン4: リソースをクローズするパターン
// 省略
