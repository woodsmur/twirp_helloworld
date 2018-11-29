package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/twitchtv/twirp"
	"github.com/woodsmur/twirp_helloworld/rpc/helloworld"
)

func main() {
	header := make(http.Header)
	header.Set("UID", "uDRlDxQYbFVXarBvmTncBoWKcZKqrZTY")
	header.Set("AppKey", "1")
	ctx := context.Background()
	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)

	client := helloworld.NewAwesomeTwirpServiceProtobufClient("http://localhost:8080", &http.Client{})
	resp, err := client.SayHelloTypesAsField(ctx, &helloworld.ReqHello{
		Hello: &helloworld.Hello{
			UserID:  "1",
			Message: "Hello World Protobuf",
		},
	})

	if err != nil {
		fmt.Printf("client call error, err : %v\n", err)
		return
	}

	fmt.Printf("%v\n", resp.GetGoodBye())

	//
	resp2, err := client.SayHello(ctx, &helloworld.Hello{
		UserID:  "1",
		Message: "Hello World Protobuf",
	})

	if err != nil {
		fmt.Printf("client call error, err : %v\n", err)
		return
	}

	fmt.Printf("%v\n", resp2)

}
