# Builder container
FROM golang:1.15 AS builder

# goの設定
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE on

# 作業ディレクトリの設定
WORKDIR /app

COPY . .
# モジュールのダウンロード・airの追加
RUN go mod download && \
    go get -u github.com/cosmtrek/air

# airの起動コマンド
CMD ["air"]