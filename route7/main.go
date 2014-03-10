package main

import (
	"config"
	"flag"
	"fmt"
	"os"
)

const CONFIG_FILE = "etc/config.json"

const CONFIG_COUNT = 6

var conf config.Config

func main() {
	fmt.Println(APP_NAME, " version ", VERSION)

	flag.Usage = show_usage
	conf.Load(CONFIG_FILE)
}

func show_usage() {
	fmt.Fprintf(os.Stderr,
		"Usage: %s \n",
		os.Args[0])
	flag.PrintDefaults()

}
