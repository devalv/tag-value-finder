package errors

const (
	ErrorToken            = "error occurred during tokenization"
	RMQConnectError       = "failed to connect to RabbitMQ"
	RMQChanOpenError      = "failed to open channel to RabbitMQ"
	RMQInQueueError       = "failed to declare an incoming queue"
	RMQOutQueueError      = "failed to declare an outgoing queue"
	RMQPublishError       = "failed to publish a message to RabbitMQ"
	HTTPHealthListenError = "failed to listen health check port:`%e`"
	ConfigError           = "failed to read config:`%e`"
)
