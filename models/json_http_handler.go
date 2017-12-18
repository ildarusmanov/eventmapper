package models

import (
	"strconv"
	"bytes"
	"errors"
	"log"
	"net/http"
	"crypto/tls"
)

var incorrectJsonHttpHandlerOptionsError = errors.New("Incorrect options")

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

func (h *JsonHttpHandler) getOptionByKey(key string) (string, error) {
	if _, ok := h.Options[key]; !ok {
		return "", incorrectJsonHttpHandlerOptionsError
	}

	return h.Options[key], nil
}

/**
 * get RabbitMQ connection url
 * @return string
 */
func (h *JsonHttpHandler) GetMqUrl() string {
	v, _ := h.getOptionByKey("mq_url")

	return v
}

/**
 * Get AMQP routing key
 * @return string
 */
func (h *JsonHttpHandler) GetRKey() string {
	v, _ := h.getOptionByKey("r_key")

	return v
}

/**
 * Get processes count
 * @return int
 */
func (h *JsonHttpHandler) GetPCount() int {
	pCountStr, err := h.getOptionByKey("p_count")

	if err != nil {
		return 1
	}

	pCountInt, err := strconv.Atoi(pCountStr)

	if err != nil {
		return 1
	}

	return pCountInt
}

/**
 * Get TLS flag
 * @return bool
 */
func (h *JsonHttpHandler) getTLS() (bool, error) {
	tlsStr, err := h.getOptionByKey("TLS")

	if err != nil {
		return false, err
	}

	tlsBool, err := strconv.ParseBool(tlsStr)

	if err != nil {
		return false, err
	}

	return tlsBool, nil
}

/**
 * Load options from map
 * @param  options map[string]string
 */
func (h *JsonHttpHandler) Init() error {
	if _, err := h.getOptionByKey("r_key"); err != nil {
		return err
	}

	if _, err := h.getOptionByKey("mq_url"); err != nil {
		return err
	}

	if _, err := h.getOptionByKey("url"); err != nil {
		return err
	}

	return nil
}

/**
 * send event through HTTP JSON api
 * @param  eventBody   []byte
 */
func (h *JsonHttpHandler) ProcessMessage(eventBody []byte) error {
	req, err := h.buildHttpRequest(eventBody)

	if err != nil {
		log.Printf("[x] %s", err)

		return err
	}
	resp, err := h.sendHttpRequest(req)

	if err != nil {
		log.Printf("[x] %s", err)

		return err
	}

	defer resp.Body.Close()

	log.Printf("[x] POST %s", req.URL.String(), resp.Status)

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
func (h *JsonHttpHandler) buildHttpRequest(eventBody []byte) (*http.Request, error) {
	url, err := h.getUrl()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewReader(eventBody),
	)

	if err != nil {
		return nil, err
	}

	h.setAuthData(req)

	return req, nil
}

/**
 * Send request via http
 * @param  r *http.Request
 * @return *http.Response, error
 */
func (h *JsonHttpHandler) sendHttpRequest(r *http.Request) (*http.Response, error) {
    hasTls, err := h.getTLS()

    if err != nil {
    	return nil, err
    }

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: hasTls},
    }

	client := &http.Client{Transport: tr}

	return client.Do(r)
}

/**
 * Get handler url
 * @return string
 */
func (h *JsonHttpHandler) getUrl() (string, error) {
	return h.getOptionByKey("url")
}

/**
 * get authentication type
 * @return string, error
 */
func (h *JsonHttpHandler) getAuthType() (string, error) {
	return h.getOptionByKey("auth_type")
}

/**
 * add auth data into HTTP request
 * @param  *http.Request
 */
func (h *JsonHttpHandler) setAuthData(req *http.Request) {
	if _, err := h.getAuthType(); err != nil {
		return
	}

	switch authType, _ := h.getAuthType(); authType {
	case "http-basic":
		h.setHttpBasicAuthData(req)
	case "get-token":
		h.setGetTokenAuthData(req)
	case "header-token":
		h.setHeaderTokenAuthData(req)
	}
}

func (h *JsonHttpHandler) setHttpBasicAuthData(req *http.Request) error {
    username, err := h.getOptionByKey("username")

    if err != nil {
    	return err
    }

    password, err := h.getOptionByKey("username")

    if err != nil {
    	return err
    }

    req.SetBasicAuth(username, password)

    return nil
}

func (h *JsonHttpHandler) setGetTokenAuthData(req *http.Request) error {
	tokenValue, err := h.getOptionByKey("token_value")

	if err != nil {
		return err
	}

	tokenName, err := h.getOptionByKey("token_name")

	if err != nil {
		return err
	}

	req.URL.Query().Set(tokenName, tokenValue)

	return nil
}

func (h *JsonHttpHandler) setHeaderTokenAuthData(req *http.Request) error {
	tokenValue, err := h.getOptionByKey("token_value")

	if err != nil {
		return err
	}

	tokenHeader, err := h.getOptionByKey("token_header")

	if err != nil {
		return err
	}

	req.Header.Set(tokenHeader, tokenValue)

	return nil
}