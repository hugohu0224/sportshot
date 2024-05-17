#!/bin/sh

case "$SERVER" in
  crawler)
    echo "Starting crawler service..."
    exec ./crawler
    ;;
  eventserver)
    echo "Starting gRPC server service..."
    exec ./eventserver
    ;;
  webserver)
    echo "Starting webserver service..."
    exec ./webserver
    ;;
  *)
    echo "Unknown service: $SERVER"
    exit 1
    ;;
esac
