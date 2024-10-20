package util

import (
	"fmt"
	"os"
)

// 指定された環境変数を取得する
func GetEnv(name string) (string, error) {

	if len(name) < 1 {
		return "", fmt.Errorf("error: 環境変数(%s)の取得に失敗しました", name)
	}

	env, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("error: 環境変数(%s)が設定されていません", name)
	} else {
		fmt.Printf("環境変数(%s): %s\n", name, env)
	}

	return env, nil
}
