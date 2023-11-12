package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sm/model"
)

func GetConfig(configPath string) (model.Config, error) {
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		return model.Config{}, fmt.Errorf("unable to read file:%w", err)
	}

	// 解析 JSON 对象
	var config model.Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return model.Config{}, fmt.Errorf("unable to parse JSON:%w", err)
	}
	return config, nil
}

func FindConfigByAlias(configPath string, aliasInput string) (model.Server, error) {
	var serverResult model.Server
	config, err := GetConfig(configPath)
	if err != nil {
		return serverResult, errors.New("no matching server found")
	}
	serverList := config.ServerList
	for _, server := range serverList {
		if server.Alias == aliasInput {
			return server, nil
		}
	}
	return serverResult, errors.New("no matching server found")
}

func WriteConfig(config model.Config, filePath string) error {
	configJSON, err := json.MarshalIndent(config, "", "  ")

	if err != nil {
		return errors.New("unable to parse JSON")
	}

	err = os.WriteFile(filePath, configJSON, os.ModePerm)
	if err != nil {
		return errors.New("unable to write configuration file")
	}

	return nil
}
