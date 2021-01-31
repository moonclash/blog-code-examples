FROM golang:latest

WORKDIR /gopi

RUN apt update -y && apt install -y sqlite3
# RUN go get -u github.com/mattn/go-sqlite3 github.com/gin-gonic/gin

COPY . .