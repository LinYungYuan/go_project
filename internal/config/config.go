package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	Environment   string // 新增的字段，用於存儲當前環境
}

func Load() (*Config, error) {
	// 設置默認值
	viper.SetDefault("SERVER_ADDRESS", ":8080")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("ENVIRONMENT", "development") // 默認環境

	// 讀取環境變量
	viper.AutomaticEnv()

	// 將環境變量名轉換為小寫
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 獲取當前環境
	env := viper.GetString("ENVIRONMENT")

	// 設置配置文件的名稱和類型
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")

	// 添加配置文件的搜索路徑
	viper.AddConfigPath("./configs") // 項目根目錄的 config 文件夾
	viper.AddConfigPath(".")         // 當前工作目錄

	// 讀取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到錯誤；如果需要可以忽略
			fmt.Printf("No config file found for environment: %s\n", env)
		} else {
			// 配置文件被找到，但產生了另外的錯誤
			return nil, err
		}
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.Environment = env

	return &cfg, nil
}
