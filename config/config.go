package config

type ServerConfig struct {
	DatabaseParams PostgresConnectionParams `toml:"database"`
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
