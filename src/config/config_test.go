package config

import (
	// "fmt"
	"testing"
)

const CONFIG_FILE = "../../etc/config.json"

const CONFIG_COUNT = 6

//测试配置文件解析
func TestLoadConfig(t *testing.T) {
	var config Config
	config.Load(CONFIG_FILE)
}
