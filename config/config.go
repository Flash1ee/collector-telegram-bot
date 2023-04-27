package config

type ServerConfig struct {
	ServerParams   ServerParams             `toml:"server"`
	DatabaseParams PostgresConnectionParams `toml:"database"`
}

type ServerParams struct {
	StartPort string
}

type PostgresConnectionParams struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func CreateConfigForServer() *ServerConfig {
	return &ServerConfig{}
}
