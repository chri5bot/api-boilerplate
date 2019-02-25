package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// DBConfiguration database configuration
type DBConfiguration struct {
	URL string
}

// Configuration service configuration
type Configuration struct {
	// Port int `default:"4000"`
	Port int
	DB   *DBConfiguration
}

func loadEnviroment(filename string) error {
	var err error

	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()

		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

// LoadConfig loads the configuration from env file
func LoadConfig(filename string) (*Configuration, error) {
	if err := loadEnviroment(filename); err != nil {
		return nil, err
	}

	config := new(Configuration)

	if err := envconfig.Process("API", config); err != nil {
		return nil, err
	}

	return config, nil
}
