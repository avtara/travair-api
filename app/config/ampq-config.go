package config

import (
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"os"
)

func SetupAMPQConnection()  *amqp.Channel {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	uri := os.Getenv("AMQP_URI")
	conn, err := amqp.Dial(uri)
	if err != nil {
		panic("Failed Initializing Broker Connection")
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}