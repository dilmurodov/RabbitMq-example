package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main () {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed on loading environment variables!")
		panic(err)
	}

	url := os.Getenv("RABBIT_MQ")

	fmt.Println(url)

	// connecting Dial rabbit-mq message broker

	conn, err := amqp.Dial(url)

	if err != nil {
		fmt.Println("Failed on connecting on rabbit-mq")
		panic(err)
	}

	defer conn.Close()

	// opening channel

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	// Connecting Comsumer to Channel

	msgs, err := ch.Consume(
		"Test Queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Channel For blockin and forever wait messages from Publisher

	forever := make(chan bool)

	go func () {
		i := 0
		for ms := range msgs {
			fmt.Println("Received message [*] - ", string(ms.Body), i)
			i++	
		}
	}()
	fmt.Println("Consumer Succesfully Registered!")
	<- forever
	fmt.Println("Comsumer Failed!")
}