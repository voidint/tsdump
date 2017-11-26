package main

import (
	"os"
	"os/user"

	"github.com/urfave/cli"
	"github.com/voidint/tsdump/build"
	"github.com/voidint/tsdump/config"
)

var (
	username string
)

func init() {
	u, err := user.Current()
	if err == nil {
		username = u.Username
	}
}

var c config.Config

func main() {
	app := cli.NewApp()
	app.Name = ""
	app.Usage = ""
	app.Version = build.Version("0.1.0")

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "H, host",
			Value:       "127.0.0.1",
			Usage:       "Connect to host.",
			Destination: &c.Host,
		},
		cli.IntFlag{
			Name:        "P, port",
			Value:       3306,
			Usage:       "Port number to use for connection.",
			Destination: &c.Port,
		},
		cli.StringFlag{
			Name:        "u, user",
			Value:       username,
			Usage:       "User for login if not current user.",
			Destination: &c.Username,
		},
		cli.StringFlag{
			Name:        "p, password",
			Usage:       "Password to use when connecting to server.",
			Destination: &c.Password,
		},
		cli.StringFlag{
			Name:        "d, db",
			Usage:       "MySQL database name",
			Destination: &c.DB,
		},
		cli.StringFlag{
			Name:        "f, formatter",
			Value:       "txt",
			Usage:       "text, csv, markdown",
			Destination: &c.Formatter,
		},
		cli.StringFlag{
			Name:        "o, output",
			Usage:       "Write to a file, instead of STDOUT.",
			Destination: &c.Output,
		},
	}
	app.Action = func(ctx *cli.Context) error {
		return nil
	}

	app.Run(os.Args)
}
