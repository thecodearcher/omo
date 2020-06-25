FROM golang:1.14

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /go/src/auth

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build ." --command=./auth