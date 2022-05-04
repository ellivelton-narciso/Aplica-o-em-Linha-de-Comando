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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "127.0.0.1",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet.",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "cname",
			Usage:  "Busca o DNS de quem pertence o Host",
			Flags:  flags,
			Action: buscarCNAME,
		},
		{
			Name:   "",
			Usage:  "",
			Flags:  flags,
			Action: ab,
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
	cname := c.String("host")
	fmt.Println(`O DNS para o servidor ` + cname + ` é:`)
	cips, cerro := net.LookupCNAME(cname)
	if cerro != nil {
		log.Fatal(cerro)
	}
	fmt.Println(cips)
}

func ab(c *cli.Context) {

}
