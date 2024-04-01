FROM golang:1.20 as builder

WORKDIR /app
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -o hello

FROM scratch
COPY --from=builder /app/hello /
ENTRYPOINT ["/hello"]
