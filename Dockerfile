FROM golang:1.24

RUN net user /add nonroot

USER nonroot

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY database/ ./database/
COPY controllers/ ./controllers/
COPY models/ ./models/

ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN apt-get update \
&& apt-get --no-install-recommends install -y gcc libsqlite3-dev sqlite3 \
&& apt-get clean \
&& go mod tidy \
&& go build -o server .

EXPOSE 8080
CMD ["./server"]
