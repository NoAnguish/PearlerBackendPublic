FROM golang:1.18.1-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./backend ./backend

RUN go build -o ./pearler ./backend/main/main.go

EXPOSE 3422

CMD [ "./pearler" ]