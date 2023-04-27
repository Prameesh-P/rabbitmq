package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")  // connecting to localhost amqp server
	if err!=nil{
			checkError(err)
	}
	defer conn.Close()

	ch,err:=conn.Channel()
	if err!=nil{
		checkError(err)
	}
	defer ch.Close()

	msgs,err:=ch.Consume(  //Consuming from queue 
		"Queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if  err!=nil {
		checkError(err)
	}

	forever:=make(chan bool)

	go func(){
		for d:=range msgs{
			fmt.Printf("Recieved Msg >>> %s\n",d.Body)
		}
	}()
	fmt.Println("Successfully connected to the instance ")
	fmt.Println("Waiting for the messages>>>")
	<-forever 
}

func checkError(err error){ // error handling function
	log.Fatalf("Error while get %v",err)
}