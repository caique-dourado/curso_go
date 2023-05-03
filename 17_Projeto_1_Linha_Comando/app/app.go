package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Busca ip's e nomes de servidores"
	app.Usage = "Busca ip's e nomes de servidores na internet"
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "amazon.com",
		},
	}
	app.Commands = []cli.Command{

		{

			Name:   "ip",
			Action: buscaIp,
			Flags:  flag,
		},
		{
			Name:   "servidores",
			Action: buscaServidor,
			Flags:  flag,
		},
	}
	return app

}

func buscaIp(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscaServidor(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}
