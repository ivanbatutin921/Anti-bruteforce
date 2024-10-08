FROM golang:alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /anti-bruteforce

EXPOSE 3000

CMD [ "/anti-bruteforce" ]