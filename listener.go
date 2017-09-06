package main

import (
	"eventmapper/models"
	"eventmapper/mq"

	"encoding/json"
	"fmt"
)

func StartEventsListener(mqChannel *mq.Channel) {
	msgs, err := mqChannel.ConsumeEvents()

	if err != nil {
		return
	}

	forever := make(chan bool)

	go func() {
		for m := range msgs {
			processEvent(m.Body)
		}
	}()

	<-forever
}

func processEvent(body []byte) {
	event := models.CreateNewEvent()

	if err := json.Unmarshal(body, event); err != nil {
		panic(err)
	}

	fmt.Println(event)
}
