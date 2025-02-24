package configs

import (
	"time"

	"github.com/spf13/viper"
	"gohub/internal/libs/logger"
)

const (
	ProductionEnv = "production"

	DatabaseTimeout    = time.Second * 5
	ProductCachingTime = time.Minute * 1
)

var AuthIgnoreMethods = []string{
	"/user.UserService/Login",
	"/user.UserService/Register",
}

type Config struct {
	Environment            string `mapstructure:"ENVIRONMENT"`
	HttpPort               int    `mapstructure:"HTTP_PORT"`
	GrpcPort               int    `mapstructure:"GRPC_PORT"`
	SocketPort             int    `mapstructure:"SOCKET_PORT"`
	AuthSecret             string `mapstructure:"AUTH_SECRET"`
	DatabaseURI            string `mapstructure:"DATABASE_URI"`
	RedisURI               string `mapstructure:"REDIS_URI"`
	RedisPassword          string `mapstructure:"REDIS_PASSWORD"`
	RedisDB                int    `mapstructure:"REDIS_DB"`
	GoogleClientID         string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret     string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	CloudinaryCloudName    string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryApiKey       string `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryApiSecret    string `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudinaryUploadFolder string `mapstructure:"CLOUDINARY_UPLOAD_FOLDER"`
	UrlCloudinary          string `mapstructure:"URL_CLOUDINARY"`
	StripeSecretKey        string `mapstructure:"STRIPE_SECRET_KEY"`
}

var (
	cfg Config
)

func LoadConfig(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Error on load configuration file, error: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		logger.Fatalf("Error on parsing configuration file, error: %v", err)
	}

	return &cfg
}

func GetConfig() *Config {
	return &cfg
}
