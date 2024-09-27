#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./cmd/server/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY deploy/server/quotes.txt /quotes.txt
ENTRYPOINT /app -diff 5 -quotes /quotes.txt
LABEL Name=tcppow-server Version=0.0.1
EXPOSE 8080
