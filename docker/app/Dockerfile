FROM golang:1.24-alpine

WORKDIR /app

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY app/ .
COPY .env .

EXPOSE 8085

CMD ["go", "run", "main.go"]
