# syntax=docker/dockerfile:1


FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /es-task-prom-log-exporter

EXPOSE 2112

CMD [ "/es-task-prom-log-exporter" ]
