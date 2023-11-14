package conf

import (
	// "fmt"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	PortServer     string `mapstructure:"PORT_SERVER"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBCluster      string `mapstructure:"DB_CLUSTER"`
	DBName         string `mapstructure:"DB_NAME"`
	// AccessTokenScecret
	PublishDataUrl string `mapstructure:"PUBLISH_DATA_URL"`
}

func NewEnv() *Env {
	env := Env{}

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

	return &env
}
