package main

import (
	"config"
	"fmt"
)

const CONFIG_FILE = "../../etc/config_test.json"

const CONFIG_COUNT = 6

var conf Config

func main() {
	fmt.Println("hello")

	conf.Load(CONFIG_FILE)
}
