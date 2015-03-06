package main

import (
	"fmt"
	"github.com/johntdyer/anybar-go/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/johntdyer/anybar-go/Godeps/_workspace/src/github.com/wsxiaoys/terminal"
	"log"
	"net"
	"os"
)

func main() {

	app := cli.NewApp()

	app.Version = "0.0.1"
	app.Name = "anybar-go"
	app.Usage = "Anybar CLI"
	app.Author = "John Dyer"
	app.Email = "johntdyer@gmail.com"

	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:   "port, p",
			Value:  "1738",
			Usage:  "Port to connect to anybar",
			EnvVar: "ANYBAR_PORT",
		},

		cli.StringFlag{Name: "address, a",
			Value: "localhost",
			Usage: "Address to send message.",
		},
	}

	app.Action = func(c *cli.Context) {
		msg := ""
		if len(c.Args()) > 0 {
			msg = c.Args()[0]
			addr := fmt.Sprintf("%s:%s", c.String("address"), c.String("port"))
			sendPacket(msg, addr)
		} else {

			terminal.Stdout.Color("r").Print("Message is required").Nl().Reset()

		}

	}

	app.Run(os.Args)

}

func sendPacket(msg string, addr string) {
	laddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	maddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	_, err = c.WriteToUDP([]byte(msg), maddr)
	if err != nil {
		log.Fatal(err)
	}

	terminal.Stdout.Color("b").Print(fmt.Sprintf("Sent '%s' successfully", msg)).Nl().Reset()

}
