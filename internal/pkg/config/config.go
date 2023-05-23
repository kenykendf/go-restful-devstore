package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver                string        `mapstructure:"DB_Driver"`
	DBConnection            string        `mapstructure:"DB_Connection"`
	ServerPort              string        `mapstructure:"Server_Port"`
	LogLevel                string        `mapstructure:"LOG_LEVEL"`
	AccessTokenKey          string        `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey         string        `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration    time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	CloudinaryName          string        `mapstructure:"CLOUDINARY_NAME"`
	CloudinaryAPIKey        string        `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryAPISecret     string        `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudinaryDir           string        `mapstructure:"CLOUDINARY_DIR"`
	PaginateDefaultPage     int           `mapstructure:"PAGINATE_DEFAULT_PAGE"`
	PaginateDefaultPageSize int           `mapstructure:"PAGINATE_DEFAULT_PAGE_SIZE"`
}

// nolint
func LoadConfig(fileConfigPath string) (Config, error) {
	var config Config

	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	viper.Unmarshal(&config)
	return config, nil
}
