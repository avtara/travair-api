package config

import (
	"fmt"
	"github.com/streadway/amqp"
)

func SetupAMPQConnection()  *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic("Failed Initializing Broker Connection")
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}
	return ch
}