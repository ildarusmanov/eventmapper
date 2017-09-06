package mq

import (
	"github.com/streadway/amqp"
)

type Event interface {
	GetBody() ([]byte, error)
}

type Channel struct {
	connection *amqp.Connection
	channel *amqp.Channel
}

func CreateChannelFromString(mqUrl string) (*Channel, error) {
	conn, err := CreateNewConnection(mqUrl)

	if err != nil {
		return nil, err
	}

	return CreateNewChannel(conn)
}

func CreateNewChannel(connection *amqp.Connection) (*Channel, error) {
	amqpCh, err := connection.Channel()

	if err != nil {
		return nil, err
	}

	return &Channel{connection, amqpCh}, nil
}

func (c *Channel) ExchangeDeclare(exName, exType string) error {
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

func (c *Channel) QueueDeclare(qName string) (amqp.Queue, error) {
	return c.channel.QueueDeclare(
		qName,    // name
		true, // durable
		false, // delete when usused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
}

func (c *Channel) Publish(body []byte, qName, exName string) error {
	return c.channel.Publish(
		exName,  // exchange
		qName, // routing key
		false, // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body: body,
		})
}

func (c *Channel) PublishEvent(event Event) error {
	q, err := c.QueueDeclare("events")

	if err != nil {
		return err
	}

	body, err := event.GetBody()

	if err != nil {
		return err
	}

	return c.Publish(body, q.Name, "") 
}

func (c *Channel) Consume(qName, cName string) (<-chan amqp.Delivery, error) {
    return c.channel.Consume(
        qName, // queue
        cName,     // consumer
        true,   // auto ack
        false,  // exclusive
        false,  // no local
        false,  // no wait
        nil,    // args
    )
}

func (c *Channel) ConsumeEvents() (<-chan amqp.Delivery, error) {
	q, err := c.QueueDeclare("events")

	if err != nil {
		return nil, err
	}

	return c.Consume(q.Name, "")
}

func (c *Channel) Close() {
	if c.channel != nil {
		c.channel.Close()
	}

	if c.connection != nil {
		c.connection.Close()
	}
}