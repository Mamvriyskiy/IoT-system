FROM golang:latest

WORKDIR /app

COPY ./main ./
COPY ./configs/config.yml ./configs/

RUN go mod download

CMD ["go", "run", "./cmd/main.go"]
