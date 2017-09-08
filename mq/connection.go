package mq

import (
	"github.com/streadway/amqp"
)

func CreateNewConnection(url string) (*amqp.Connection, error) {
	return amqp.Dial(url)
}
