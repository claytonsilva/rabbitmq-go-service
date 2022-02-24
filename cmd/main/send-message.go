package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func send(ampqQueryString string, queueName string, interval int) {
	conn, err := amqp.Dial(ampqQueryString)
	failOnError(err, "Failed to connect to RabbitMQ", true)
	defer closeConn(conn)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel", true)
	defer closeCh(ch)

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue", true)

	i := 1

	for {
		body := fmt.Sprintf("Message %d sent", i)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
			})
		failOnError(err, "Failed to publish a message", true)
		log.Printf(" [x] Sent (%s)", body)
		i++
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
