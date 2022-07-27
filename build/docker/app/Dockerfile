# syntax=docker/dockerfile:1

##
## Build
##
FROM debian:latest

WORKDIR /app

RUN apt update -y
RUN apt install curl wget -y

RUN wget -O /tmp/go1.18.4.linux-amd64.tar.gz https://go.dev/dl/go1.18.4.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.18.4.linux-amd64.tar.gz

ENV PATH="${PATH}:/usr/local/go/bin"

RUN curl -fsSL https://deb.nodesource.com/setup_17.x | bash -
RUN apt update -y
RUN apt install nodejs -y

RUN npm install -g yarn

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /marlaone-website

RUN cd ./web && yarn && yarn build

EXPOSE 1819

RUN useradd -u 8877 nonroot

USER nonroot

ENTRYPOINT ["/marlaone-website", "serve"]