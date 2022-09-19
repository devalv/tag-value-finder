package errors

const (
	ErrorToken       = "error occurred during tokenization"
	RmqConnectError  = "failed to connect to RabbitMQ"
	RmqChanOpenError = "failed to open channel to RabbitMQ"
	RmqInQueueError  = "failed to declare an incoming queue"
	RmqOutQueueError = "failed to declare an outgoing queue"
	RmqPublishError  = "failed to publish a message to RabbitMQ"
)
