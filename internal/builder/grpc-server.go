package builder

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	impl "supercollector/internal/grpc"
	"supercollector/internal/pb"
)

func NewGRPCServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		return errors.Wrap(err, "listen grpc")
	}

	s := grpc.NewServer()

	pb.RegisterSupercollectorServer(s, &impl.Server{})
	if err = s.Serve(lis); err != nil {
		return errors.Wrap(err, "register server impl")
	}

	return nil
}
