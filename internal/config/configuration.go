/*
    Main Configuration
*/
package config

import (
	"github.com/reshimahendra/gin-starter/pkg/logger"
	"github.com/spf13/viper"
)


var Config *Configuration

type Configuration struct {
   Server   ServerConfiguration
   Database DatabaseConfiguration
}

func GetConfig() *Configuration {
    return Config
}

// Initiate Configuration by calling our config file.
// Supported configuration format ['*.yaml', '*.toml', '*.json', etc]
func Setup() error {
    var config *Configuration

    // Place the config file at root folder with name '.config.yaml'
    viper.SetConfigName(".config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    // try to read the config 
    if err := viper.ReadInConfig(); err != nil {
        logger.Errorf("Error reading the config file, %v", err)
        return err
    }

    if err := viper.Unmarshal(&config); err != nil {
        logger.Errorf("Unable to decode into struct, %v", err)
        return err
    }

    Config = config
    return nil
}
