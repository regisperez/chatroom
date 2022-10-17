FROM golang:1.18

ENV GO111MODULE=on

COPY ./ /go/src/chatroom

WORKDIR /go/src/chatroom

COPY go.mod ./
COPY go.sum ./

RUN go mod download

WORKDIR /go/src/chatroom/main

RUN go build -o /chatroom

CMD [ "/chatroom" ]
