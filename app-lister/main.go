package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

type AppLister struct{}

func (c *AppLister) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Application Lister",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "list-app",
				HelpText: "call CLI plugin APIs to list apps from all orgs",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "cf list-app",
				},
			},
		},
	}
}
func (c *AppLister) Run(cliConnection plugin.CliConnection, args []string) {
	orgs, err := cliConnection.GetOrgs()
	if err != nil {
		fmt.Println("Error getting orgs:", err)
		os.Exit(1)
	}

	for _, org := range orgs {
		_, err := cliConnection.CliCommandWithoutTerminalOutput("t", "-o", org.Name)
		if err != nil {
			fmt.Println("Error targeting org: ", org.Name)
			os.Exit(1)
		}

		apps, err := cliConnection.GetApps()
		if err != nil {
			fmt.Println("Error getting applications from org: ", org.Name)
			os.Exit(1)
		}

		for _, app := range apps {
			fmt.Println(app.Name)
		}
	}
}

func main() {
	plugin.Start(new(AppLister))
}
