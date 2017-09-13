package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"eventmapper/configs"
	"eventmapper/models"
	"eventmapper/mq"
	"eventmapper/pb"
	"eventmapper/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var InvalidEventRequestSignature = errors.New("Invalid EventRequest signature")

// server is used to implement helloworld.GreeterServer.
type GrpcServer struct {
	mqUrl string
	token string
}

func CreateNewGrpcServer(mqUrl, token string) *GrpcServer {
	return &GrpcServer{mqUrl, token}
}

func (s *GrpcServer) validateSignature(eventReq *pb.EventRequest) error {
	toMd5 := []byte(s.token + ":" + eventReq.GetUserToken())
	hasher := md5.New()
	hasher.Write(toMd5)
	hash := hex.EncodeToString(hasher.Sum(nil))
	if hash != eventReq.GetSignature() {
		return InvalidEventRequestSignature
	}
	return nil
}

func (s *GrpcServer) publishEvent(rKey string, pbEvent *pb.Event) (mq.Event, error) {
	event := models.BuildNewEvent(
		pbEvent.GetEventName(),
		pbEvent.GetEventTarget(),
		pbEvent.GetUserId(),
		pbEvent.GetCreatedAt(),
		pbEvent.GetParams(),
	)

	return services.PublishEvent(event, s.mqUrl, rKey)
}

func (s *GrpcServer) buildResponse(isOk bool, status string) *pb.EventResponse {
	return &pb.EventResponse{isOk, status}
}

// SayHello implements helloworld.GreeterServer
func (s *GrpcServer) CreateEvent(ctx context.Context, in *pb.EventRequest) (*pb.EventResponse, error) {
	if err := s.validateSignature(in); err != nil {
		return s.buildResponse(false, "invalid signature"), err
	}

	if _, err := s.publishEvent(in.GetRKey(), in.GetEvent()); err != nil {
		return s.buildResponse(false, "publish error"), err
	} else {
		return s.buildResponse(true, "ok"), nil
	}
}

func StartGrpcServer(config *configs.Config) {
	lis, err := net.Listen("tcp", config.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEventMapperServer(s, CreateNewGrpcServer(config.MqUrl, config.GrpcToken))
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
