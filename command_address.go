package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/FactomProject/ed25519"

	"gopkg.in/yaml.v2"

	"github.com/urfave/cli"
)

func address(c *cli.Context) error {
	filename := os.Getenv("HOME") + "/.hm"
	s, _ := ioutil.ReadFile(filename)

	config := Config{}
	yaml.Unmarshal(s, &config)

	fmt.Printf("Address: %s\n", getPublicKey(config.PrivateKey))

	return nil
}

func getPublicKey(s string) string {
	var data [64]byte
	d, _ := base64.StdEncoding.DecodeString(s)

	copy(data[:], d)
	p := ed25519.GetPublicKey(&data)
	var pub []byte
	pub = p[:]
	return base64.StdEncoding.EncodeToString(pub)
}
