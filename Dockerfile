FROM golang:1.21-alpine as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/platform-games cmd/main.go

FROM alpine
WORKDIR /etc/platform-games
COPY --from=build /usr/local/bin/platform-games /usr/local/bin/platform-games
CMD ["platform-games"]
