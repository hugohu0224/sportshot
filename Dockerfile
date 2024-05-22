# first stage
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# allow for independent implementation
ARG SERVER
RUN if [ "$SERVER" = "crawler" ]; then \
      go build -o /app/crawler ./cmd/crawler; \
    elif [ "$SERVER" = "eventserver" ]; then \
      go build -o /app/eventserver ./cmd/grpcserver/event; \
    elif [ "$SERVER" = "webserver" ]; then \
      go build -o /app/webserver ./cmd/webserver; \
    else \
      echo "Unknown SERVER: $SERVER" && exit 1; \
    fi

# second stage
FROM golang:1.21

WORKDIR /app

ARG SERVER
COPY --from=builder /app/${SERVER} .
COPY --from=builder /app/scripts/entrypoint.sh .

COPY scripts/entrypoint.sh .
RUN chmod +x entrypoint.sh

ENV SERVER=${SERVER}

ENTRYPOINT ["./entrypoint.sh"]
