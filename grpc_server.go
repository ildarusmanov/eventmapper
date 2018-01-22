package main

import (
	"github.com/ildarusmanov/eventmapper/configs"
	"github.com/ildarusmanov/eventmapper/models"
	"github.com/ildarusmanov/eventmapper/mq"
	"github.com/ildarusmanov/eventmapper/pb"
	"github.com/ildarusmanov/eventmapper/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GrpcServer struct {
	mqUrl string
}

/**
 * create event method
 * @param  ctx context.Context
 * @param in *pb.EventRequest
 * @return *pb.EventResponse, error
 */
func (s *GrpcServer) CreateEvent(ctx context.Context, in *pb.EventRequest) (*pb.EventResponse, error) {
	if _, err := s.publishEvent(in.GetRKey(), in.GetEvent()); err != nil {
		return s.buildResponse(false, "publish error"), err
	} else {
		return s.buildResponse(true, "ok"), nil
	}
}

/**
 * start server
 * @param config *configs.Config
 */
func StartGrpc(config *configs.Config) {
	go runGrpcServer(config)
}

/**
 * start GRPC server listnening for requests
 * @param  config *configs.Config
 */
func runGrpcServer(config *configs.Config) {
	lis, err := net.Listen("tcp", config.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts, err := getGrpcServerOptions(config)

	if err != nil {
		log.Fatalf("failed with %v", err)
	}

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterEventMapperServer(grpcServer, createNewGrpcServer(config.MqUrl))
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

/**
 * create new grpc server
 * @param mqUrl string RabbitMQ connection string
 * @return *GrpcServer
 */
func createNewGrpcServer(mqUrl string) *GrpcServer {
	return &GrpcServer{mqUrl}
}

/**
 * publish event
 * @param  rKey string
 * @param pbEvent *pb.Event
 * @return mq.Event, error
 */
func (s *GrpcServer) publishEvent(rKey string, pbEvent *pb.Event) (mq.Event, error) {
	event := models.BuildNewEvent(
		models.CreateNewEventSource(
			pbEvent.GetSource().GetSourceType(),
			pbEvent.GetSource().GetSourceId(),
			pbEvent.GetSource().GetOrigin(),
			pbEvent.GetSource().GetParams(),
		),
		models.CreateNewEventTarget(
			pbEvent.GetTarget().GetTargetType(),
			pbEvent.GetTarget().GetTargetId(),
			pbEvent.GetTarget().GetParams(),
		),
		pbEvent.GetEventName(),
		pbEvent.GetUserId(),
		pbEvent.GetCreatedAt(),
		pbEvent.GetParams(),
	)

	return services.PublishEvent(event, s.mqUrl, rKey)
}

/**
 * build response
 * @param isOk bool
 * @param status string
 * @return *pb.EventResponse
 */
func (s *GrpcServer) buildResponse(isOk bool, status string) *pb.EventResponse {
	return &pb.EventResponse{isOk, status}
}

/**
 * get server options
 * @param  config *configs.Config
 * @return []grpc.ServerOption, err
 */
func getGrpcServerOptions(config *configs.Config) ([]grpc.ServerOption, error) {
	var opts []grpc.ServerOption

	if config.GrpcTls {
		creds, err := credentials.NewServerTLSFromFile(
			config.GrpcCertFile,
			config.GrpcKeyFile,
		)

		if err != nil {
			return nil, err
		}

		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	return opts, nil
}
