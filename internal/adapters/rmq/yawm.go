package rmq

import (
	"context"
	"tag-value-finder/internal/domain/crawler"
	"tag-value-finder/internal/domain/errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type YawmRmq struct {
	c           *amqp.Connection
	ch          *amqp.Channel
	inQueryName string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panic().Msgf("%s: %s", msg, err)
	}
}

func NewYawm(ctx context.Context, rmqConnURI, inQueryName string) (*YawmRmq, error) {
	log.Debug().Msgf("Trying to connect to %s", rmqConnURI)
	conn, err := amqp.Dial(rmqConnURI)
	failOnError(err, errors.RmqConnectError)

	ch, err := conn.Channel()
	failOnError(err, errors.RmqChanOpenError)

	mq, err := ch.QueueDeclare(
		inQueryName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Debug().Msgf("Connected to %s", rmqConnURI)
	return &YawmRmq{c: conn, ch: ch, inQueryName: mq.Name}, nil
}

func (y *YawmRmq) Disconnect() error {
	y.ch.Close()
	y.c.Close()
	return nil
}

func (y *YawmRmq) LaunchConsumer() error {
	msgs, _ := y.ch.Consume(
		y.inQueryName,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args)
	)

	var forever chan struct{}

	go func() {
		for m := range msgs {
			log.Debug().Msgf("Received a message: %s", m.Body)
			tagValue := crawler.GetH1(string(m.Body))
			log.Debug().Msgf("Tag value is: %s", tagValue)
			// TODO: post response to Producer
		}
	}()

	log.Debug().Msgf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
