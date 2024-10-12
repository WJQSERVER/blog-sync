package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   ServerConfig
	Log      LogConfig
	Download DownloadConfig
	Hugo     HugoConfig
}

type ServerConfig struct {
	CycleInterval int `toml:"cycle_interval"`
}

type LogConfig struct {
	LogFilePath string `toml:"logfilepath"`
	MaxLogSize  int    `toml:"maxlogsize"`
}

type DownloadConfig struct {
	DownloadUrl string `toml:"downloadurl"`
	Username    string `toml:"username"`
	Password    string `toml:"password"`
	SavePath    string `toml:"savepath"`
}

type HugoConfig struct {
	BaseUrl  string `toml:"base_url"`
	UnzipDir string `toml:"unzip_dir"`
}

// LoadConfig 从 TOML 配置文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
