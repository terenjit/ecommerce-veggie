package config

import "github.com/spf13/viper"

type App struct {
	AppPort      string `json:"app_port"`
	AppEnv       string `json:"app_env"`
	JWTSecretKey string `json:"jwt_secret_key"`
	JWTIssuer    string `json:"jwt_issuer"`
}

type PsqlDB struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	User           string `json:"user"`
	Password       string `json:"password"`
	DBName         string `json:"db_name"`
	DBMaxIdleConns int    `json:"db_max_idle"`
	DBMaxOpenConns int    `json:"db_max_open"`
}

type Config struct {
	App    App    `json:"app"`
	PsqlDB PsqlDB `json:"psql_db"`
}

var ConfigInstance Config

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort:      viper.GetString("APP_PORT"),
			AppEnv:       viper.GetString("APP_ENV"),
			JWTSecretKey: viper.GetString("JWT_SECRET_KEY"),
			JWTIssuer:    viper.GetString("JWT_ISSUER"),
		},
		PsqlDB: PsqlDB{
			Host:           viper.GetString("DATABASE_HOST"),
			Port:           viper.GetInt("DATABASE_PORT"),
			User:           viper.GetString("DATABASE_USER"),
			Password:       viper.GetString("DATABASE_PASSWORD"),
			DBName:         viper.GetString("DATABASE_NAME"),
			DBMaxIdleConns: viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
			DBMaxOpenConns: viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
		},
	}
}
