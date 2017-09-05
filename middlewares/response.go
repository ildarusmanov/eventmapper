package middlewares

import (
	"net/http"
)

type Response struct {
	contentType string
	charset string
}

func CreateNewOkResponse(contentType string, charset string) *Response {
	return &Response{contentType, charset}
}

func CreateNewJsonOkResponse() *Response {
	return CreateNewOkResponse("application/json", "UTF-8")
}

func (resp *Response) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", resp.contentType + "; charset=" + resp.charset)
	w.WriteHeader(http.StatusOK)
}
