package models

type Config struct {
	HealthCheckAddr string `env:"HEALTH_CHECK_ADDR" env-default:":3333"`
	MQConnURI       string `env:"MQ_CONN_URI" env-default:"amqp://guest:guest@localhost:5672/"`
	InQueryName     string `env:"IN_QUERY_NAME" env-default:"biba"`
	OutQueryName    string `env:"OUT_QUERY_NAME" env-default:"boba"`
}
