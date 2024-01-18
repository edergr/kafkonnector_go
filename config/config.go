package config

type Config struct {
	MongoURI string
}

func NewConfig() *Config {
	return &Config{
		MongoURI: "mongodb://localhost:27017",
	}
}
