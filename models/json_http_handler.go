package models

import (
	"bytes"
	"errors"
	"log"
	"net/http"
)

var IncorrectOptions = errors.New("Incorrect options")

/**
 * Handler implementation for HTTP JSON handlers
 */
type JsonHttpHandler struct {
	Options map[string]string
}

/**
 * get options
 * @return map[string]string
 */
func (h *JsonHttpHandler) GetOptions() map[string]string {
	return h.Options
}

/**
 * get RabbitMQ connection url
 * @return string
 */
func (h *JsonHttpHandler) GetMqUrl() string {
	return h.Options["mq_url"]
}

/**
 * Get AMQP routing key
 * @return string
 */
func (h *JsonHttpHandler) GetRKey() string {
	return h.Options["r_key"]
}

/**
 * Load options from map
 * @param  options map[string]string
 */
func (h *JsonHttpHandler) Init() error {
	if _, ok := h.Options["r_key"]; !ok {
		return IncorrectOptions
	}

	if _, ok := h.Options["mq_url"]; !ok {
		return IncorrectOptions
	}

	return nil
}

/**
 * send event through HTTP JSON api
 * @param  eventBody   []byte
 */
func (h *JsonHttpHandler) ProcessMessage(eventBody []byte) error {
	req, err := h.BuildRequest(eventBody)

	if err != nil {
		log.Printf("[x] %s", err)

		return err
	}
	resp, err := h.SendHttpRequest(req)

	if err != nil {
		log.Printf("[x] %s", err)

		return err
	}

	log.Printf("[x] POST %s", h.getUrl(), resp.Status)

	return err
}

/**
 * runs on handler start
 * @return error
 */
func (h *JsonHttpHandler) Start() error {
	return nil
}

/**
 * runs on handler finished
 */
func (h *JsonHttpHandler) Stop() {
	return
}

/**
 * Create http request
 * @param  eventBody []byte
 * @return *http.Request, error
 */
func (h *JsonHttpHandler) BuildHttpRequest(eventBody []byte) (*http.Request, error) {
	req, err := http.NewRequest(
		"POST",
		h.getUrl(),
		bytes.NewReader(eventBody),
	)

	if err != nil {
		return nil, err
	}

	if h.hasBasicAuth() {
		req.SetBasicAuth(h.getAuthUName(), h.getAuthPwd())
	}

	return req, nil
}

/**
 * Send request via http
 * @param  r *http.Request
 * @return *http.Response, error
 */
func (h *JsonHttpHandler) SendHttpRequest(r *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(r)
}

/**
 * Get handler url
 * @return string
 */
func (h *JsonHttpHandler) getUrl() string {
	return h.Options["url"]
}

/**
 * Check Http BasicAuth is enabled
 * @return bool
 */
func (h *JsonHttpHandler) hasBasicAuth() bool {
	_, ok := h.Options["auth_name"]

	return ok
}

/**
 * Get http basic auth username
 * @return string
 */
func (h *JsonHttpHandler) getAuthUName() string {
	return h.Options["auth_uname"]
}

/**
 * Get http basic auth password
 * @return string
 */
func (h *JsonHttpHandler) getAuthPwd() string {
	return h.Options["auth_pwd"]
}
