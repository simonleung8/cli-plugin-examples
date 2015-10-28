package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/simonleung8/cli-plugin-examples/app-lister2/apps_repository"
	"github.com/simonleung8/flags"
)

type AppLister struct{}

func (c *AppLister) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Application Lister 2",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "all-apps",
				HelpText: "Access Cloud Controller directly for all apps",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "cf all-apps",
				},
			},
		},
	}
}
func (c *AppLister) Run(cliConnection plugin.CliConnection, args []string) {
	fc, err := parseArguments(args)
	if err != nil {
		fmt.Println("invalid flag provided:", err)
		os.Exit(1)
	}

	if fc.IsSet("organization") {
		// list only apps in org
	} else {
		// list all apps
		apps, err := appsRepository.GetAllApps(cliConnection)
		if err != nil {
			fmt.Println("Error getting list of applications", err)
			os.Exit(1)
		}
		for i, app := range apps {
			fmt.Println(i, app.Entity.Name)
		}
	}

}

func main() {
	plugin.Start(new(AppLister))
}

func parseArguments(args []string) (flags.FlagContext, error) {
	fc := flags.New()
	fc.NewStringFlag("organization", "o", "The organization to target for when listing apps")
	err := fc.Parse(args...)

	return fc, err
}
