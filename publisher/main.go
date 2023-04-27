package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func CheckErrors(err error) {  // error handling function
	log.Fatalf("error occured !! %v", err) 
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")  // connecting to localhost amqp server

	if err != nil {
		CheckErrors(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMq")

	ch, err := conn.Channel()

	if err != nil {
		CheckErrors(err)
	}
	defer ch.Close()

	q,err:=ch.QueueDeclare(  // queue declearing 
		"Queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err !=nil {
		CheckErrors(err)
	}

	fmt.Println(q)
	err=ch.Publish( // queue publishing 
		"",
		"Queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("HYYY"),
		},
	)
	if err !=nil{
		CheckErrors(err)
	}
	fmt.Println("Succefully published message into the queue!!")

	
}
