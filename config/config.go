package config

import (
	"bytes"
	"io"
	"net/url"
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
	Database struct {
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Timeout  string `yaml:"timeout"`
	} `yaml:"database"`

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

func (c *Config) GetDatabaseURL() string {
	databaseURL := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.Database.Username, c.Database.Password),
		Host:     c.Database.Host,
		Path:     c.Database.Database,
		RawQuery: "sslmode=disable",
	}

	return databaseURL.String()
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
