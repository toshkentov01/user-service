FROM golang:1.17.5-alpine3.14

RUN mkdir /alif-tech-task

ADD . /alif-tech-task

WORKDIR /alif-tech-task

RUN go build -o main ./cmd/main.go
RUN go mod tidy
RUN go mod vendor

CMD ["./main"]
