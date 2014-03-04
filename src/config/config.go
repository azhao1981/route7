/*
  config.go by weizhao
  config file for etc/config.json as default
  1. json only
  2. remote file support
  3. test config json file support
*/
package config

import (
	"../route"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Default route.Route
	Routes  []route.Route
	Listen  struct {
		Host string
		Port int
	}
	Log        string `json:"log"`
	Error_log  string `json:"error_log"`
	Admin      string `json:"admin"`
	PprofHttpd string `json:"pprof_httpd"`
}

func (c *Config) Load(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Load Config File Error.")
	}

	if json.Unmarshal([]byte(b), &c) != nil {
		panic("Parse json failed.")
	}
	fmt.Println(c)
	fmt.Println("Load config ", path, "OK!!")
}
