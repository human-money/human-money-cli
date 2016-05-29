package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ed25519"
)

func pay(c *cli.Context) error {
	sourcePublicKey, _ := generateKeys()
	publicKey, privateKey := generateKeys()
	t := &Transaction{
		Source:      sourcePublicKey,
		Destination: publicKey,
		Amount:      proto.Uint32(10),
	}

	data, _ := proto.Marshal(t)
	t.Signature = ed25519.Sign(privateKey, data)

	data, _ = proto.Marshal(t)
	postTransaction(data)

	return nil
}

func postTransaction(data []byte) {
	url := "http://localhost:4000/transactions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func generateKeys() (ed25519.PublicKey, ed25519.PrivateKey) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	return publicKey, privateKey
}

func main() {
	app := cli.NewApp()
	app.Name = "pay"
	app.Usage = "A stable cryptocurrency built for everyone"
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "generate a new key",
			Action:  new,
		},
		{
			Name:    "address",
			Aliases: []string{"a"},
			Usage:   "show your address",
			Action:  address,
		},
		{
			Name:    "balance",
			Aliases: []string{"b"},
			Usage:   "show your balance",
			Action:  balance,
		},
		{
			Name:    "setup",
			Aliases: []string{"s"},
			Usage:   "Creates a .hm file",
			Action:  setup,
		},
		{
			Name:    "pay",
			Aliases: []string{"p"},
			Usage:   "pay <address>",
			Action:  pay,
		},
	}

	app.Run(os.Args)
}
