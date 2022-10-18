# Chatroom

A simple browser-based chat application using Go.
This application allow several users to talk in a chatroom and also to get stock quotes from an API using a specific command.

## Features
- Allow registered users to log in and talk with other users in a chatroom.
- Allow users to post messages as commands into the chatroom with the following format /stock=stock_code
- Chat messages ordered by their timestamps and show only the last 50 messages.
- Decoupled bot that will call an API using the stock_code as a parameter (https://stooq.com/q/l/?s=aapl.us&f=sd2t2ohlcv&h&e=csv,here aapl.us is the stock_code)
- The bot parse the received CSV file and send a message back into the chatroom using the following format: “APPL.US quote is $93.42 per share”.
- Unit tests

## Installation Guide
- Golang
- Docker
- Docker-compose
## Building
After the installation, inside the folder execute the command

```bash
  docker-compose up
```
## Running
```
  http://localhost:8010/
```

## First Access
```
  user admin
  password admin
```

## Postman Collection
The API CRUD endpoints and inputs are described at https://www.getpostman.com/collections/36139809334d99a93c57
