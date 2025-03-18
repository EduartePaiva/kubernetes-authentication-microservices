# Authentication microservice with kubernetes deployment

![kubernetes cluster](./draw.png "kubernetes cluster")

This is a authentication microservice project it's divided in two separated internal services that communicate with each other and a external mongo database.

## Objective

The objective of this project is to be a reference on how to work with microservice that allow both gRPC or REST communication, while setting it up on different environments like docker-compose and kubernetes.

## Used technologies

- Golang
- gRPC
- Rest API
- MongoDB
- Docker
- Docker-Compose
- Kubernetes

## Auth-api

Auth api is responsible for handling tree authentication tasks which are in the following routes:

- **/hashed-pw/:password** this route accepds a password and return it's hashed counterpart.
- **/token** compare if a hashed password and an unhashed password are the same, and then returns a JWT token.
- **/verify-token** Verifying if a JWT token is valid.

Auth-api also have the possibility of exposing gRPC routing, if the environment variable **COMMUNICATION_PROTOCOL** is set to "gRPC"

## Users-api

It's responsible for handling users authentication requests from the web, it have the following apis:

- **/signup** it receives a email and a password, then it validate the credentials and call the **/hash-pw** from the auth-api, after that the it's saved in the mongodb the email and hashed password

- **/login** it queries the hashed password from the user email and then calls **/token** to verify if the password if valid and returns the JWT token.

Users-api communicate with auth-api via gRPC or Rest API. it's configurable by the **COMMUNICATION_PROTOCOL** environment variable.

## About the deployment

Both microservices was containerized with docker, and pushed to docker hub. Then the kubernetes definitions orchestrate the microservices for easy deployment using gRPC in a cluster.

## gRPC setting and load balancing.

The gRPC communication between the auth-api and users-api can be achieved if both have the environment variable **COMMUNICATION_PROTOCOL** with the value **gRPC**, otherwise it defaults to **REST**, when in gRPC the client (users-api) connects with all auth-api servers and uses the round robin load balancing strategy.

## Environment variables

### auth-api

| Variable Name            | Default Value | Description                                          |
| ------------------------ | ------------- | ---------------------------------------------------- |
| `TOKEN_KEY`              | (required)    | Secret key for signing JTW tokens, should be secure. |
| `AUTH_SERVER_PORT`       | `3000`        | The port that the server will listen to.             |
| `COMMUNICATION_PROTOCOL` | `REST`        | This value can be `REST` or `gRPC`                   |

### users-api

| Variable Name            | Default Value    | Description                                                                                                      |
| ------------------------ | ---------------- | ---------------------------------------------------------------------------------------------------------------- |
| `AUTH_API_ADDRESS`       | `localhost:3000` | The auth api address, if using gRPC this address would be a dns address e.g. `dns:///auth-service.default:50051` |
| `USERS_SERVER_PORT`      | `3001`           | The port that users-api will receive http requests                                                               |
| `COMMUNICATION_PROTOCOL` | `REST`           | This value can be `REST` or `gRPC`                                                                               |
| `MONGODB_CONNECTION_URI` | (required)       | MongoDB connection URI, for kubernetes it's highly recommended to setup a secretKey e.g. opaque key              |
