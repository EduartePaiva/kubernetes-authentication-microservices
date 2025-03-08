FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.work go.work.sum ./

COPY auth-api/ auth-api/
COPY common/ common/
COPY users-api/go.mod users-api/go.mod

WORKDIR /app/auth-api

# RUN go mod tidy 
RUN go mod download

RUN mkdir -p deploy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./deploy/app ./

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/auth-api/deploy/app /app/app

ENTRYPOINT [ "./app" ]