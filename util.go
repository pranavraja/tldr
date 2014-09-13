package main

import (
	"encoding/json"
	"os/user"
	"fmt"
	"os"
	"io/ioutil"
)

type Config struct {
	UseCache string
	CacheDirectory string
}

func initialize() () {
	usr, _ := user.Current()
	os.MkdirAll(usr.HomeDir + "/.config/tldr/", 0755)
	configFile, err := os.Create(usr.HomeDir + "/.config/tldr/client.json")
	if err != nil {
		fmt.Println(err)
	}
	_, err = configFile.WriteString("{\n	\"useCache\": \"yes\",\n	\"cacheDirectory\": \"~/.tldr/cache/\"\n}\n")
	os.MkdirAll(usr.HomeDir + "/.tldr/cache/", 0755)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadConfig() (c *Config, err error) {
	var bfile []byte
	usr, _ := user.Current()
	cfgpath := usr.HomeDir + "/.config/tldr/client.json"
	if bfile, err = ioutil.ReadFile(cfgpath); err != nil {
		initialize()
		LoadConfig()
	}
	c = new(Config)
	err = json.Unmarshal(bfile, c)
	return
}
