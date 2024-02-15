package config

type Config struct {
	MongoURI       string
	DatabaseName   string
	CollectionName string
}

func NewConfig() *Config {
	return &Config{
		MongoURI:       "mongodb://localhost:27017",
		DatabaseName:   "kafkonnector",
		CollectionName: "connectors",
	}
}
