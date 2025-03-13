I'll implement increment rewrite
Starting with the auth pod
auth-api will communicate via gRPC, or via RestAPI,
it'll depends on environment variable setting

# todos

[_] - finish the service for the users API
[_] - build a gRPC connection after the http
[_] - build a client side load balancer
[_] - Users-API is not really a gRPC server, only http. the only thing it have to know is if it'll use a http communication or gRPC, and that will be flexible
