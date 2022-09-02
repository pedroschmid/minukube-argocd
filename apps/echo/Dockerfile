# Go
FROM golang:1.19-alpine AS build

WORKDIR /opt/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server server.go

# Alpine
FROM alpine:3.15.6

WORKDIR /opt/app

COPY --from=build /opt/app/server server
CMD ["./server"]