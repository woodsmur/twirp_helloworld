# generate pb files

    for i in `ls rpc`; do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/${i}/service.proto; done

# run specific server

    go run cmd/helloworldserver/main.go

# run specific client

    go run cmd/helloworldclient/main.go