/* SPDX-License-Identifier: GPL-3.0-or-later */

package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const file_name string = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home_dir, file_name), nil
}

func Read() (Config, error) {
	config := Config{}
	filePath, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer f.Close()

	jsonData, err := io.ReadAll(f)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(jsonData, &config); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return config, nil
}

func write(cfg Config) error {

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonData, os.ModePerm)
}

func (c *Config) SetUser(name string) error {

	c.CurrentUserName = name

	return write(*c)
}

func GetUsername() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
