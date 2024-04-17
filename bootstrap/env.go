package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {	
	AppEnv                 	string `mapstructure:"APP_ENV"`
	Port                   	string `mapstructure:"PORT"`
	ServerAddress          	string `mapstructure:"SERVER_ADDRESS"`
	FrontendAddress        	string `mapstructure:"FRONTEND_ADDRESS"`
	ContextTimeout         	int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 	string `mapstructure:"DB_HOST"`
	DBPort                 	string `mapstructure:"DB_PORT"`
	DBUser                 	string `mapstructure:"DB_USER"`
	DBPass                 	string `mapstructure:"DB_PASS"`
	DBName                 	string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  	int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour 	int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      	string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     	string `mapstructure:"REFRESH_TOKEN_SECRET"`
	GoogleClientID         	string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret     	string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	VKClientID             	string `mapstructure:"VK_CLIENT_ID"`
	VKClientSecret         	string `mapstructure:"VK_CLIENT_SECRET"`
	TinkoffTerminalKey     	string `mapstructure:"TINKOFF_TERMINAL_KEY"`
	TinkoffTerminalPassword	string `mapstructure:"TINKOFF_TERMINAL_PASSWORD"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
