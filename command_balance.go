package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func balance(c *cli.Context) error {
	filename := os.Getenv("HOME") + "/.hm"
	s, _ := ioutil.ReadFile(filename)

	fmt.Printf(string(s))

	return nil
}
