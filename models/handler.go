package models

import (
	"bytes"
	"eventmapper/mq"
	"log"
	"net/http"
)

const (
	HANDLER_TYPE_HTTP_JSON = "http_json"
)

type Handler struct {
	MqUrl       string            `yaml:"mq_url"`
	RKey        string            `yaml:"r_key"`
	HandlerType string            `yaml:"handler_type"`
	Options     map[string]string `yaml:"handler_options"`
}

/**
 * Create new Handler
 * @param mqUrl string
 * @param rKey string
 * @param handlerType string
 * @param options map[string]string
 * @return *Handler
 */
func CreateNewHandler(mqUrl string, rKey string, handlerType string, options map[string]string) *Handler {
	return &Handler{
		MqUrl:       mqUrl,
		RKey:        rKey,
		HandlerType: handlerType,
		Options:     options,
	}
}

/**
 * Create handler by given config map
 * @param cfg map[string]string
 * @return *Handler
 */
func BuildHandlerFromConfig(cfg map[string]string) *Handler {
	return CreateNewHandler(cfg["mq_url"], cfg["r_key"], cfg["handler_type"], cfg)
}

/**
 * Start listener
 * @param closeCh chan bool
 * @param errCh chan error
 */
func (h *Handler) StartListening(closeCh chan bool, errCh chan error) {
	log.Printf("[x] Starting handler %s for %s", h.HandlerType, h.RKey)

	mqConn, err := mq.CreateNewConnection(h.MqUrl)
	defer mqConn.Close()

	if err != nil {
		panic(err)
	}

	mqChannel, err := mq.CreateNewChannel(mqConn)
	defer mqChannel.Close()

	if err != nil {
		panic(err)
	}

	msgs, err := mqChannel.ConsumeEvents(h.RKey)

	if err != nil {
		panic(err)
	}

	log.Printf("[x] Start listener")

	for m := range msgs {
		select {
		case <-closeCh:
			return
		default:
			log.Printf("[x] %s", m.Body)
			errCh <- h.process(m.Body)
		}
	}

	log.Printf("[x] Finish listener")
}

/**
 * send event through defined transport
 * @param  eventBody   []byte
 */
func (h *Handler) process(eventBody []byte) error {
	log.Printf("[x] New message received %s", eventBody)

	if h.HandlerType == HANDLER_TYPE_HTTP_JSON {
		return h.httpJsonTransport(eventBody)
	}

	log.Printf("[x] Unknown handler type")

	return nil
}

/**
 * send event through HTTP JSON api
 * @param  eventBody   []byte
 */
func (h *Handler) httpJsonTransport(eventBody []byte) error {
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		h.Options["url"],
		bytes.NewReader(eventBody),
	)

	if err != nil {
		return err
	}

	if _, ok := h.Options["auth_uname"]; ok {
		req.SetBasicAuth(h.Options["auth_uname"], h.Options["auth_pwd"])
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("[x] %s", err)

		return err
	}

	log.Printf("[x] POST %s", h.Options["url"], resp.Status)

	return err
}
