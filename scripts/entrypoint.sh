#!/bin/sh

case "$SERVER" in
  crawler)
    echo "Starting crawler..."
    exec ./crawler
    ;;
  eventserver)
    echo "Starting eventserver..."
    exec ./eventserver
    ;;
  webserver)
    echo "Starting webserver..."
    exec ./webserver
    ;;
  *)
    echo "Unknown service: $SERVER"
    exit 1
    ;;
esac
