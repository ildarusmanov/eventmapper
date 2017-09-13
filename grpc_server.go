package main

import (
	"eventmapper/mq"
	"eventmapper/pb"
	"eventmapper/services"
	pb "eventmapper/event_service"
	"log"
	"net"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var InvalidEventRequestSignature = errors.New("Invalid EventRequest signature")
// server is used to implement helloworld.GreeterServer.
type GrpcServer struct {}

func CreateNewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) validateSignature(eventReq *pb.EventRequest) error {
	return InvalidEventRequestSignature
}

func (s *GrpcServer) publishEvent(pbEvent *pb.Event) (mq.Event, error) {
	return nil, nil
}

func (s *GrpcServer) buildResponse(event mq.Event) (*pb.EventResponse, error) {
	return nil, nil
}
// SayHello implements helloworld.GreeterServer
func (s *GrpcServer) CreateEvent(ctx context.Context, in *pb.EventRequest) (*pb.EventResponse, error) {
	if err := s.validateSignature(in); err != nil {
		return nil, err
	}

	if event, err := s.publishEvent(in); err != nil {
		return nil, err
	} else {
		return s.buildResponse(event)
	}
}

func StartGrpcServer(config *configs.Config) {
	lis, err := net.Listen("tcp", config.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEventMapperServer(s, CreateNewGrpcServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
