package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/urfave/cli"
	"github.com/voidint/tsdump/build"
	"github.com/voidint/tsdump/config"
	"github.com/voidint/tsdump/model"
	"github.com/voidint/tsdump/model/mysql"
	"github.com/voidint/tsdump/view/md"
	"github.com/voidint/tsdump/view/txt"
)

const (
	// TextView 纯文本视图
	TextView = "txt"
	// MarkdownView markdown文本视图
	MarkdownView = "md"
	// CSVView CSV文本视图
	CSVView = "csv"
	// JSONView JSON文本视图
	JSONView = "json"
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

var (
	c   config.Config
	out io.Writer = os.Stdout
)

func main() {
	app := cli.NewApp()
	app.Name = "tsdump"
	app.Usage = "Database table structure dump tool."
	app.Version = build.Version("0.1.0")
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "voidnt",
			Email: "voidint@126.com",
		},
	}

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
			Usage:       "Database name.",
			Destination: &c.DB,
		},
		cli.StringFlag{
			Name:        "V, viewer",
			Value:       "txt",
			Usage:       "Output viewer. Optional values: txt|md|csv|json",
			Destination: &c.Viewer,
		},
		cli.StringFlag{
			Name:        "o, output",
			Usage:       "Write to a file, instead of STDOUT.",
			Destination: &c.Output,
		},
		cli.BoolFlag{
			Name:        "D, debug",
			Usage:       "Enable debug mode.",
			Destination: &c.Debug,
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if c.Debug {
			fmt.Println(c)
		}

		repo, err := mysql.NewRepo(&c)
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		// Get metadata
		var dbs []model.DB
		if c.DB != "" {
			dbs, err = repo.GetDBs(&model.DB{
				Name: c.DB,
			})
		} else {
			dbs, err = repo.GetDBs(nil)
		}
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		if len(c.Output) > 0 {
			var f *os.File
			if f, err = os.Create(c.Output); err != nil {
				return cli.NewExitError(err, 1)
			}
			defer f.Close()
			out = f
		}

		// Output as target viewer
		switch c.Viewer {
		case TextView:
			_ = txt.NewView().Do(dbs, out)
		case MarkdownView:
			_ = md.NewView().Do(dbs, out)
		case CSVView:
			return cli.NewExitError(fmt.Sprintf("%q is not supported yet.", c.Viewer), 0)
		case JSONView:
			return cli.NewExitError(fmt.Sprintf("%q is not supported yet.", c.Viewer), 0)
		default:
			return cli.NewExitError(fmt.Sprintf("Unsupported viewer: %s", c.Viewer), 1)
		}

		return nil
	}

	app.Run(os.Args)
}
