package config

import (
	"fmt"

	"github.com/chubin518/kestrel-layout-basic/pkg/env"
	"github.com/chubin518/kestrel-layout-basic/pkg/graceful"
	"github.com/chubin518/kestrel-layout-basic/pkg/logging"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf = &Config{
	Logging: &logging.LoggingOptions{},
	Server:  &graceful.GracefulOptions{},
}

type Config struct {
	Logging *logging.LoggingOptions   `mapstructure:"logging"`
	Server  *graceful.GracefulOptions `mapstructure:"server"`
}

func Init(paths ...string) error {
	file := fmt.Sprintf("conf/%s.yaml", env.Active())
	if len(paths) >= 1 {
		file = paths[0]
	}
	v := viper.New()
	v.SetConfigFile(file)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		logging.Infof("config changed: %v", in.Name)
		if err := v.Unmarshal(conf); err != nil {
			logging.Errorf("error unmarshaling changed config: %v", err)
			return
		}
	})
	if err := v.Unmarshal(conf); err != nil {
		return err
	}
	return nil
}

func Logging() *logging.LoggingOptions {
	return conf.Logging
}

func Server() *graceful.GracefulOptions {
	return conf.Server
}
