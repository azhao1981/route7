/*
  config.go by weizhao
  config file for etc/config.json as default
  1. json only
  2. remote file support
  3. test config json file support
*/
package config

import (
	"fmt"
	// "route"
	"io/ioutil"
)

type Config struct {
	path string
	// routes []Route
	listen struct {
		host string
		port int
	}
	log        string `json:"log"`
	error_log  string `json:"error_log"`
	admin      string `json:"admin`
	PprofHttpd string `json:"pprof_httpd"`
}

func (c *Config) Load(path string) {
	c.path = path
	b, err := ioutil.ReadFile(c.path)
	if err != nil {
		panic("Load Config File Error.")
	}
	fmt.Printf("%s", b)
	fmt.Println("Load config ", path, "OK!!")
}
