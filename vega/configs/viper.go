package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var cfgFile string

var Config Configuration

type Configuration struct {
	Vega `mapstructure:"vega" json:"vega"`
}

type Vega struct {
	API       `mapstructure:"api" json:"api"`
	Semaphore `mapstructure:"semaphore" json:"semaphore"`
	Awx       `mapstructure:"awx" json:"awx"`
}

type API struct {
	Port        int    `mapstructure:"port" json:"port"`
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	TokenSecret string `mapstructure:"token_secret" json:"token_secret"`
}

type Semaphore struct {
	URL      string `mapstructure:"url" json:"url"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

type Awx struct {
	URL      string `mapstructure:"url" json:"url"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

var (
	defaults = map[string]interface{}{
		"debug.enabled": false,
		"vega.api.port": 8000,
	}
	envPrefix   = "CRUCIBLE"
	configName  = "config"
	configType  = "yaml"
	configPaths = []string{
		".",
		fmt.Sprintf("%s/.crucible", getUserHomePath()),
	}
)

var allowedEnvVarKeys = []string{
	"vega.semaphore.url",
	"vega.semaphore.username",
	"vega.semaphore.password",
	"debug.enabled",
}

// Bootstrap reads in config file and ENV variables if set.
func Bootstrap() {
	err := godotenv.Load(".env")

	if err != nil {
		if viper.GetString("debug.enabled") == "true" {
			log.Println("Error loading .env file")
		}
	}

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		for _, p := range configPaths {
			viper.AddConfigPath(p)
		}
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
		err := viper.ReadInConfig()
		if err != nil {
			if viper.GetString("debug.enabled") == "true" {
				fmt.Println(err)
			}
		}
	}
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(envPrefix)

	for _, key := range allowedEnvVarKeys {
		err := viper.BindEnv(key)
		if err != nil {
			fmt.Println(err)
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("could not decode config into struct: %v", err)
	}
}

func getUserHomePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}
