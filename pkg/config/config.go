// TODO: add a func to get config by key and unmarshal into a passed in structure.

package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type (
	// Service is the structure for the service information configuration.
	Service struct {
		Group   string
		Name    string
		Version string
	}

	// DB is the structure for the main database configuration.
	DB struct {
		Type     string
		Host     string
		Port     int
		User     string
		Password string
		Database string
		Log      bool
	}

	// Management is the structure for the management http endpoint configuration.
	Management struct {
		Endpoint struct {
			Port            int
			BaseRoutingPath string
		}
		Health struct {
			Path string
			Full bool
		}
	}

	// Log is the structure for the logger configuration.
	// If not present, the Machinery will use a default logger provided
	// by the "gm-log" package.
	Log struct {
		Path     string
		Filename string
		Console  struct {
			Enabled       bool
			DisableColors bool
			Colors        bool
		}
		Level           string
		JSON            bool
		MaxSize         int
		MaxBackups      int
		MaxAge          int
		Compress        bool
		LocalTime       bool
		TimestampFormat string
		FullTimestamp   bool
		ForceFormatting bool
	}

	// API is the structure for the Http API server and app configuration.
	API struct {
		Endpoint struct {
			Port            int
			BaseRoutingPath string
		}
		Security struct {
			Enabled bool
			Jwt     struct {
				Secret     string
				Expiration struct {
					Enabled bool
					Minutes int32
				}
			}
		}
	}

	// Ldap configuration
	Ldap struct {
		Base   string
		Host   string
		Port   int
		UseSSL bool
		Bind   struct {
			DN       string
			Password string
		}
		UserFilter  string
		GroupFilter string
		Attributes  []string
	}

	// Configuration describe the type for the configuration file
	// Log "required" is used for side-effect check: if err != nil => use default logger.
	// Then we go by convention: if Log exists then we assert its mandatory fields exixsts.
	Configuration struct {
		Service    Service
		API        API
		DB         DB
		Management Management
		Log        Log `validate:"required"`
		Ldap       Ldap
	}
)

var instance *Configuration
var once sync.Once

// GetConfiguration returns the Configuration structure singleton instance.
func GetConfiguration() *Configuration {
	once.Do(func() {
		loadConfiguration()
	})

	return instance
}

func loadConfiguration() {
	viper.SetDefault("logPath", "./log")

	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if os.Getenv("ENV") != "" {
		viper.SetConfigName("config-" + os.Getenv("ENV"))
	} else {
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(&instance); err != nil {
		panic(fmt.Errorf("fatal error decoding configuration into struct: %v", err))
	}

}

// Get returns a configuration map by key. Used for custom or gear configurations.
func Get(key string) interface{} {
	// just in case!
	conf := GetConfiguration()
	if conf == nil {
		panic("No configuration at all!")
	}
	return viper.Get(key)
}

// IsSet checks to see if the key has been set in any of the data locations.
// IsSet is case-insensitive for a key.
func IsSet(key string) bool {
	// just in case!
	conf := GetConfiguration()
	if conf == nil {
		panic("No configuration at all!")
	}
	return viper.IsSet(key)
}
