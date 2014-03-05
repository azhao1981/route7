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
	ErrorLog   string `json:"error_log"`
	Admin      string `json:"admin"`
	PprofHttpd string `json:"pprof_httpd"`
}

// load config file and init route
func (c *Config) Load(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Load Config File Error.")
	}
	if json.Unmarshal([]byte(b), &c) != nil {
		panic("Parse json failed.")
	}
	for index, _ := range c.Routes {
		c.Routes[index].Afterload()
	}
	fmt.Println("Load config ", path, "OK!!")
}

// test config file . this function is for support just like nginx -t
// simple one , TODO is tell the error detail
func (c *Config) TestLoad(path string) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
			fmt.Println("Load config ", path, "Error!!", r)
		}
	}()
	c.Load(path)
	ok = true
	return
}
