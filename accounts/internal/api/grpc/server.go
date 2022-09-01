package grpc

import (
	pb "accounts/internal/api/gen/proto"
	"accounts/internal/utils"
	"context"
)

type GRPCServer struct {
	pb.UnimplementedAccountsServer
	//grpcServer *grpc.Server
}

func (s *GRPCServer) Test(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	utils.Logger.Info("Start test GRPC")
	return &pb.Response{
		Z: in.X + in.Y,
	}, nil
}

//type GRPCServer struct {
//	grpcServer *grpc.Server
//}
//
//func NewGRPCServer() *GRPCServer {
//	return &GRPCServer{}
//}

//
//func (s *GRPCServer) ListenAndServe(ctx context.Context, addr string) error {
//	// logger "server started"
//
//	lis, err := net.Listen("tcp", addr)
//	if err != nil {
//		return err
//	}
//	s.grpcServer = grpc.NewServer()
//	reflection.Register(s.grpcServer)
//	Account.RegisterAccountsServer(s.grpcServer, s)
//
//	return s.grpcServer.Serve(lis)
//}
//
//func (s *GRPCServer) Close(ctx context.Context) {
//	// logger "shutting down"
//	if s.grpcServer != nil {
//		s.grpcServer.GracefulStop()
//	}
//}
//
//func RunGRPCApi(ctx context.Context, cfg *config.GRPC, grpcAPIServer *GRPCServer) {
//	go func() {
//		<-ctx.Done()
//		grpcAPIServer.Close(ctx)
//		logger.NewEntry().Info(ctx, "grpc api server gracefully stopped")
//	}()
//
//	logger.NewEntry().Info(ctx, fmt.Sprintf("gRPC listening on %s:%s",
//		cfg.Host, cfg.Port))
//	err := grpcAPIServer.ListenAndServe(ctx, fmt.Sprintf("%s:%s",
//		cfg.Host, cfg.Port))
//	if err != nil {
//		logger.NewEntry().WithError(err).Error(ctx, "error starting grpc server")
//	}
//}
