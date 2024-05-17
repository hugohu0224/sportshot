# first stage
FROM golang:1.21 as builder

WORKDIR /app

ARG SERVER=crawler

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# allow for independent implementation
RUN if [ "SERVER" = "crawler" ]; then \
      go build -o /app/crawler ./cmd/crawler; \
    elif [ "SERVER" = "eventserver" ]; then \
      go build -o /app/grpcserver ./cmd/grpcserver/event; \
    elif [ "SERVER" = "webserver" ]; then \
      go build -o /app/webserver ./cmd/webserver; \
    else \
      echo "Unknown service: SERVER" && exit 1; \
    fi

# second stage
FROM golang:1.21

WORKDIR /app

COPY --from=builder /app/${SERVER} .

COPY scripts/entrypoint.sh .
RUN chmod +x entrypoint.sh

ENV SERVER=${SERVER}

ENTRYPOINT ["./entrypoint.sh"]
