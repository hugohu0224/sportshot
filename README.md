# Introduction

Sportshot is a sports odds scanning tool that automatically crawls and retrieves sports betting odds. It's designed to be easy to use, enabling users to quickly access and track the historical data of the odds they're interested in. This tool helps you keep track of changes in sports odds efficiently and effectively.

## Todo List
### Crawler
* :white_check_mark: Implement basketball crawler.
* :white_check_mark: Persist the crawled data to MongoDB.

* :black_square_button: Remove spaces or symbols from the team names.

### WebServer
* :white_check_mark: Implement the get method to receive user parameters.
* :white_check_mark: Implement the accessing process the event server via the grpc client.
* :large_blue_diamond: Implement user registration function to make the service available to others(by local).

### WebUI
* :black_square_button: Implement the front-end page for odds searching.

### EventServer
* :white_check_mark: Implement the interaction between grpc server and MongoDB.
* :black_square_button: Receives parameters from the API to filter the data.
* :black_square_button: Implement the distributed crawler structure.

### DbServer
* :white_check_mark: Implement MongoDB.
* :black_square_button: Implement Redis.

### Others
* :black_square_button: Implement One-click activation of all services by docker-compose.
* :black_square_button: Implement the services splitability.
* :black_square_button: Enhance the maintainability of the config file. 

### Emoji meaning
* :black_square_button: Not finish yet.
* :white_check_mark: Finished.
* :large_blue_diamond: Might be added.