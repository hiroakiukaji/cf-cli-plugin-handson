package main

import (
	"fmt"
	"os/exec"

	"github.com/cloudfoundry/cli/plugin"
)

type MyFirstPlugin struct{}

func (mfp *MyFirstPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("----------------")
	fmt.Println(" branch checker ")
	fmt.Println("----------------")

	output, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()
	fmt.Println("On branch " + string(output))

	var yn string
	fmt.Printf("Is it OK? (yes/no) :")
	fmt.Scanf("%s", &yn)

	if yn == "yes" {
		pushargs := []string{"push"}
		pushargs = append(pushargs, args[1:]...)
		cliConnection.CliCommand(pushargs...)
	} else {
		fmt.Println("Bye!")
	}
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
