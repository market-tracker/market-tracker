FROM golang:1.17

WORKDIR /home/market-tracker

RUN go env -w GO111MODULE=on

CMD ["tail", "-f", "/dev/null"]
