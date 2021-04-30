package application

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type StoreConfig struct {
	Dsn string
}

type Config struct {
	Port     string
	LogLevel string
	Store    StoreConfig
}

func NewConfig() *Config {
	return &Config{
		Port:     "8080",
		LogLevel: "debug",
	}
}

func (c *Config) SetUp(filename string) error {
	v := viper.New()

	v.SetDefault("LOG_LEVEL", "debug")

	if err := c.configureFromFile(filename, v); err != nil {
		if err := c.configureFromEnvironment(v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) configureFromFile(filename string, v *viper.Viper) error {
	v.SetConfigName(filename)
	v.SetConfigType("toml")
	v.AddConfigPath("configs")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}

func (c *Config) configureFromEnvironment(v *viper.Viper) error {
	v.AutomaticEnv()

	var ok bool

	rawPort := v.Get("PORT")
	if c.Port, ok = rawPort.(string); !ok {
		c.Port = "8081"
	}

	rawLogLevel := v.Get("LOG_LEVEL")
	c.LogLevel, _ = rawLogLevel.(string)

	rawDbDsn := v.Get("DB_DSN")
	if c.Store.Dsn, ok = rawDbDsn.(string); !ok {
		return errors.New("не заполнен дсн от бд")
	}

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf(`
Config
----------------------
Port: %s,
LogLevel: %s,
Store {
	Dsn: %s, 
}`,
		c.Port,
		c.LogLevel,
		c.Store.Dsn,
	)
}
