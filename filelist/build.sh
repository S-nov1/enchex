#!/bin/bash 

docker run --rm -v "$PWD":/usr/src/encapp -w /usr/src/encapp golang:alpine go build -v
docker build -t enclist .