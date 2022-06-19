package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)



func main() {

	// Writing --- Publisher

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
		panic(err)
	}

	defer ch.Close()

	// Creating queue for messages 

	q, err := ch.QueueDeclare(
		"Test Queue",
		false,
		false,
		false,
		false,
		nil, 
	)

	if err != nil {
		fmt.Println("Failed on defining queue")
		panic(err)
	} 
	
	// Information about queue

	fmt.Println(q) 

	// Creating Publisher 

	err = ch.Publish(
		"",
		"Test Queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello from publisher!"),
		},
	)

	if err != nil {
		fmt.Println("Failed on publishing message :(")
	}

	fmt.Println("Succesfully connecting on Rabbit-MQ")
}