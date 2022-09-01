package config

type ServerConfig struct {
	BindAddr string
}

func NewCofig() *ServerConfig {
	return &ServerConfig{
		BindAddr: ":8080",
	}
}
