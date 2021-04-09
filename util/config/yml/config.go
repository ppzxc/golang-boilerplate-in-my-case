package yml

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-arch/util/config/logger"
	"go-arch/util/err"
	"go.uber.org/zap"
)

func Read(configFilePath string, logLevel string, fileName string, useLogFile bool, dsn string) (*Config, error) {
	if len(configFilePath) <= 0 {
		return nil, err.ConfigFilePathIsInvalid
	}

	viper.SetConfigFile(configFilePath)
	if readError := viper.ReadInConfig(); readError != nil {
		return nil, readError
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.L().Info("config file changed",
			zap.String("file name", e.Name))
		if readError := logger.Init(logLevel, fileName, useLogFile, dsn); readError != nil {
			zap.L().Error("viper config change method is error occurred", zap.Error(readError))
		}
	})

	viper.WatchConfig()

	config, readError := validateConfig()
	if readError != nil {
		zap.L().Error("invalid configuration", zap.Error(readError))
		return nil, readError
	}

	return config, nil
}

func validateConfig() (*Config, error) {
	config := Config{}
	if len(viper.GetString("database.type")) <= 0 {
		return nil, errors.New("missing database.type")
	}
	config.DataBase.Type = viper.GetString("database.type")

	if len(viper.GetString("database.ip")) <= 0 {
		return nil, errors.New("missing database.type")
	}
	config.DataBase.Ip = viper.GetString("database.ip")

	if len(viper.GetString("database.port")) <= 0 {
		return nil, errors.New("missing database.type")
	}
	config.DataBase.Port = viper.GetString("database.port")

	if len(viper.GetString("database.username")) <= 0 {
		return nil, errors.New("missing database.username")
	}
	config.DataBase.Username = viper.GetString("database.username")

	if len(viper.GetString("database.password")) <= 0 {
		return nil, errors.New("missing database.password")
	}
	config.DataBase.Password = viper.GetString("database.password")

	if len(viper.GetString("database.instance")) <= 0 {
		return nil, errors.New("missing database.instance")
	}
	config.DataBase.Instance = viper.GetString("database.instance")

	return &config, nil
}
