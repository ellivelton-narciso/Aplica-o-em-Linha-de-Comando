package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flagHost := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "127.0.0.1",
		},
	}

	app.Commands = []cli.Command{

		{
			Name:   "ip",
			Usage:  "BuscaIPs de endereços na internet.",
			Flags:  flagHost,
			Action: buscarIps,
		},

		{
			Name:   "cname",
			Usage:  "Busca o DNS de quem pertence o Host",
			Flags:  flagHost,
			Action: buscarCNAME,
		},
		{
			Name:   "ns",
			Usage:  "Busca os servidores que o Host possui.", //nameservers
			Flags:  flagHost,
			Action: buscarNS,
		},
	}
	return app
}
func buscarIps(c *cli.Context) {
	host := c.String("host")
	fmt.Println(`Esses são os IPs do host: "` + host + `"`)
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarCNAME(c *cli.Context) {
	host := c.String("host")
	fmt.Println(`O DNS para o servidor ` + host + ` é:`)
	ips, erro := net.LookupCNAME(host)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(ips)

}

func buscarNS(c *cli.Context) {
	host := c.String("host")
	fmt.Println(`Esses são os servidores disponíveis para: ` + host)
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
