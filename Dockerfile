FROM golang:1.17

WORKDIR /home/market-tracker

COPY . .

RUN go get -u ./...

RUN go build -o market-tracker .

CMD ["./market-tracker"]
