FROM golang:alpine AS build

WORKDIR /app-build

COPY go.mod .
COPY go.sum .
# COPY .env .

COPY main.go .

RUN go build -o app .

FROM alpine

WORKDIR /app

COPY --from=build /app-build/app .

CMD ["./app"]