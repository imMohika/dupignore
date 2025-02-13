package cmd

import (
	"github.com/adrg/xdg"
	"github.com/alecthomas/kong"
	"github.com/immohika/dupignore/internal/config"
	"log"
)

type GlobalVars struct {
	Config string `help:"Path to config file. Empty value will use XDG data directory." default:""`
}

type CLI struct {
	GlobalVars

	Dedup  DedupCmd  `cmd:"" help:"Remove duplicate entries from .gitignore"`
	Config ConfigCmd `cmd:"" help:"Manage global config file"`
}

func Run() {
	cli := CLI{}

	err := defaultGlobals(&cli.GlobalVars)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.From(cli.GlobalVars.Config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := kong.Parse(&cli,
		kong.Name("dupignore"),
		kong.Description("A CLI tool to remove duplicate entries from .gitignore"),
		kong.UsageOnError(),
		kong.Bind(cfg),
		kong.Bind(&cli.GlobalVars))

	err = ctx.Run(&cfg, &cli.GlobalVars)
	ctx.FatalIfErrorf(err)
}

func defaultGlobals(g *GlobalVars) error {
	if g.Config == "" {
		var err error
		g.Config, err = xdg.ConfigFile("dupignore/config.toml")
		if err != nil {
			return err
		}
	}
	return nil
}
