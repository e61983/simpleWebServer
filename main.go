package main

import (
	"log"
	"net/http"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Simple"
	app.Usage = "Create a simple web server"
	app.Action = run
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Usage:   "use port",
			EnvVars: []string{"WEB_PORT"},
			Aliases: []string{"p"},
		},
		&cli.StringFlag{
			Name:    "path",
			Usage:   "web root dir",
			EnvVars: []string{"WEB_ROOT_PATH"},
		},
	}
	app.Run(os.Args)
}

var port = "8080"
var path = "./webroot"

func run(c *cli.Context) error {
	if userPath := c.String("path"); "" != userPath {
		path = userPath
	}

	if userPort := c.String("port"); "" != userPort {
		port = userPort
	}

	log.Println("Web root at : " + path)
	log.Println("Listen at port : " + port)

	webroot := http.FileServer(http.Dir(path))
	log.Fatal(http.ListenAndServe(":"+port, webroot))
	return nil
}
