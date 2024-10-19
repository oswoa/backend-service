## 動作確認

### 前提条件

1. [evans](https://github.com/ktr0731/evans?tab=readme-ov-file#installation)がインストールされていること

### 手順

1. `docker compose up -d`でコンテナを起動
1. `$ git submodule update --init`
1. `src/grpc/spec/README.md` を参照し、proto ファイルをコンパイルする
1. 生成ファイルを `src/grpc/proto` に配置する
1. ターミナルを起動し、`src/grpc/proto`に移動
1. evans を起動する
1. 確認対象の API を叩く
