package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/simonleung8/cli-plugin-examples/ssh-disabler/commands"
	"github.com/simonleung8/flags"
)

type AppLister struct{}

func (c *AppLister) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "SSH-Disabler",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "list-apps",
				HelpText: "List all apps across organizations",
				UsageDetails: plugin.Usage{
					Usage: "cf list-apps [-o ORG_NAME]",
					Options: map[string]string{
						"o": "The organization to be targeted",
					},
				},
			},
			plugin.Command{
				Name:     "disable-app-ssh",
				HelpText: "Disable application's ssh feature",
				UsageDetails: plugin.Usage{
					Usage: "cf disable-app-ssh [-o ORG_NAME]",
					Options: map[string]string{
						"o": "The organization to be targeted",
					},
				},
			},
		},
	}
}

func (c *AppLister) Run(cliConnection plugin.CliConnection, args []string) {
	fc, err := parseArguments(args)
	if err != nil {
		exitFail("Invalid flag provided: " + err.Error())
	}

	switch args[0] {
	case "list-apps":
		if fc.IsSet("organization") {
			err = commands.ListAppsInOneOrg(cliConnection, fc.String("organization"))
		} else {
			err = commands.ListAllApps(cliConnection)
		}
	case "disable-app-ssh":
		if fc.IsSet("organization") {
			err = commands.DisableAppsInOneOrg(cliConnection, fc.String("organization"))
		} else {
			err = commands.DisableAppsInAllOrg(cliConnection)
		}
	case "CLI-MESSAGE-UNINSTALL":
		fmt.Println("Thank you for using this plugin")
	}
	if err != nil {
		exitFail(err.Error())
	}
}

func parseArguments(args []string) (flags.FlagContext, error) {
	fc := flags.New()
	fc.NewStringFlag("organization", "o", "The organization to target for when listing apps")
	err := fc.Parse(args...)

	return fc, err
}

func exitFail(err string) {
	fmt.Println("exitFail\n", err)
	os.Exit(1)
}
