# gurugurume-api

## 1.Docker Desktopのインストール

Docker Desktopをインストールしてない場合は先にインストール

[macの方はこちら](https://matsuand.github.io/docs.docker.jp.onthefly/desktop/mac/install/)
[windowsの方はこちら](https://matsuand.github.io/docs.docker.jp.onthefly/desktop/windows/install/)

## 2.イメージを作成する
```
docker-compose build
```
#### ※注意点
DockerDesktopを開いていないと下記のエラーが発生する.
```
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
```

## 3.コンテナを作成する
```
docker-compose up -d
```
#### ※注意点
同じポート番号のコンテナを作成している場合エラーが発生します。被っていないか確認してください.

ポート番号:8080

## 4.実行コマンド
```
docker-compose exec app go run main.go
```
下記のサイトで{"message":"hello world"}と表示されたら完了。

http://localhost:8080/
