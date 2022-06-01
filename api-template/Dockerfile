FROM dwarvesf/sql-migrate-git as migrate

FROM golang:1.16 as builder
WORKDIR /go/src/github.com/msation/api
COPY . .
RUN go mod download -x && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server cmd/server/*.go && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o indexer cmd/indexer/*.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/github.com/msation/api/server .
COPY --from=builder /go/src/github.com/msation/api/indexer .
COPY --from=migrate /bin/sql-migrate /bin
COPY dbconfig.yml /app
COPY migrations /app/migrations
CMD ["./server"]