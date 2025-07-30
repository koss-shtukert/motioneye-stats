package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment              string `mapstructure:"APP_ENV"`
	LogLevel                 string `mapstructure:"LOG_LEVEL"`
	CronRunDiskUsageJob      bool   `mapstructure:"CRON_RUN_DISK_USAGE_JOB"`
	CronDiskUsageJobPath     string `mapstructure:"CRON_DISK_USAGE_JOB_PATH"`
	CronDiskUsageJobInterval string `mapstructure:"CRON_DISK_USAGE_JOB_INTERVAL"`
	TgBotApiKey              string `mapstructure:"TGBOT_API_KEY"`
	TgBotChatId              string `mapstructure:"TGBOT_CHAT_ID"`
}

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("CRON_RUN_DISK_USAGE_JOB", false)
	viper.SetDefault("CRON_DISK_USAGE_JOB_PATH", "")
	viper.SetDefault("CRON_DISK_USAGE_JOB_INTERVAL", "")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if cfg.CronDiskUsageJobPath == "" {
		cfg.CronRunDiskUsageJob = false
	}

	required := map[string]string{
		"APP_ENV":       cfg.Environment,
		"LOG_LEVEL":     cfg.LogLevel,
		"TGBOT_API_KEY": cfg.TgBotApiKey,
		"TGBOT_CHAT_ID": cfg.TgBotChatId,
	}

	for key, value := range required {
		if value == "" {
			return nil, fmt.Errorf("required variable %s is empty", key)
		}
	}

	return &cfg, nil
}
