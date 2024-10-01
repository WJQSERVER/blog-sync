package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LogFilePath string `yaml:"logfilepath"`
	DownloadUrl string `yaml:"downloadurl"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	SavePath    string `yaml:"savepath"`
	Unzipdir    string `yaml:"unzip_dir"`
	BaseURL     string `yaml:"base_url"`
}

// LoadConfig 从 YAML 配置文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
