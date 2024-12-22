package builder

import (
	"context"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"supercollector/internal/pb"
)

func NewRESTServer(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterSupercollectorHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	if err != nil {
		return errors.Wrap(err, "register server")
	}

	if err = http.ListenAndServe(":8081", mux); err != nil {
		return errors.Wrap(err, "listen and serve")
	}

	return nil
}
