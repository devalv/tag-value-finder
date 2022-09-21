package rmq

import (
	"context"
	"strings"
	"time"

	"github.com/devalv/tag-value-finder/internal/domain/crawler"
	"github.com/devalv/tag-value-finder/internal/domain/errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type YawmRmq struct {
	c            *amqp.Connection
	ch           *amqp.Channel
	inQueryName  string
	outQueryName string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal().Msgf(msg, err)
	}
}

func declareQueue(ch *amqp.Channel, queryName string) (string, error) {
	mq, err := ch.QueueDeclare(
		queryName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return "", err
	}
	return mq.Name, nil
}

func NewYawm(ctx context.Context, rmqConnURI, inQueryName, outQueryName string) (*YawmRmq, error) {
	log.Debug().Msgf("Trying to connect to %s", rmqConnURI)
	conn, err := amqp.Dial(rmqConnURI)
	failOnError(err, errors.RMQConnectError)

	ch, err := conn.Channel()
	failOnError(err, errors.RMQChanOpenError)

	imqName, err := declareQueue(ch, inQueryName)
	failOnError(err, errors.RMQInQueueError)

	omqName, err := declareQueue(ch, outQueryName)
	failOnError(err, errors.RMQOutQueueError)

	log.Debug().Msgf("Connected to %s", rmqConnURI)
	return &YawmRmq{c: conn, ch: ch, inQueryName: imqName, outQueryName: omqName}, nil
}

func (y *YawmRmq) Disconnect() error {
	_ = y.ch.Close()
	_ = y.c.Close()
	return nil
}

func (y *YawmRmq) PublishResponse(msg string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := y.ch.PublishWithContext(ctx,
		"",             // exchange
		y.outQueryName, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	log.Debug().Msgf("Pushed a message: %s", msg)
	return err
}

func (y *YawmRmq) LaunchConsumer() error {
	messages, _ := y.ch.Consume(
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
		for m := range messages {
			log.Debug().Msgf("Received a message: %s", m.Body)
			tagValue := crawler.GetH1(string(m.Body))
			err := y.PublishResponse(strings.TrimSpace(tagValue))
			failOnError(err, errors.RMQPublishError)
		}
	}()

	log.Debug().Msg(" [*] Waiting for messages.")
	<-forever

	return nil
}
