package models

import (
	"testing"
)

func TestJsonHttpHandlerGetters(t *testing.T) {
	o := map[string]string{
		"mq_url":       "mq_url",
		"r_key":        "r_key",
		"handler_type": "http_json",
		"url":          "url",
	}

	h, err := CreateNewHandler(o)

	if err != nil {
		t.Error(err)
	}

	ho := h.GetOptions()

	if o["mq_url"] != ho["mq_url"] {
		t.Error("Error")
	}

	if o["r_key"] != ho["r_key"] {
		t.Error("Error")
	}

	if o["handler_type"] != ho["handler_type"] {
		t.Error("Error")
	}

	if o["url"] != ho["url"] {
		t.Error("Error")
	}

	if h.GetMqUrl() != o["mq_url"] {
		t.Error("Error")
	}

	if h.GetRKey() != o["r_key"] {
		t.Error("Error")
	}
}
