package services

import (
	"eventmapper/mq"
)

func PublishEvent(event *mq.Event, mqUrl, rKey string) (mq.Event, error) {
	mqConn, err := mq.CreateNewConnection(mqUrl)
	defer mqConn.Close()

	if err != nil {
		return nil, err
	}

	mqChannel, err := mq.CreateNewChannel(mqConn)
	defer mqChannel.Close()

	if err != nil {
		return nil, err
	}

	if err := event.Validate(); err != nil {
		return nil, err
	}

	if err := event.Publish(mqChannel, rKey); err != nil {
		return nil, err
	}

	return event, nil
}
