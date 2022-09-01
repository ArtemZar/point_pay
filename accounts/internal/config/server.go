package config

type GRPC struct {
	Host string `envconfig:"HOST" required:"true" default:""`
	Port string `envconfig:"GRPC_PORT" required:"true" default:""`
}
