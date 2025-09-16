package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

// d
var progressBarTopic = "go2ecowatch/inner/progressbar"

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

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

func init() {
	viper.SetEnvPrefix("G2E")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/go2ecowatch")
	viper.BindEnv("HOST")
	viper.BindEnv("PORT")
	viper.BindEnv("USER")
	viper.BindEnv("PASSWORD")
	viper.BindEnv("ECOWATCH_ID")

}

func main() {
	viper.SetDefault("HOST", "")
	viper.SetDefault("PORT", "1883")
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Println("No config file found in $HOME/.ha2ecowatch/config.yaml, not using auth")
		}
	}
	username := viper.GetString("USER")
	password := viper.GetString("PASSWORD")
	host := viper.Get("HOST")
	port := viper.Get("PORT")
	ecowatchId := viper.GetString("ECOWATCH_ID")
	connectionString := fmt.Sprintf("mqtt://%v:%v", host, port)

	fmt.Println(connectionString)

	opts := mqtt.NewClientOptions().
		SetUsername(username).
		SetPassword(password).
		AddBroker(connectionString).
		SetAutoReconnect(true)
	opts.OnConnect = connectHandler

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	subscribeTopic := fmt.Sprintf("%v", progressBarTopic)
	log.Printf("Subscribing to %v", subscribeTopic)
	if token := client.Subscribe(subscribeTopic, 1, getProgressBarListener(ecowatchId)); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	select {}
}
