package yml

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/ppzxc/golang-boilerplate-in-my-case/util/config/logger"
	"github.com/ppzxc/golang-boilerplate-in-my-case/util/err"
	"github.com/spf13/viper"
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

	// database
	if len(viper.GetString("database.type")) <= 0 {
		return nil, errors.New("missing database.type")
	}
	config.DataBase.Type = viper.GetString("database.type")

	if len(viper.GetString("database.host")) <= 0 {
		return nil, errors.New("missing database.host")
	}
	config.DataBase.Host = viper.GetString("database.host")

	if len(viper.GetString("database.port")) <= 0 {
		return nil, errors.New("missing database.port")
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

	// http
	if len(viper.GetString("http.addr")) <= 0 {
		return nil, errors.New("missing http.addr")
	}
	config.Http.Addr = viper.GetString("http.addr")

	if len(viper.GetString("http.context.timeout")) <= 0 {
		return nil, errors.New("missing http.context.timeout")
	}
	config.Http.Context.Timeout = viper.GetInt("http.context.timeout")

	return &config, nil
}
