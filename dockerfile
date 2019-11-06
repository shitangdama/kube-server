FROM golang:latest AS builder

RUN mkdir app
RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.io

WORKDIR /app/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir app
WORKDIR /app/

RUN mkdir app

COPY --from=builder /app/app .

CMD ["./app"]  