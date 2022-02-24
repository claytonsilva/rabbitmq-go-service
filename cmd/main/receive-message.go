package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func receive(ampqQueryString string, queueName string, interval int) {
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

	err = ch.Qos(
		3,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS", true)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer", true)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s with content: <%s>", d.MessageId, d.Body)
			time.Sleep(time.Duration(interval) * time.Millisecond)
			log.Printf("Done")
			err := d.Ack(false)
			failOnError(err, "fail to ack message", true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
