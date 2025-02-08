package app

import (
	"naslook/internal/domain/file"

	"github.com/urfave/cli/v2"
)

func New() *cli.App {
	app := cli.NewApp()
	app.Name = "naslook"
	app.Commands = []*cli.Command{
		{
			Name: "dup-file",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{Name: "path"},
				&cli.StringSliceFlag{Name: "ignore"},
			},
			Action: actionDupFile,
		},
	}
	return app
}

func actionDupFile(ctx *cli.Context) error {
	return file.FileDeleteDup(ctx.Context,
		ctx.StringSlice("path"),
		ctx.StringSlice("ignore"),
	)
}
