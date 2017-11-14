// Copyright (C) 2017 Miquel Sabaté Solà <msabate@suse.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

var gitCommit, version string

func versionString() string {
	str := version

	if gitCommit != "" {
		str += fmt.Sprintf(" with commit '%v'", gitCommit)
	}
	return fmt.Sprintf(`%v.
Copyright (C) 2017 Miquel Sabaté Solà <msabate@suse.com>
License GPLv3+: GNU GPL version 3 or later "<http://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.`, str)
}

func main() {
	// TODO: exec command (implementing also rake), logs (with warning on
	// containerized setups) and api thingies.

	app := cli.NewApp()
	app.Name = "portusctl"
	app.Usage = "Client for your Portus instance"
	app.UsageText = "portusctl <command> [arguments...]"
	app.HideHelp = true
	app.Version = versionString()

	app.CommandNotFound = func(context *cli.Context, cmd string) {
		fmt.Printf("Incorrect usage: command '%v' does not exist.\n\n", cmd)
		cli.ShowAppHelp(context)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server, s",
			Usage:  "The location where the Portus instance serves requests",
			EnvVar: "PORTUSCTL_API_SERVER",
		},
		cli.StringFlag{
			Name:   "token, t",
			Usage:  "The authentication token of the user for the Portus REST API",
			EnvVar: "PORTUSCTL_API_TOKEN",
		},
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "The user of the Portus REST API",
			EnvVar: "PORTUSCTL_API_USER",
		},
	}

	// TODO: add validate command
	app.Commands = []cli.Command{
		{
			Name:   "create",
			Usage:  "Create the given resource",
			Action: resourceDecorator(create),
			ArgsUsage: `<resource> [arguments...]

Where <resource> is the resource that you want to create.`,
		},
		{
			Name:   "delete",
			Usage:  "Delete the given resource",
			Action: resourceDecorator(delete),
			ArgsUsage: `<resource> [arguments...]

Where <resource> is the resource that you want to delete.`,
		},
		{
			Name:   "exec",
			Usage:  "Execute an arbitrary command on the environment of your Portus instance",
			Action: execCmd,
			ArgsUsage: `<command> [arguments...]

Where <command> is the command that you want to run on the environment of your
Portus instance. The successive arguments will be passed also to this command.`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "local, l",
					Value:  "/srv/Portus",
					Usage:  "The location on the current host of your Portus instance",
					EnvVar: "PORTUSCTL_EXEC_LOCATION",
				},
				cli.BoolFlag{
					Name:   "vendor, v",
					Usage:  "Use the local 'vendor' directory as the gem environment",
					EnvVar: "PORTUSCTL_EXEC_VENDOR",
				},
			},
		},
		{
			Name:   "get",
			Usage:  "Fetches info for the given resource",
			Action: resourceDecorator(get),
			ArgsUsage: `<resource> [arguments...]

Where <resource> is the resource that you want to fetch.`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "format, f",
					Value:  "default",
					Usage:  "The output format. Available options: default and json",
					EnvVar: "PORTUSCTL_FORMAT",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Update the given resource",
			Action: resourceDecorator(update),
			ArgsUsage: `<resource> [arguments...]

Where <resource> is the resource that you want to update.`,
		},
	}

	app.RunAndExitOnError()
}