package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sm/model"
)

func GetConfig(configPath string) model.Config {
	fileContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("无法读取文件：", err)
	}

	// 解析JSON对象
	var config model.Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		fmt.Println("无法解析JSON：", err)
	}
	return config
}

func FindConfigByAlias(configPath string, aliasInput string) (model.Server, error) {
	var serverResult model.Server
	config := GetConfig(configPath)
	serverList := config.ServerList
	for _, server := range serverList {
		if server.Alias == aliasInput {
			return server, nil
		}
	}
	return serverResult, errors.New("未找到匹配的服务器")
}
