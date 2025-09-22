package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/dbahlbeck/go2ecowatch/internal"
	"github.com/spf13/viper"
)

var progressBarTopic = "go2ecowatch/inner/progressbar"

func init() {
	viper.SetEnvPrefix("G2E")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/go2ecowatch")
	viper.MustBindEnv("HOST")
	viper.MustBindEnv("PORT")
	viper.MustBindEnv("USER")
	viper.MustBindEnv("PASSWORD")
	viper.MustBindEnv("ECOWATCH_ID")

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

	f := internal.DefaultMqttClientFactory{}
	client := f.CreateClient(username, password, connectionString)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	subscribeTopic := fmt.Sprintf("%v", progressBarTopic)
	log.Printf("Subscribing to %v", subscribeTopic)
	if token := client.Subscribe(subscribeTopic, 1, internal.GetProgressBarListener(ecowatchId)); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	select {}
}
