package cli

import (
	"fmt"
	"gotranslator/src/translator"

	"github.com/urfave/cli"
)

func Prepare() *cli.App {
	app := cli.NewApp()
	app.Name = "Go Translator ~ Developed by @instalando"
	app.Usage = "Translate any text to any language"

	Flags := []cli.Flag{
		cli.StringFlag{
			Name:  "from",
			Usage: "Language to translate from",
			Value: "auto",
		},
		cli.StringFlag{
			Name:  "to",
			Usage: "Language to translate to",
			Value: "pt",
		},
		cli.StringFlag{
			Name:  "text",
			Usage: "Text to translate",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "translate",
			Usage:  "Translate text",
			Action: translate,
			Flags:  Flags,
		},
	}

	return app
}

func translate(ctx *cli.Context) {
	from := ctx.String("from")
	to := ctx.String("to")
	text := ctx.String("text")

	res, err := translator.Translate(text, from, to)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
