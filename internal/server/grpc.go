package server

import (
	"leo/api/snippet/v1"
	"leo/internal/service"
	"net"
	"net/url"
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(ss *service.SnippetService) (*grpc.Server, error) {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	dir, error := os.UserHomeDir()
	if error != nil {
		return nil, error
	}
	socketPath := filepath.Join(dir, ".peon/leo.sock")
	os.Remove(socketPath)
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.Address("unix://"+socketPath))
	opts = append(opts, grpc.Listener(l))
	u, err := url.Parse("unix://" + socketPath)
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.Endpoint(u))
	srv := grpc.NewServer(opts...)

	snippet.RegisterSnippetServer(srv, ss)
	return srv, nil
}
