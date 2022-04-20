FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /build/zero
ADD go.mod .
ADD go.sum .
RUN go mod download
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/main string/main.go


FROM alpine
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /app/main /app/main

CMD ["./main"]