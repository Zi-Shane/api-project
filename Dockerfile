FROM docker.io/library/golang:alpine as builder
# RUN mkdir -p /app
WORKDIR /app

COPY . .
RUN go mod download &&\
    go build -o app


# FROM docker.io/library/golang
FROM docker.io/library/alpine
WORKDIR /app
COPY --from=builder /app .

EXPOSE 9090

ENTRYPOINT ["./app"]
