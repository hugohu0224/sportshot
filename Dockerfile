# first stage
FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# allow for independent build
ARG SERVER
RUN if [ "$SERVER" = "crawler" ]; then \
      go build -o /app/crawler ./cmd/crawler; \
    elif [ "$SERVER" = "eventserver" ]; then \
      go build -o /app/eventserver ./cmd/grpcserver/events; \
    elif [ "$SERVER" = "webserver" ]; then \
      go build -o /app/webserver ./cmd/webserver; \
    else \
      echo "Unknown SERVER: $SERVER" && exit 1; \
    fi

# second stage
FROM golang:1.21-alpine

WORKDIR /app

ARG SERVER
# get the compiled application
COPY --from=builder /app/${SERVER} .

# get static for webserver needed
COPY internal/webserver/static/ internal/webserver/static/
COPY internal/webserver/templates/ internal/webserver/templates/

# get Jason data for eventserver initial needed
COPY pkg/files/sportevents.basketball.json pkg/files/sportevents.basketball.json

# allow for independent run
COPY scripts/entrypoint.sh .
RUN chmod +x entrypoint.sh

ENV SERVER=${SERVER}

ENTRYPOINT ["./entrypoint.sh"]