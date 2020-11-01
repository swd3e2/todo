package application

import (
	"fmt"

	"github.com/spf13/viper"
)

type StoreConfig struct {
	User     string
	Pwd      string
	Dsn      string
	Port     string
	Database string
}

type Config struct {
	Port     string
	LogLevel string
	Name     string
	PathMap  string `mapstructure:"path_map"`
	Store    StoreConfig
}

func NewConfig() *Config {
	return &Config{
		Port:     "8080",
		LogLevel: "debug",
	}
}

func (this *Config) SetUp(filename string) error {
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&this); err != nil {
		return err
	}

	return nil
}

func (this *Config) String() string {
	return fmt.Sprintf(`Port: %s
	Store {User: %s, Pwd: %s, Dsn: %s, Port: %s, Database: %s}`,
		this.Port,
		this.Store.User,
		this.Store.Pwd,
		this.Store.Dsn,
		this.Store.Port,
		this.Store.Database,
	)
}
