#!/bin/sh

case "$SERVER" in
  webcrawler)
    echo "Starting crawler service..."
    exec ./webcrawler
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
