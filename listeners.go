package main

import (
	"log"
	"strconv"

	"github.com/eclipse/paho.mqtt.golang"
)

func getProgressBarListener(topic string) func(mqtt.Client, mqtt.Message) {
	return func(client mqtt.Client, message mqtt.Message) {
		percent, err := strconv.Atoi(string(message.Payload()))
		if err != nil {
			log.Println("could not parse message as int")
			publishInnerErrorRing(client)
			return
		}

		if percent < 0 || percent > 100 {
			log.Printf("invalid percentage: %v\n", percent)
			publishInnerErrorRing(client)
		}

		numberOfPixels := 24
		ring, err := MakeGradientProgressBar(&V{255, 0, 0}, &V{0, 255, 0}, numberOfPixels, percent)
		if err != nil {
			publishInnerErrorRing(client)
		}

		msgToPublish := pixelSliceToMessage(ring)
		log.Println("Publishing progress bar status")
		client.Publish(innerRingTopic(topic), 0, false, msgToPublish)

	}
}
