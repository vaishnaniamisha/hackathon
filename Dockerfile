FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o main .

EXPOSE 8070

CMD ["./main"]
