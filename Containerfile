FROM golang:1.22-alpine as build

WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY core core
COPY helpers helpers
COPY routes routes
RUN go build -o ./dist/sailing-seas .

FROM alpine

WORKDIR /usr/app

COPY --from=build /usr/app/dist/sailing-seas sailing-seas
CMD ["./sailing-seas"]
