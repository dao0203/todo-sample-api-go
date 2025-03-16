# ベースイメージ
FROM golang:1.23

# 作業ディレクトリを設定
WORKDIR /app

# ファイルをコピー
COPY . .

# 依存関係を取得
RUN go mod tidy

# ポートを指定
EXPOSE 8080

# アプリケーション実行
CMD ["go", "run", "cmd/app/main.go"]
