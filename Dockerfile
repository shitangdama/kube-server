FROM golang:latest AS builder

WORKDIR /build
# RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.io

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /build/demo /app/

CMD ["./app"]  