package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/eclipse/paho.mqtt.golang"
)

func innerRingTopic(id string) string {
	return fmt.Sprintf("ecowatch/%v/set/pixels", id)
}

func pixelSliceToMessage(pSlice []Pixel) []byte {
	message := EcowatchMessage{
		Inner: pSlice,
	}
	jsonData, _ := json.Marshal(message)
	return jsonData

}

func publishInnerErrorRing(client mqtt.Client, ecoWatchId string) {
	topic := innerRingTopic(ecoWatchId)
	errorRing := pixelSliceToMessage(SingleColourPixelSlice(&V{255, 0, 0}, 24))
	client.Publish(topic, 0, false, errorRing)
}

func GetProgressBarListener(ecoWatchId string) func(mqtt.Client, mqtt.Message) {
	return func(client mqtt.Client, message mqtt.Message) {
		percent, err := strconv.Atoi(string(message.Payload()))
		if err != nil {
			log.Println("could not parse message as int")
			publishInnerErrorRing(client, ecoWatchId)
			return
		}

		if percent < 0 || percent > 100 {
			log.Printf("invalid percentage: %v\n", percent)
			publishInnerErrorRing(client, ecoWatchId)
		}

		numberOfPixels := 24
		ring, err := GradientPixelSliceProgress(&V{255, 0, 0}, &V{0, 255, 0}, numberOfPixels, percent)
		if err != nil {
			log.Printf("Could not generate progress bar:%v", err)
			publishInnerErrorRing(client, ecoWatchId)
		}

		msgToPublish := pixelSliceToMessage(ring)
		topic := innerRingTopic(ecoWatchId)
		log.Printf("Publishing progress bar status to %v\n", topic)
		client.Publish(topic, 0, false, msgToPublish)
	}
}
