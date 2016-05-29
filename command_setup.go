package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

const DEFAULT_SERVER = "http://money-1.human.money"

type Config struct {
	PrivateKey string
	Server     string
}

func setup(c *cli.Context) error {
	fmt.Println("Generating human money address...")
	publicKey, privateKey := generateKeys()
	server := getServer()
	config := Config{
		PrivateKey: base64.StdEncoding.EncodeToString(privateKey),
		Server:     server,
	}
	a, _ := yaml.Marshal(&config)
	filename := os.Getenv("HOME") + "/.hm"
	ioutil.WriteFile(filename, a, 0600)

	fmt.Println("Success! Saved settings to ~/.hm")
	fmt.Printf("Server: %s\n", server)
	fmt.Printf("Address: %s\n", base64.StdEncoding.EncodeToString(publicKey))
	fmt.Printf("Private key: %s\n", base64.StdEncoding.EncodeToString(privateKey))
	return nil
}

func getServer() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Server (%s): ", DEFAULT_SERVER)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")

	if text == "" {
		return DEFAULT_SERVER
	} else {
		return text
	}
}
