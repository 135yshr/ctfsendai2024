# ビルドステージ
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 依存関係のファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o main ./main.go

# 実行ステージ
#FROM scratch
FROM scratch

EXPOSE 8080

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/configs ./configs
COPY --from=builder /app/web ./web

ENTRYPOINT ["./main"]
