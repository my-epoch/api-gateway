FROM golang:latest

RUN mkdir "build"
ADD . /build
WORKDIR /build

RUN go build -o "build/server" cmd/server/main.go
EXPOSE 8080
ENTRYPOINT ["build/server"]

