package app

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
	"github.com/tomochain/backend-matching-engine/utils"
)

// Config stores the application-wide configurations
var Config appConfig
var logger = utils.Logger

type appConfig struct {

	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// the elasticsearch url. required.
	ElasticsearchURL string `mapstructure:"elasticsearch_url"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.ElasticsearchURL, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "RESTFUL_" in their names are also read automatically.
func LoadConfig(configPath string, env string) error {
	v := viper.New()

	if env != "" {
		v.SetConfigName("config." + env)
	}

	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)

	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	v.SetEnvPrefix("api")
	v.AutomaticEnv()

	err = v.Unmarshal(&Config)
	if err != nil {
		return err
	}

	// log information
	logger.Infof("Server port: %v", Config.ServerPort)
	logger.Infof("Elasticsearch url: %v", Config.ElasticsearchURL)
	return Config.Validate()
}
