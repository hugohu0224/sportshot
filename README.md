﻿# Introduction

Sportshot is a sports odds scanning tool that automatically crawls and retrieves sports betting odds. It's designed to
be easy to use, enabling users to quickly access and track the historical data of the odds they're interested in. This
tool helps you keep track of changes in sports odds efficiently and effectively.

## Architecture Reference

![Alt text](pkg/files/architecture.png)

## Todo List

### Crawler

* :white_check_mark: Implement basketball crawler.
* :white_check_mark: Persist the crawled data to MongoDB.
* :black_square_button: Remove spaces or symbols from the team names.

### WebServer

* :white_check_mark: Implement the get method to receive user parameters.
* :white_check_mark: Implement the accessing process the event server via the grpc client.
* :white_check_mark: Implement user registration function to make the service available to others(by local).
* :white_check_mark: Implement JWT to make sure that the user is allowed to access a specific page.

### WebUI

* :white_check_mark: Implement the front-end page for odds searching.
* :white_check_mark: Implement a button to control the on/off of the crawler.
* :large_blue_diamond: Made the front-end webpage to display the crawled odds in real time.

### EventServer(grpc)

* :white_check_mark: Implement the interaction between grpc server and MongoDB.
* :white_check_mark: Receives parameters from the API to filter the data.
* :white_check_mark: Implement gRPC naming and discovery using etcd.
* :white_check_mark: Implement etcd load balancer.
* :x: Implement the distributed crawler structure.

### DbServer

* :white_check_mark: MongoDB for web crawler.
* :white_check_mark: etcd for grpc.
* :white_check_mark: MySQL for user login authentication.
* :x: Redis.
* :x: Kafka.

### Others

* :white_check_mark: Implement One-click activation of all services by docker-compose.
* :white_check_mark: Implement the services splitability.
* :white_check_mark: Enhance the maintainability of the config file.

### Emoji meaning

* :black_square_button: Not finish yet.
* :white_check_mark: Finished.
* :large_blue_diamond: Might be added.
* :x: Cancel
