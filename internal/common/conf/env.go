package conf

import (
	// "fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBCluster      string `mapstructure:"DB_CLUSTER"`
	DBName         string `mapstructure:"DB_NAME"`
	// AccessTokenScecret
}

func NewEnv() *Env {
	env := Env{}
	if os.Getenv("APP_ENV") == "" {
		viper.SetConfigFile(".env.yaml")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Can't find the file .env.yaml: ", err)
		}

		err = viper.Unmarshal(&env)
		if err != nil {
			log.Fatal("Enviroment can't be loaded: ", err)
		}

		if env.AppEnv == "development" {
			log.Println("The app is running in development env")
		}
	} else {
		timeOut, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
		env = Env{
			AppEnv:         os.Getenv("APP_ENV"),
			ServerAddress:  os.Getenv("SERVER_ADDRESS"),
			ContextTimeout: timeOut,
			DBHost:            os.Getenv("DB_HOST"),
			DBPort:            os.Getenv("DB_PORT"),
			DBUser:            os.Getenv("DB_USER"),
			DBPass:            os.Getenv("DB_PASS"),
			DBCluster:         os.Getenv("DB_CLUSTER"),
			DBName:            os.Getenv("DB_NAME"),
			// AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		}
	}

	return &env
}
