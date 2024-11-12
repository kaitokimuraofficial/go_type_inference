# go_type_inference

## 実行方法

このプログラムを実行するには、事前に[`Golang`](https://go.dev/)の実行環境をセットアップする必要があります。
以下に簡単な実行手順を紹介するので、ぜひ試してみてください。([`Docker`](https://www.docker.com/ja-jp/)がインストールされていることが前提です)

### **1. Dockerを起動する**

Dockerを起動してください。コマンドラインからでも、[`Docker Desktop`](https://www.docker.com/ja-jp/products/docker-desktop/)を起動する方法でも、何でも構いません。

### **2. コマンドを実行する**

[`Makefile`](./Makefile)に記載されている２つのコマンドを順番に実行します。

以下のコマンドで、[`Dockerfile`](./Dockerfile)をもとにイメージを作成します。
```md
make dcb
```

次に、上記で作成したイメージからコンテナを起動し、プログラムを実行します。
```md
make dcr
```
