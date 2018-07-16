FROM golang:1.9.5

WORKDIR /go/src/interpreter_monkey
COPY . .

RUN go build main.go

CMD ["./main"]