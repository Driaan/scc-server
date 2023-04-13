package main

// Project: System Configuration Collector (SCC)
// Description: Collects system configuration information and stores it in a database

import (
	"bytes"
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

// var log = logrus.New()

func main() {
	viper.SetDefault("loglevel", "info")
	viper.SetDefault("logpath", "/var/log/scc.log")
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths

	viper.SetEnvPrefix("scc")
	viper.SetConfigName("scc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.SetFormatter(&log.JSONFormatter{})

	fmt.Println("hello world")
	log.Printf("hello %s", "world")

	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	viper.Get("name")
	scan_os()
}

func visit(path string, di fs.DirEntry, err error) error {
	// fmt.Printf("Visited: %s\n", path)
	return nil
}

func scan_os() {
	// walk entire OS filepath and get os.stat for each and measure time taken
	// using filepath.walk

	root := "/etc"
	current_time := time.Now()
	fmt.Println(current_time)
	err := filepath.WalkDir(root, visit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(current_time))

}

// TODO:
