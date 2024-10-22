package util

import (
	"fmt"
	"log"
	"os"
)

// 指定された環境変数を取得する
func GetEnv(name string) (string, error) {

	if len(name) < 1 {
		return "", fmt.Errorf("error: 環境変数(%s)が設定されていません", name)
	}

	env, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("error: 環境変数(%s)の取得に失敗しました", name)
	}
	log.Printf("環境変数(%s): %s\n", name, env)

	return env, nil
}
