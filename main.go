package main

import (
	"app"
	"config"
	"flag"
	"fmt"
	"os"
	"runtime"
)

const CONFIG_FILE = "etc/config.json"

const CONFIG_COUNT = 6

var conf config.Config

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
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
