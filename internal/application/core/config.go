package core

import (
	"fmt"

	"github.com/spf13/viper"
)

// StoreConfig Конфиг подключения к бд.
type StoreConfig struct {
	Host     string
	Port     uint
	User     string
	Password string
	Database string
}

// Config Структура со всеми настройками проекта.
type Config struct {
	Port           string
	LogLevel       string
	MigrationsPath string
	Store          StoreConfig
}

// NewConfig Создание нового конфига с настройками.
func NewConfig() *Config {
	return &Config{
		Port:     "8080",
		LogLevel: "debug",
	}
}

// SetUp Заполнение конфига.
// Сначала пытается прочитать конфиг из file, если не удалось пытается достать настройки из файла пытается достать
// их из енвов.
func (c *Config) SetUp(filename string) error {
	v := viper.New()

	v.SetDefault("LOG_LEVEL", "debug")
	v.SetDefault("MIGRATIONS_PATH", "/migrations")

	if err := c.configureFromFile(filename, v); err != nil {
		if err := c.configureFromEnvironment(v); err != nil {
			return err
		}
	}

	return nil
}

// configureFromFile Получение настроек из файла
func (c *Config) configureFromFile(filename string, v *viper.Viper) error {
	v.SetConfigName(filename)
	v.SetConfigType("toml")
	v.AddConfigPath("configs")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}

// configureFromFile Получение настроек из енвов
func (c *Config) configureFromEnvironment(v *viper.Viper) error {
	v.AutomaticEnv()

	c.Port = v.GetString("PORT")
	c.LogLevel = v.GetString("LOG_LEVEL")
	c.MigrationsPath = v.GetString("MIGRATIONS_PATH")

	c.Store.Host = v.GetString("DB_HOST")
	c.Store.Port = v.GetUint("DB_PORT")
	c.Store.User = v.GetString("DB_USER")
	c.Store.Password = v.GetString("DB_PASSWORD")
	c.Store.Database = v.GetString("DB_DATABASE")

	return nil
}

// String Выводит содержимое конфига
func (c *Config) String() string {
	return fmt.Sprintf(`
Config
----------------------
Port: %s,
LogLevel: %s,
MigrationsPath: %s,
Store {
	Host: %s,
	Port: %d,
	User: %s,
	Password: %s,
	Database: %s
}`,
		c.Port,
		c.LogLevel,
		c.MigrationsPath,
		c.Store.Host,
		c.Store.Port,
		c.Store.User,
		c.Store.Password,
		c.Store.Database,
	)
}
