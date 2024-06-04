package config

import (
	"bytes"
	"io"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Secret string `yaml:"secret"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
}

// Config ...
type Config struct {
	Settings Settings `yaml:"settings"`
}

// NewConfig ...
func NewConfig() (*Config, error) {
	absPath, _ := filepath.Abs(".")
	file, err := ExpandEnv(filepath.Join(path.Clean(absPath), "config", "config.yaml"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func ExpandEnv(configs string) (io.Reader, error) {
	file, err := os.Open(configs)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
	}()

	bufferConfigs := new(bytes.Buffer)
	_, err = bufferConfigs.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	bytesConfigs := []byte(os.ExpandEnv(bufferConfigs.String()))
	return bytes.NewReader(bytesConfigs), nil
}
