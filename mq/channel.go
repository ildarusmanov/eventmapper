package mq

import (
	"eventmapper/models"
	"github.com/streadway/amqp"
)

type Channel struct {
	connection *amqp.Connection
	channel *amqp.Channel
}

func CreateChannelFromString(mqUrl string) (*Channel, error) {
	conn, err := CreateNewConnection(config.MqUrl)

	if err != nil {
		return nil, err
	}

	return CreateNewChannel(conn)
}

func CreateNewChannel(connection *amqp.Connection) (*Channel, error) {
	amqpCh, err := conn.Channel()

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

func (c *Channel) QueueDeclare(qName string) (*amqp.Queue, error) {
	return ch.QueueDeclare(
	  qName,    // name
	  false, // durable
	  false, // delete when usused
	  true,  // exclusive
	  false, // noWait
	  nil,   // arguments
	)
}

func (c *Channel) Publish(body []byte, qName, exName string) error {
	return ch.Publish(
			exName,  // exchange
			qName, // routing key
			false, // mandatory
			false,
			amqp.Publishing {
				DeliveryMode: amqp.Persistent,
				ContentType: "text/plain",
				Body: body,
			}
		)
}

func (c *Channel) PublishEvent(event *models.Event) error {
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

func (c *Channel) Close() {
	if c.channel != nil {
		c.channel.Close()
	}

	if c.connection != nil {
		c.connection.Close()
	}
}