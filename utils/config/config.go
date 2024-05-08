package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sync"
)

var once sync.Once

func InitConfigByViper() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../")
		viper.AddConfigPath("../../")

		zap.S().Infof("reading config file: %s", viper.ConfigFileUsed())
		if err := viper.ReadInConfig(); err != nil {
			zap.S().Fatal("error reading config file", zap.Error(err))
		}
	},
	)
}
