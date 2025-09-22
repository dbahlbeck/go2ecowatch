package internal

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClientFactory interface {
	CreateClient(string, string, string) mqtt.Client
}

type DefaultMqttClientFactory struct{}

func (f *DefaultMqttClientFactory) CreateClient(username string, password string, connectionString string) mqtt.Client {
	opts := mqtt.NewClientOptions().
		SetUsername(username).
		SetPassword(password).
		AddBroker(connectionString).
		SetAutoReconnect(true)
	opts.OnConnect = connectHandler
	return mqtt.NewClient(opts)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}
