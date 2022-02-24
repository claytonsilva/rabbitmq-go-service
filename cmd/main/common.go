package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string, hasFatal bool) {
	if err != nil {
		if hasFatal {
			log.Fatalf("%s: %s", msg, err)
		} else {
			log.Printf("%s: %s", msg, err)
		}
	}
}

func closeConn(connection *amqp.Connection) {
	err := connection.Close()
	failOnError(err, "Error closing connection", true)
}

func closeCh(channel *amqp.Channel) {
	err := channel.Close()
	failOnError(err, "Error closing channel", true)
}
