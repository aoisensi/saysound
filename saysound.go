package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"time"
)

var (
	config struct {
		Sounds map[string]struct {
			File []string `json:"file"`
		} `json:"sounds"`
	}
	list = flag.Bool("list", false, "show all sound list.")
)

func init() {
	path := os.Getenv("SAYSOUND")
	if path == "" {
		path = "$HOME/saysound"
	}
	os.Chdir(path)

	c, err := ioutil.ReadFile(path + "/config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	json.Unmarshal(c, &config)
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	flag.Parse()
	if *list {
		l := config.Sounds
		ls := make([]string, 0, len(l))
		for key := range l {
			ls = append(ls, key)
		}
		sort.Strings(ls)
		for _, k := range ls {
			fmt.Println(k)
		}
		return
	}
	args := flag.Args()
	for _, arg := range args {
		say(arg)
		break
	}
}

func say(sound string) {
	snd, ok := config.Sounds[sound]
	if !ok {
		return
	}
	count := len(snd.File)
	file := snd.File[rand.Intn(count)]
	exec.Command("cvlc", "--play-and-exit", file).Start()
}
