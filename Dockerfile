# ビルド用ステージ
FROM golang:1.20.12-alpine3.19 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# 実行ステージ
FROM alpine:3.19

WORKDIR /root/

# ビルドステージから実行バイナリをコピー
COPY --from=builder /app/main .

# 8080ポートを露出
EXPOSE 8080

# アプリケーション実行
CMD ["./main"]
