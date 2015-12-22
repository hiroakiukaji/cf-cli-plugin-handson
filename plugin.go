package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type MyFirstPlugin struct{}

func (mfp *MyFirstPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("`cf push` from Plugin")

	var pushargs []string
	pushargs = append(pushargs, "push")
	pushargs = append(pushargs, args[1:]...)
	cliConnection.CliCommand(pushargs...)
}

func (mfp *MyFirstPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "MyFirstPlugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "my-first-plugin",
				HelpText: "This is a help text.",

				UsageDetails: plugin.Usage{
					Usage: "my-first-plugin\n   cf my-first-plugin",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(MyFirstPlugin))
}
