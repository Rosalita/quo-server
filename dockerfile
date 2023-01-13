FROM golang:1.19
LABEL maintainer="rosie@rosie.dev"

RUN mkdir /quote-service

# Copy go.mod and go.sum files to container
COPY go.* /quote-service/ 

WORKDIR /quote-service

# Fetch dependencies
RUN go mod download

# Copy go files to container
COPY *.go /quote-service/

RUN go build -o app 
CMD ["./app"]