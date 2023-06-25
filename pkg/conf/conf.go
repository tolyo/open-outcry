package conf

import (
	"bytes"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	// Loads yaml config files
	_ "embed"
)

//go:embed dev.env
var dev string

// EnvName type constraint for Environment types
type EnvName string

// Environment types
const (
	DEV  EnvName = "DEV"
	PROD EnvName = "PROD"
	TEST EnvName = "TEST"
)

// Environment types to yaml mapper
var mapper = map[EnvName]string{
	DEV: dev,
}

// configuration - structure that contains configuration information from config variables
type configuration struct {
	DBDsn      string
	UpdateFees bool
}

var config *configuration
var envName EnvName

// LoadConfig - loads configurations from config variables into Environment struct
func LoadConfig(conf string) (*configuration, error) {
	envName = EnvName(conf)
	switch envName {
	case DEV:
		viper.SetConfigType("env")

		err := viper.ReadConfig(bytes.NewBuffer([]byte(mapper[envName])))
		if err != nil {
			log.Fatal("Invalid config")
		}
	case PROD:
		viper.AutomaticEnv()

	default:
		log.Fatal("Environment variable not supplied")
	}

	config = &configuration{
		DBDsn: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			viper.GetString("POSTGRES_HOST"),
			viper.GetString("POSTGRES_USER"),
			viper.GetString("POSTGRES_PASSWORD"),
			viper.GetString("POSTGRES_DB"),
			viper.GetString("POSTGRES_PORT"),
		),
		UpdateFees: viper.GetBool("UPDATE_FEES"),
	}

	return config, nil
}

// Get config
func Get() *configuration {
	if config == nil {
		log.Fatal("Env not initialized")
	}
	return config
}

// LoadTestConfig - Helper for calling in tests
func LoadTestConfig() (*configuration, error) {
	return LoadConfig(string(DEV))
}

// IsDevEnvironment to check if the current env is dev environment
func IsDevEnvironment() bool {
	return envName == DEV
}

// IsTestEnvironment to check if the current env is test environment
func IsTestEnvironment() bool {
	return envName == TEST
}
