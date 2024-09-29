FROM golang:alpine as build

WORKDIR /app

COPY . /app

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
