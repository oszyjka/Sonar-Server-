FROM golang:1.24

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

RUN go mod tidy
RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
