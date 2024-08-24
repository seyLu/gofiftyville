FROM golang:1.23-alpine
WORKDIR /gofiftyville

RUN apk add --no-cache gcc musl-dev

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/server ./cmd/server

CMD ["/gofiftyville/bin/server"]
EXPOSE 8080
