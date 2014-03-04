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
	if config.Log != "log/access.log" {
		t.Error("config log file error")
	}
	if config.Listen.Host != "0.0.0.0" {
		t.Error("config listen host error")
	}

	if config.Listen.Port != 9000 {
		t.Error("config listen port error")
	}

	if config.Routes[0].Id != 900 {
		t.Error("config route 1 id is 900 error")
	}

	if config.Routes[0].need_check_sparams {
		t.Error("config listen port error")
	}

}
