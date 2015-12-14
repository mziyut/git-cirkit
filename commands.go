package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mziyut/git-cirkit/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "init",
		Usage:  "",
		Action: command.CmdInit,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "feature",
		Usage:  "",
		Action: command.CmdFeature,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "hotfix",
		Usage:  "",
		Action: command.CmdHotfix,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "release",
		Usage:  "",
		Action: command.CmdRelease,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "tag",
		Usage:  "",
		Action: command.CmdTag,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
