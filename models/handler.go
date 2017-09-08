package models

import (
	"eventmapper/mq"
	"net/http"
	"bytes"
	"log"
)

const (
	HANDLER_TYPE_HTTP_JSON = "http_json"
)

type Handler struct {
	MqUrl       string             `yaml:"mq_url"`
	RKey        string             `yaml:"r_key"`
	HandlerType string             `yaml:"handler_type"`
	Options     map[string]string  `yaml:"handler_options"`
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
		MqUrl: mqUrl,
		RKey: rKey,
		HandlerType: handlerType,
		Options: options,
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
 * Start listening for events
 */
func (h *Handler) StartListening() {
	log.Printf("[x] Starting handler %s for %s", h.HandlerType, h.RKey)

	forever := make(chan bool)

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

	go func() {
		log.Printf("[x] Start listener")
		for m := range msgs {
			log.Printf("[x] %s", m.Body)
			h.process(m.Body)

		}
		log.Printf("[x] Finish listener")
	}()

	log.Printf("[x] Wait for messages")

	<- forever
}

/**
 * send event through defined transport
 * @param  eventBody   []byte
 */
func (h *Handler) process(eventBody []byte) {
	log.Printf("[x] New message received", eventBody)

	if h.HandlerType == HANDLER_TYPE_HTTP_JSON {
		h.httpJsonTransport(eventBody)
	} else {
		log.Printf("[x] Unknown handler type")
	}
}

/**
 * send event through HTTP JSON api
 * @param  eventBody   []byte
 */
func(h *Handler) httpJsonTransport(eventBody []byte) error {
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

	log.Printf("[x] POST %s - %s", h.Options["url"], resp.Status)
	
	return err
}
