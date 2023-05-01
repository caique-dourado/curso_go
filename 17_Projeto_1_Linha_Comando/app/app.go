package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IP's e nomes de servidores"
	//cria uma flag padrão para a cli.command
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}
	app.Commands = []cli.Command{
		//Criando os comandos do terminal
		{
			Name:  "ip",
			Usage: "Busca endereços de Ip's na internet",
			Flags: flags,
			//Ação do comando
			Action: BuscaIp,
		},

		{
			Name:  "servidores",
			Usage: "Busca nome de servidor na internet",
			Flags: flags,
			//Ação do comando
			Action: BuscaNome,
		},
	}
	return app
}

func BuscaIp(c *cli.Context) {
	host := c.String("host")
	//Obtem o Ip onde o site estar hospedado
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func BuscaNome(c *cli.Context) {
	host := c.String("host")
	//Obtem o nome do servidor
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
