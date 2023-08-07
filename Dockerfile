FROM golang:1.20

WORKDIR /go-bankmate

COPY . .

RUN go mod tidy

RUN go build -o app

CMD [ "./app" ]