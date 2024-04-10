package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	KEY = "key"
	URL = "url"
)

type Config struct {
	configName string
	configDir  string
	homeDir    string
}

func NewConfig(configName string) (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return &Config{
		homeDir:    homeDir,
		configName: configName,
		configDir:  ".gosrt",
	}, nil
}

func (cfg *Config) Validate() error {
	abs := filepath.Join(cfg.homeDir, cfg.configDir, cfg.configName)
	if _, err := os.Stat(abs); os.IsNotExist(err) {
		return err
	}

	if key := cfg.GetKey(); len(key) == 0 {
		return fmt.Errorf("config 'key' is emty")
	}
	return nil
}

func (cfg *Config) Config() error {
	configDir := filepath.Join(cfg.homeDir, cfg.configDir)
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return err
	}

	configName := filepath.Join(configDir, cfg.configName)

	if _, err := os.Stat(configName); os.IsNotExist(err) {
		os.Create(configName)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	err := viper.ReadInConfig()
	return err
}

func (cfg *Config) Set(key, value string) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Printf("fatal error writing config file: %s \n", err)
		return
	}
}

func (cfg *Config) GetKey() string {
	return viper.GetString(KEY)
}
func (cfg *Config) GetUrl() string {
	return viper.GetString(URL)
}
func (cfg *Config) Get(key string) any {
	return viper.Get(key)
}
