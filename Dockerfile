FROM golang:1.20.1-alpine AS builder

WORKDIR /go/src/github.com/andyantrim/qstore
COPY . .
RUN go build -o /go/bin/qstore cmd/qstore/main.go

# Path: Dockerfile
FROM alpine AS runner
COPY --from=builder /go/bin/qstore /usr/local/bin/qstore
ENTRYPOINT ["qstore"]
CMD ["serve"]
