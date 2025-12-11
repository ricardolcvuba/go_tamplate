FROM golang:1.25.0-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["go", "run", "src/cmd/api/main.go"]
