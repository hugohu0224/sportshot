FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o crawler ./cmd/crawler

FROM alpine:latest as crawler
WORKDIR /root/
COPY --from=builder /app/crawler .
ENV SERVICE=crawler
ENV PORT=50051
EXPOSE 50051
CMD ["./crawler"]