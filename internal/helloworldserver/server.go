package helloworldserver

import (
	"context"
	"fmt"

	pb "github.com/woodsmur/twirp_helloworld/rpc/helloworld"
)

type Server struct {
}

// SayHello says hello to client
func (s *Server) SayHello(ctx context.Context, req *pb.Hello) (*pb.GoodBye, error) {
	fmt.Printf("[SayHello]Req : %#v", req)
	return &pb.GoodBye{
		UserID:  req.GetUserID(),
		Message: req.GetMessage(),
	}, nil
}

// SayHello2 says hello to client
func (s *Server) SayHelloTypesAsField(ctx context.Context, req *pb.ReqHello) (*pb.ResHello, error) {
	fmt.Printf("[SayHello2]Req : %#v", req)
	return &pb.ResHello{
		GoodBye: &pb.GoodBye{
			UserID:  req.GetHello().GetUserID(),
			Message: req.GetHello().GetMessage(),
		},
	}, nil
}
