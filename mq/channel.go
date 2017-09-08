package mq

import (
	"github.com/streadway/amqp"
	"strings"
	"log"
)

const (
	EVENTS_EXCH_NAME = "events_topic"
	EVENTS_EXCH_TYPE = "topic"
)

type Event interface {
	GetEventName() string
	GetBody() ([]byte, error)
}

type Channel struct {
	channel *amqp.Channel
}

func CreateNewChannel(mqConn *amqp.Connection) (*Channel, error) {
	mqCh, err := mqConn.Channel()

	if err != nil {
		return nil, err
	}

	return &Channel{mqCh}, nil
}

func (c *Channel) PublishEvent(event Event, rKey string) error {
	err := c.exchangeDeclare(EVENTS_EXCH_NAME, EVENTS_EXCH_TYPE)

	if err != nil {
		return err
	}

	fullRKey := strings.Join([]string{rKey, event.GetEventName()}, ".")
	body, err := event.GetBody()

	if err != nil {
		return err
	}

	log.Printf("[x] Sending %s", body)

	return c.publish(body, fullRKey, EVENTS_EXCH_NAME)
}

func (c *Channel) ConsumeEvents(rKey string) (<-chan amqp.Delivery, error) {
	err := c.exchangeDeclare(EVENTS_EXCH_NAME, EVENTS_EXCH_TYPE)

	if err != nil {
		return nil, err
	}

	q, err := c.queueDeclare("")

	if err != nil {
		return nil, err
	}

	err = c.queueBind(EVENTS_EXCH_NAME, q.Name, rKey)

	if err != nil {
		return nil, err
	}

	return c.consume(q.Name, "")
}

func (c *Channel) Close() {
	c.channel.Close()
}

func (c *Channel) exchangeDeclare(exName, exType string) error {
	return c.channel.ExchangeDeclare(
        exName, // name
        exType, // type
        true, // durable
        false, // auto-deleted
        false, // internal
        false, // no-wait
        nil, // arguments
    )
}

func (c *Channel) queueDeclare(qName string) (amqp.Queue, error) {
	return c.channel.QueueDeclare(
		qName,    // name
		true, // durable
		false, // delete when usused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
}

func (c *Channel) queueBind(exName, qName, rKey string) error {
	return c.channel.QueueBind(
		qName, // queue name
		rKey,     // routing key
		exName, // exchange
		false,
		nil,
	)
}

func (c *Channel) publish(body []byte, rKey, exName string) error {
	return c.channel.Publish(
		exName,  // exchange
		rKey, // routing key
		false, // mandatory
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: body,
		})
}

func (c *Channel) consume(qName, cName string) (<-chan amqp.Delivery, error) {
    return c.channel.Consume(
        qName,  // queue
        cName,  // consumer
        true,   // auto ack
        false,  // exclusive
        false,  // no local
        false,  // no wait
        nil,    // args
    )
}
