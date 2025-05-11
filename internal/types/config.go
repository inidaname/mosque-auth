package types

import "time"

type Config struct {
	Server struct {
		GRPCPort string `yaml:"grpc_port"`
		HTTPPort string `yaml:"http_port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Schema   string `yaml:"schema"`
		SSLMode  string `yaml:"sslmode"`
		Url      string `yaml:"url"`
	} `yaml:"database"`
	DBurl  string
	APIurl string
	Auth   AuthConfig
}

type AuthConfig struct {
	Basic BasicConfig
	Token TokenConfig
}

type BasicConfig struct {
	User string
	Pass string
}

type TokenConfig struct {
	Secret string
	Exp    time.Duration
	Iss    string
}
