package main

import (
	"fmt"
	"github.com/johntdyer/anybar-go/Godeps/_workspace/src/github.com/codegangsta/cli"
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
		cli.StringFlag{Name: "port, p", Value: "1738", Usage: "Port to connect to any-bar. Default 1738"},
		cli.StringFlag{Name: "address, a", Value: "localhost", Usage: "Address to send message. Default localhost"},
		cli.StringFlag{Name: "msg, m", Value: "", Usage: "Message to send to anybar"},
	}

	app.Action = func(c *cli.Context) {
		if c.String("msg") == "" {
			fmt.Println("Message is required")
			os.Exit(0)
		}

		addr := fmt.Sprintf("%s:%s", c.String("address"), c.String("port"))

		sendPacket(c.String("msg"), addr)

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
	fmt.Printf("Sent %s, success\n", msg)

}
