package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type LogConfig struct {
	LogDir   string `json:"log_dir"`
	LogLevel uint   `json:"log_level"`
}

type EnvConfig struct {
	LogConfig
}

/*******************************************************************************
* Description: Read config.json
*       Input: config file path
*      Output:
*      Return:
*      Others:
*******************************************************************************/
func (c *EnvConfig) LoadConfig(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Failed to load config file:", err.Error())
		return err
	}

	c.LogConfig = LogConfig{
		LogDir:   "../logs/",
		LogLevel: 4, // 1:Fatal 2:Error 3:Warn 4:Info 5:Debug
	}

	if err := json.Unmarshal(data, c); err != nil {
		fmt.Println("Failed to unmarshal config json:", err.Error())
		return err
	}

	return nil
}

var CONF EnvConfig
