package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

type config struct {
	List map[string]string
}

var (
	pathRoot, pathConfig, pathSounds string
	cfg                              config
)

func init() {
	u, _ := user.Current()
	pathRoot = u.HomeDir + "/saysound/"
}

func main() {
	flag.Parse()
	flag.Arg(0)

	fs, err := ioutil.ReadDir(pathRoot)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, fi := range fs {
		if fi.IsDir() {
			continue
		}
		ext := filepath.Ext(fi.Name())
		if ext != "mp3" && ext != "wav" {
			continue
		}
	}
}
