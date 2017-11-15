package models

import (
	"errors"
	"eventmapper/mq"
	"log"
)

const (
	HANDLER_TYPE_HTTP_JSON = "http_json"
	HANDLER_TYPE_GRPC      = "grpc"
)

var UndefinedHandlerError = errors.New("Undefined handler type")

type Handler interface {
	Init() error
	Start() error
	Stop()
	ProcessMessage([]byte) error
	GetOptions() map[string]string
	GetMqUrl() string
	GetRKey() string
}

/**
 * Create new Handler
 * @param options map[string]string
 * required keys for options map: mq_url, r_key, handler_type
 * @return *Handler
 */
func CreateNewHandler(options map[string]string) (Handler, error) {
	if options["handler_type"] == HANDLER_TYPE_HTTP_JSON {
		h := &JsonHttpHandler{options}

		if err := h.Init(); err != nil {
			return nil, err
		}

		return h, nil
	}

	if options["handler_type"] == HANDLER_TYPE_GRPC {
		h := &GrpcHandler{options}

		if err := h.Init(); err != nil {
			return nil, err
		}

		return h, nil
	}

	return nil, UndefinedHandlerError
}

/**
 * Start listening for messages
 * @param h Handler
 * @param closeCh chan bool
 * @param errCh chan error
 */
func StartHandler(options map[string]string, closeCh chan bool) error {
	h, err := CreateNewHandler(options)

	if err != nil {
		return err
	}

	err = h.Start()

	if err != nil {
		return err
	}

	defer h.Stop()

	log.Printf("[x] Starting handler for %s", h.GetRKey())

	mqConn, err := mq.CreateNewConnection(h.GetMqUrl())
	defer mqConn.Close()

	if err != nil {
		return err
	}

	mqChannel, err := mq.CreateNewChannel(mqConn)
	defer mqChannel.Close()

	if err != nil {
		return err
	}

	msgs, err := mqChannel.ConsumeEvents(h.GetRKey())

	if err != nil {
		return err
	}

	log.Printf("[x] Start listener")

	for m := range msgs {
		log.Printf("[x] %s", m.Body)

		select {
		case <-closeCh:
			return nil
		default:
			log.Printf("[x] error: %s", h.ProcessMessage(m.Body))
		}
	}

	log.Printf("[x] Finish listener")

	return nil
}
