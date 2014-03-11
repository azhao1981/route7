package main

import (
	"flag"
	"fmt"
	"github.com/azhao1981/route7/app"
	"github.com/azhao1981/route7/config"
	"os"
)

const CONFIG_FILE = "etc/config.json"

const CONFIG_COUNT = 6

var conf config.Config

func main() {
	fmt.Println(app.APP_NAME, "version", app.VERSION)

	flag.Usage = show_usage
	if !conf.TestLoad(CONFIG_FILE) {
		panic("Config File error ,Try it later!")
	}
	app.Run(&conf)
}

func show_usage() {
	fmt.Fprintf(os.Stderr,
		"Usage: %s \n",
		os.Args[0])
	flag.PrintDefaults()
}
