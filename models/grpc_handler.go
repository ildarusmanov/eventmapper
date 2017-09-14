package models

import (
	"encoding/json"
	"errors"
	"eventmapper/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var incorrectGrpcHandlerOptionsError = errors.New("Incorrect options")

/**
 * Handler implementation for HTTP JSON handlers
 */
type GrpcHandler struct {
	Options map[string]string
}

/**
 * get options
 * @return map[string]string
 */
func (h *GrpcHandler) GetOptions() map[string]string {
	return h.Options
}

/**
 * get RabbitMQ connection url
 * @return string
 */
func (h *GrpcHandler) GetMqUrl() string {
	return h.Options["mq_url"]
}

/**
 * Get AMQP routing key
 * @return string
 */
func (h *GrpcHandler) GetRKey() string {
	return h.Options["r_key"]
}

/**
 * Load options from map
 * @param  options map[string]string
 */
func (h *GrpcHandler) Init() error {
	if _, ok := h.Options["r_key"]; !ok {
		return incorrectGrpcHandlerOptionsError
	}

	if _, ok := h.Options["mq_url"]; !ok {
		return incorrectGrpcHandlerOptionsError
	}

	if _, ok := h.Options["user_token"]; !ok {
		return incorrectGrpcHandlerOptionsError
	}

	if _, ok := h.Options["addr"]; !ok {
		return incorrectGrpcHandlerOptionsError
	}

	if _, ok := h.Options["tls"]; !ok {
		return incorrectGrpcHandlerOptionsError
	}

	if _, ok := h.Options["cert"]; !ok {
		return incorrectGrpcHandlerOptionsError
	}

	if _, ok := h.Options["host"]; !ok {
		return incorrectGrpcHandlerOptionsError
	}

	return nil
}

/**
 * send event through gRPC
 * @param  eventBody   []byte
 */
func (h *GrpcHandler) ProcessMessage(eventBody []byte) error {
	opts, err := h.getConnOptions()

	if err != nil {
		return err
	}

	conn, err := grpc.Dial(h.getAddr(), opts...)

	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewEventMapperClient(conn)

	event := &pb.Event{}

	if err := json.Unmarshal(eventBody, event); err != nil {
		return err
	}

	eventReq := &pb.EventRequest{
		UserToken: h.getUserToken(),
		RKey:      "",
		Event:     event,
	}

	_, err = client.CreateEvent(context.Background(), eventReq)

	if err != nil {
		return err
	}

	return nil
}

/**
 * Get connection options
 */
func (h *GrpcHandler) getConnOptions() ([]grpc.DialOption, error) {
	var opts []grpc.DialOption

	if h.getTls() {
		creds, err := credentials.NewClientTLSFromFile(
			h.getCert(),
			h.getHost(),
		)

		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	return opts, nil
}

/**
 * runs on handler start
 * @return error
 */
func (h *GrpcHandler) Start() error {
	return nil
}

/**
 * runs on handler finished
 */
func (h *GrpcHandler) Stop() {
	return
}

/**
 * Get handler addr
 * @return string
 */
func (h *GrpcHandler) getAddr() string {
	return h.Options["addr"]
}

/**
 * Get handler user token
 * @return string
 */
func (h *GrpcHandler) getUserToken() string {
	return h.Options["user_token"]
}

/**
 * Get tls
 * @return bool
 */
func (h *GrpcHandler) getTls() bool {
	return h.Options["tls"] == "true"
}

/**
 * Get cert
 * @return string
 */
func (h *GrpcHandler) getCert() string {
	return h.Options["cert"]
}

/**
 * Get host
 * @return string
 */
func (h *GrpcHandler) getHost() string {
	return h.Options["host"]
}
