package config

import "github.com/spf13/viper"

type Config struct {
	SrvConf        HttpServer
	GrpcClientConf GrpcClient
}

type HttpServer struct {
	BindAddr string
}

type GrpcClient struct {
	Target string
}

func New() (*Config, error) {
	conf := &Config{}
	viper.SetDefault("http_server.addr", ":8080")
	conf.SrvConf.BindAddr = viper.GetString("http_server.addr")

	viper.SetDefault("grpc_client.target", "localhost:8081")
	conf.GrpcClientConf.Target = viper.GetString("grpc_client.target")

	return conf, nil
}
