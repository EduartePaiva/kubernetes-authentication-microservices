I'll implement increment rewrite
Starting with the auth pod
auth-api will communicate via gRPC, or via RestAPI,
it'll depends on environment variable setting

# todos

[_] - finish the service for the users API
[_] - build a gRPC connection after the http
[_] - build a client side load balancer
[_] - Users-API is not really a gRPC server, only http. the only thing it have to know is if it'll use a http communication or gRPC, and that will be flexible

# important commands

## how to enter a container to test dns:

docker exec -it <container_id_or_name> sh

then I'll do: nslookup <some_dns>

## how to enter a kubernetes network to test dns

1. create a pod using:

kubectl apply -f dns-test-pod.yaml

then access the test pod:  
kubectl exec -it dnsutils -- /bin/sh

Then I can use nslookup inside it!!!.
