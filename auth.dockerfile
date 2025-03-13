FROM golang:1.24-alpine AS build

WORKDIR /app

# copy only mod files
COPY go.work go.work.sum ./
COPY auth-api/go.mod auth-api/go.mod
COPY users-api/go.mod users-api/go.mod
COPY common/go.mod common/go.mod

# download external packages
WORKDIR /app/users-api
RUN go mod download

# copy extra files
WORKDIR /app
COPY common/ common/
COPY auth-api/ auth-api/

# build go binary
WORKDIR /app/auth-api
RUN mkdir -p deploy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./deploy/app ./

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/auth-api/deploy/app /app/app

ENTRYPOINT [ "./app" ]