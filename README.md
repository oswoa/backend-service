## 動作確認

### 前提条件

1. 下記がインストールされていること
    1. [evans](https://github.com/ktr0731/evans?tab=readme-ov-file#installation)
    1. protoc-gen-go-grpc: `$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
    1. protoc-gen-go: `$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

### 手順

1. サブモジュールを更新: `$ git submodule update --init`
1. `src/grpc/spec/README.md` を参照し、proto ファイルをコンパイルする
    1. 生成ファイルを `src/grpc/proto` に配置する
1. コンテナを起動: `docker compose up -d`
1. 確認対象の API を叩く
    1. ターミナルを起動し、`src/grpc/proto`に移動
    1. evans を起動する: `$ evans --port 50001 --proto backend.proto`
    1. `$ call <API名>`
