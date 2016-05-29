package main

import (
	"encoding/base64"
	"fmt"

	"github.com/urfave/cli"
)

func new(c *cli.Context) error {
	publicKey, privateKey := generateKeys()
	fmt.Printf("Address: %s\n", base64.StdEncoding.EncodeToString(publicKey))
	fmt.Printf("Private key: %s\n", base64.StdEncoding.EncodeToString(privateKey))
	return nil
}
