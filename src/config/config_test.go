package config

import (
	// "fmt"
	"testing"
)

const CONFIG_FILE = "../../etc/config.json"
const CONFIG_WRONG_FILE = "../../etc/config_wrong.json"

const CONFIG_COUNT = 6

//Load config file
func TestLoadConfig(t *testing.T) {
	var config Config
	config.Load(CONFIG_FILE)
	if config.Log != "log/access.log" {
		t.Error("config log file error")
	}

	if config.ErrorLog != "log/error.log" {
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

	if !config.Routes[0].NeedCheckParams {
		t.Error("config route 1 NeedCheckParams error")
	}

}

//load error config file . this function is for just like nginx -t
func TestTestLoad(t *testing.T) {
	var config Config
	if !config.TestLoad(CONFIG_FILE) {
		t.Error("TestLoad right config file Error")
	}
	if config.TestLoad(CONFIG_WRONG_FILE) {
		t.Error("TestLoad wrong config file Error")
	}
}
