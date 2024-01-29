package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Cannnot connect to rabbitmq server")
	defer conn.Close()

	ch, ch_err := conn.Channel()
	failOnError(ch_err, "Cannnot connect to an opned channel")

	q, q_err := ch.QueueDeclare("MessageQueue", true, false, false, false, nil)
	failOnError(q_err, "Failed to declare a queue")

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := "Hello world from go server"
	publish_err := ch.PublishWithContext(context, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

	failOnError(publish_err, "Failed to publish")
	log.Printf(" [x] Sent %s\n", message)
}
