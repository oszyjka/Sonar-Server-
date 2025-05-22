FROM golang:1.24

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN apt-get update \
&& apt-get install -y gcc libsqlite3-dev sqlite3 \
&& apt-get clean \
&& go mod tidy \
&& go build -o server .

EXPOSE 8080
CMD ["./server"]
