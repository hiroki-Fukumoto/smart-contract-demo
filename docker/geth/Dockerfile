FROM ubuntu:latest

RUN apt-get update && apt install -y software-properties-common
RUN add-apt-repository -y ppa:ethereum/ethereum
RUN apt-get update && apt-get install -y \
  ethereum

WORKDIR /go/contracts
COPY ./contracts .
COPY ./src /go/src

RUN apt-get -y install nodejs npm
RUN npm i -g solc
RUN npm i -g ganache
