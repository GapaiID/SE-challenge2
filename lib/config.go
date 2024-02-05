package lib

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath = "./config/config.yaml"

var defaultConfig = Config{
	Name: "miblog",
	Http: &HttpConfig{
		Host: "127.0.0.1",
		Port: 8000,
	},
	Database: &DatabaseConfig{
		Parameters:   "",
		MigrationDir: "migrations",
	},
	JWT: &JWTConfig{
		TokenLifeTime: 1440,
	},
	Mail: &MailConfig{
		Enable:    false,
		Host:      "smtp.gmail.com",
		Port:      587,
		User:      "user",
		Password:  "password",
		UseTLS:    true,
		FromEmail: "NoReply <norepy@example.com>",
	},
	Swagger: &SwaggerConfig{
		Title:       "Miblog API",
		Description: "Miblog Endpoints",
		Version:     "1.0",
	},
	Redis: &RedisConfig{Host: "127.0.0.1", Port: 6379},
}

func NewConfig() Config {
	config := defaultConfig

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Wrap(err, "failed to read config"))
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(errors.Wrap(err, "failed to marshal config"))
	}

	return config
}

type Config struct {
	Name      string `mapstructure:"Name"`
	SecretKey string `mapstructure:"SecretKey"`

	Http     *HttpConfig     `mapstructure:"Http"`
	Database *DatabaseConfig `mapstructure:"database"`
	JWT      *JWTConfig      `mapstructure:"jwt"`
	Mail     *MailConfig     `mapstructure:"mail"`
	Swagger  *SwaggerConfig  `mapstructure:"swagger"`
	Redis    *RedisConfig    `mapstructure:"redis"`
}

type RedisConfig struct {
	Host      string `mapstructure:"Host"`
	Port      int    `mapstructure:"Port"`
	KeyPrefix string `mapstructure:"KeyPrefix"`
}

func (a *RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

type SwaggerConfig struct {
	Title       string `mapstructure:"title"`
	Description string `maptsructure:"description"`
	Version     string `mapstructure:"version"`
}

type MailConfig struct {
	Enable    bool   `mapstructure:"Enable"`
	Host      string `mapstructure:"Host"`
	Port      int    `mapstructure:"Port"`
	User      string `mapstructure:"User"`
	Password  string `mapstructure:"Password"`
	UseTLS    bool   `mapstructure:"UseTLS"`
	FromEmail string `mapstructure:"FromEmail"`
}

type JWTConfig struct {
	TokenLifeTime int `mapstructure:"TokenLifeTime"`
}

type HttpConfig struct {
	Host string `mapstructure:"Host" validate:"ipv4"`
	Port int    `mapstructure:"Port" validate:"gte=1,lte=65535"`
}

func (a *HttpConfig) ListenAddr() string {
	if err := validator.New().Struct(a); err != nil {
		return "0.0.0.0:5100"
	}

	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

type DatabaseConfig struct {
	Engine       string `mapstructure:"Engine"`
	Name         string `mapstructure:"Name"`
	Host         string `mapstructure:"Host"`
	Port         int    `mapstructure:"Port"`
	Username     string `mapstructure:"Username"`
	Password     string `mapstructure:"Password"`
	Parameters   string `mapstructure:"Parameters"`
	MigrationDir string `mapstructure:"MigrationDir"`
}

func (a *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?%s", a.Engine, a.Username, a.Password, a.Host, a.Port, a.Name, a.Parameters)
}

func SetConfigPath(path string) {
	configPath = path
}
