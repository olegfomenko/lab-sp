FROM golang:1.21-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/olegfomenko/lab-sp
COPY . .

ENV GOOS="linux"
ENV GOARCH="amd64"
RUN GOOS=linux go build  -o /usr/local/bin/lab-sp /go/src/github.com/olegfomenko/lab-sp

FROM ubuntu:latest

RUN apt-get update && apt-get install -y linux-tools-generic && apt-get install -y build-essential && apt-get install -y musl-dev
#RUN git clone https://github.com/brendangregg/FlameGraph

COPY --from=buildbase /usr/local/bin/lab-sp /usr/local/bin/lab-sp

WORKDIR /usr/lib/linux-tools/5.15.0-91-generic




