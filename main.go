package main

import (
	"fmt"
	// "log"
	"os"
	"runtime"

	// "bitbucket.org/grengojbo/ads-core/config"
	// "github.com/nu7hatch/gouuid"
	"github.com/codegangsta/cli"
)

var (
	Version   = "0.1.0"
	BuildTime = "2015-09-20 UTC"
	GitHash   = "c00"
)

type Switch struct {
	Debug   bool
	Verbose bool
}

type Application struct {
	Logs string
	// Config string
	// DB     *sql.DB
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func main() {
	app := cli.NewApp()
	app.Name = "sw-cli"
	app.Version = Version
	app.Usage = "Run Cisco Switch command"
	app.Author = "Oleg Dolya"
	app.Email = "oleg.dolya@gmail.com"
	app.EnableBashCompletion = true
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "verbose",
			// Value: "false",
			Usage: "Verbose mode",
		},
		cli.BoolFlag{
			Name: "debug",
			// Value: "false",
			Usage: "Debug mode",
		},
		cli.StringFlag{
			Name:   "passwd, p",
			Usage:  "SSH pasword",
			EnvVar: "CISCO_PASSWD",
		},
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "SSH user name",
			EnvVar: "CISCO_USER",
		},
		cli.StringFlag{
			Name:  "host",
			Usage: "connection to only host",
		},
		cli.BoolTFlag{
			Name:  "teltet, t",
			Usage: "Connect to telnet mode",
		},
		cli.StringFlag{
			Name:   "config, c",
			Value:  "config/config.yml",
			Usage:  "config file to use (config/config.yml)",
			EnvVar: "APP_CONFIG",
		},
	}
	// app.Before = func(c *cli.Context) error {
	// log.Println("Load config:", c.GlobalString("config"))
	// cfg, err := getConfig(c)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 			return
	// 		}
	// 		log.Println(cfg)
	// setting.CustomConf = c.GlobalString("config")
	// setting.NewConfigContext()
	// setting.NewServices()
	// return nil
	// }
	app.Run(os.Args)
}
