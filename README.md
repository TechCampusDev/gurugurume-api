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
docker-compose exec app go run .
```
下記のサイトで{"message":"hello world"}と表示されたらDockerとGo(Gin)の環境構築完了。

http://localhost:8080/

## 5.PostgreSQLのインストール
ローカル環境へPostgreSQLをインストールする。
[ここからインストール](https://www.postgresql.jp/download)  
```
基本的に古すぎなければどのバージョンでも問題ないと思われる。version14では動くことを確認済み。
```  
ユーザ名やパスワードなどを設定すると思うので、保存しておく。  
.local.envファイルの内容を.envファイルを作成しそこにコピーする。それぞれ、DB_USER = DB_PASS =に設定したユーザ名とパスワードを追記する。 

## 6.データベースの作成
コマンドプロンプトを開き以下のコマンドでpostgresにアクセスする。  
```
psql -U {自分の設定したユーザ名}
```  
パスワードを求められるので入力する。この時、入力内容は画面に表示されないが入力は出来ている。  
次に以下のコマンドでデータベースを作成する。  
```
create database gurugurume;
```

## 7. .envファイルの設定
1. `SECRET_KEY`に適当な値を設定する

## 8.動作確認
```
docker-compose exec app go run .
```
下記のサイトで{"message":success}と表示されたらPostgreSQLの環境構築完了。

http://localhost:8080/db

## 環境のセットアップ
以下にアクセスする。  
http://localhost:8080/setenv  
`success setting environment`と表示されたら完了。