package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func New(appName, repoName string) (*Config, error) {
	cfg := new(Config)

	err := initLocalConfig(cfg, repoName)
	if err != nil {
		log.Fatalf("error init local config: %v \n", err)
		return nil, err
	}

	return cfg, err
}

func initLocalConfig(cfg *Config, repoName string) error {
	localFile := getConfigFile(repoName)
	f, err := os.Open(localFile)
	if err != nil {
		log.Fatalf("fail to open file: %v \n", err)
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		log.Fatalf("fail to decode with err: %v \n", err)
		return err
	}

	return nil
}

func getConfigFile(repoName string) string {
	var (
		env      = "development"
		filename = fmt.Sprintf("%s.yaml", repoName)
	)

	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("fail get file path executable: %v \n", err)
	}

	return filepath.Join(ex, "../files", env, filename)
}
