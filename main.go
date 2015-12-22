package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Masterminds/glide/cfg"
)

func main() {
	yml, err := ioutil.ReadFile("glide.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf, err := cfg.ConfigFromYaml(yml)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hash, err := conf.Hash()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(hash)
}
