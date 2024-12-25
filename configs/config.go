package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBName         string `mapstructure:"DB_DATABASE"`
	DBUsername     string `mapstructure:"DB_USERNAME"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBPort         string `mapstructure:"DB_PORT"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	Cloudinary_url string `mapstructure:"CLOUDINARY_URL"`
	SlipOk_url     string `mapstructure:"SLIPOK_URL"`
	SlipOk_Secret  string `mapstructure:"SLIPOK_SECRET"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("❌ Error reading config file", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalln("❌ Unable to decode into struct", err)
	}

	return config
}
