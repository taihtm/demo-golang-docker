FROM golang:latest
LABEL maintainer="Hoang Tran Minh Tai <tai.htm@gmail.com>"
ADD . /app
WORKDIR /app
RUN rm -rf go.mod
RUN go mod init demo-golang-docker
RUN go build
ENTRYPOINT ./demo-golang-docker
