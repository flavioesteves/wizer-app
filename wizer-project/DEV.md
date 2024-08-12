# GRPC TOOL
* TERMINAL:  grpcurl -plaintext -d '{"key":"value"}' "HOST_GRPC_SERVICE" proto.protoService.rpc
* GUI: grpcui -plaintext "HOST_GRPC_SERVICE"


# REST TOOL
* TERMINAL: curl --header "Content-Type: application/json" \
    -- request POST \
    -- data '{"key":"value}' \
    "BROKER_SERVICE_ENDPOINT" 
    example: BROKER_SERVICE_ENDPOINT = http://0.0.0.0:8080/v1/profiles
* GUI: Postman


# CLI Wizer-Project
* make up
* make up_build
* make down
* make [APP]:  build_broker, build_user, build_front

* make proto-all
* make [SERVICE]: broker-service, user-service


# Migrate tool for Database
[migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
