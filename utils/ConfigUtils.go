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
		return model.Config{}, fmt.Errorf("无法读取文件：%w", err)
	}

	// 解析 JSON 对象
	var config model.Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return model.Config{}, fmt.Errorf("无法解析 JSON：%w", err)
	}
	return config, nil
}

func FindConfigByAlias(configPath string, aliasInput string) (model.Server, error) {
	var serverResult model.Server
	config, err := GetConfig(configPath)
	if err != nil {
		return serverResult, errors.New("未找到匹配的服务器")
	}
	serverList := config.ServerList
	for _, server := range serverList {
		if server.Alias == aliasInput {
			return server, nil
		}
	}
	return serverResult, errors.New("未找到匹配的服务器")
}
