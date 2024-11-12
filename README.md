# go_type_inference

## 実行方法

この

### **1. Dockerを起動する**

Dockerを起動してください。コマンドラインからでも、[`Docker Desktop`](https://www.docker.com/ja-jp/products/docker-desktop/)を起動する方法でも、何でも構いません。

### **2. コマンドを実行する**

[`Makefile`](./Makefile)に記載してある２つのコマンドを順番に実行します。

このコマンドで、[`Dockerfile`](./Dockerfile)から、イメージを作成します。
```md
make dcb
```

このコマンドで、上のコマンドで作成したイメージからコンテナを起動し、プログラムを実行します。
```md
make dcr
```
