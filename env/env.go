package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerAddress   string
	LibraryFolder   string
	KobomaticFolder string
}

func getVariable(variable string) (string, error) {
	value := os.Getenv(variable)
	if value == "" {
		return "", fmt.Errorf("environment variable %s is not set", variable)
	}
	return value, nil
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("no .env file found, expecting inline environment variables")
	}

	serverAddress, err := getVariable("SERVER_ADDRESS")
	if err != nil {
		return nil, err
	}

	libraryFolder, err := getVariable("LIBRARY_FOLDER")
	if err != nil {
		return nil, err
	}

	kobomaticFolder, err := getVariable("KOBOMATIC_FOLDER")
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerAddress:   serverAddress,
		LibraryFolder:   libraryFolder,
		KobomaticFolder: kobomaticFolder,
	}, nil
}
