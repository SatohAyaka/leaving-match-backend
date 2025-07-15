FROM golang:1.23.3


# 作業ディレクトリの作成
WORKDIR /app

# go.mod と go.sum をコピー（依存関係を先にインストールするため）
COPY app/go.mod app/go.sum ./
RUN go mod download

# ソースコードコピー
COPY app/ .
COPY .env .

# ポートの指定（必要であれば）
EXPOSE 8085

# 起動コマンド
CMD ["go", "run", "main.go"]
