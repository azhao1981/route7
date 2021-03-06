package config

import (
	"fmt"
	"net/http"
	"testing"
)

const CONFIG_FILE = "../etc/config.json"
const CONFIG_WRONG_FILE = "../etc/config_wrong.json"

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
	if config.Routes[0].Id != 100 {
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

// FindRoute test
// test request with params "DeviceID=AAA", route Id:1000 will be return
func TestFindRoute(t *testing.T) {
	var config Config
	if !config.TestLoad(CONFIG_FILE) {
		t.Error("TestLoad config file Error")
	}

	req_url := `http://localhost:10003/ticket/req.do?DeviceID=AAA`
	req, _ := http.NewRequest("GET", req_url, nil)

	r := config.FindRoute(req)
	fmt.Println("# FindRoute() id: ", r.Id)
	if r == nil || r.Id != 1000 {
		t.Error("Test #FindRoute should find route 1000")
	}
}

// FindRoute test
// test request with params "DeviceID=AAA", route Id:1000 will be return
func TestFindRouteDefault(t *testing.T) {
	var config Config
	if !config.TestLoad(CONFIG_FILE) {
		t.Error("TestLoad config file Error")
	}

	req_url := `http://localhost:10003/ticket/req.do?DeviceID=NotThisRoute`
	req, _ := http.NewRequest("GET", req_url, nil)

	r := config.FindRoute(req)
	fmt.Println("# FindRoute() id: ", r.Id)
	if r == nil || r.Id != 0 {
		t.Error("Test #FindRoute should find route 0")
	}
}
