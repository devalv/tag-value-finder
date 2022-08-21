package main

//import "fmt"

//
//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Panicf("%s: %s", msg, err)
//	}
//}
//
//func main() {
//	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
//	failOnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//
//	ch, err := conn.Channel()
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	//q, err := ch.QueueDeclare(
//	//	"hello", // name
//	//	false,   // durable
//	//	false,   // delete when unused
//	//	false,   // exclusive
//	//	false,   // no-wait
//	//	nil,     // arguments
//	//)
//	//failOnError(err, "Failed to declare a queue")
//
//	body := "Hello Yoba!"
//
//	for i := 1; i <= 10000; i++ {
//		err = ch.Publish(
//			"test-e-multi",   // exchange
//			"test-q-request", // routing key
//			false,            // mandatory
//			false,            // immediate
//			amqp.Publishing{
//				ContentType: "text/plain",
//				Body:        []byte(body),
//			})
//		failOnError(err, "Failed to publish a message")
//		log.Printf(" [x] Sent %s\n", body)
//		//time.Sleep(1 * time.Second)
//	}
//
//	// Read messages
//	msgs, err := ch.Consume(
//		"multi-q-1", // queue
//		"",          // consumer
//		true,        // auto-ack
//		false,       // exclusive
//		false,       // no-local
//		false,       // no-wait
//		nil,         // args
//	)
//	failOnError(err, "Failed to register a consumer")
//
//	var forever chan struct{}
//
//	go func() {
//		for d := range msgs {
//			log.Printf("Received a message: %s", d.Body)
//		}
//	}()
//
//	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
//	<-forever
//}
