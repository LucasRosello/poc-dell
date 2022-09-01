FROM golang:1.19

WORKDIR /
COPY . .
RUN go mod download

CMD ["go","run","main.go"]