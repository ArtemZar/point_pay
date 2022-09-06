package config

import "github.com/spf13/viper"

type Config struct {
	SrvConfig GRPCConfig
	DB        DataBase
}

type GRPCConfig struct {
	Network string
	Addr    string
}

type DataBase struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	AuthDB   string
}

func New() (*Config, error) {
	conf := &Config{}

	viper.SetDefault("grpc_server.addr", "localhost:8081")
	viper.SetDefault("grpc_server.network", "tcp")
	conf.SrvConfig.Addr = viper.GetString("grpc_server.addr")
	conf.SrvConfig.Network = viper.GetString("grpc_server.network")

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "27017")
	viper.SetDefault("database.username", "")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.database", "account-service")
	viper.SetDefault("database.authDB", "")
	conf.DB.Host = viper.GetString("database.host")
	conf.DB.Port = viper.GetString("database.port")
	conf.DB.Username = viper.GetString("database.username")
	conf.DB.Password = viper.GetString("database.password")
	conf.DB.Database = viper.GetString("database.database")
	conf.DB.AuthDB = viper.GetString("database.authDB")
	return conf, nil
}
