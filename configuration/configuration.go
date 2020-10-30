package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Configuration struct {
	dbConfig *DatabaseConfig
}

func GetConfig() (*Configuration, error) {
	fmt.Println("Loading the env")
	var err error
	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbconfig := DatabaseConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
	}

	configuration := &Configuration{dbConfig: &dbconfig}

	return configuration, nil
}

type DatabaseConfig struct {
	User, Password, Name, Port, Host string
}

func (receiver *Configuration) DBConfig() *DatabaseConfig {
	return receiver.dbConfig
}
