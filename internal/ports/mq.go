package ports

type MQ interface {
	Disconnect() error
	PublishResponse(msg string) error
	LaunchConsumer() error
}
