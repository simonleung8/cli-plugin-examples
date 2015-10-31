package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/cloudfoundry/cli/plugin/models"
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
					Usage: "cf list-apps",
				},
			},
		},
	}
}

func (c *AppLister) Run(cliConnection plugin.CliConnection, args []string) {
	fc, err := parseArguments(args)
	if err != nil {
		failed("Invalid flag provided: " + err.Error())
	}

	switch args[0] {
	case "list-apps":
		if fc.IsSet("organization") {
			orgName := fc.String("organization")
			apps, err := getAppsInOrg(cliConnection, orgName)
			if err != nil {
				failed("Error getting list of organizations: " + err.Error())
			}
			printAppsName(orgName, apps)
		} else {
			listAllApps(cliConnection)
		}
	case "disable-ssh-all":
		disableAppsInAllOrg(cliConnection)
	case "CLI-MESSAGE-UNINSTALL":
		return
	}
}

func parseArguments(args []string) (flags.FlagContext, error) {
	fc := flags.New()
	fc.NewStringFlag("organization", "o", "The organization to target for when listing apps")
	err := fc.Parse(args...)

	return fc, err
}

func listAllApps(cliConnection plugin.CliConnection) {
	orgs, err := cliConnection.GetOrgs()
	if err != nil {
		failed("Error getting list of organizations: " + err.Error())
	}

	for _, org := range orgs {
		apps, err := getAppsInOrg(cliConnection, org.Name)
		if err != nil {
			fmt.Println("Failed to get apps in organization '" + org.Name + "'")
			continue
		}

		printAppsName(org.Name, apps)
	}
}

func disableAppsInAllOrg(cliConnection plugin.CliConnection) {
	orgs, err := cliConnection.GetOrgs()
	if err != nil {
		failed("Error getting list of organizations: " + err.Error())
	}

	for _, org := range orgs {
		apps, err := getAppsInOrg(cliConnection, org.Name)
		if err != nil {
			fmt.Println("Failed to get apps in organization '" + org.Name + "'")
			continue
		}

		disableApps(org.Name, apps)
	}
}

func getAppsInOrg(cliConnection plugin.CliConnection, orgName string) ([]plugin_models.GetAppsModel, error) {
	_, err := cliConnection.CliCommandWithoutTerminalOutput("target", "-o", orgName)
	if err != nil {
		return []plugin_models.GetAppsModel{}, errors.New("Failed to target org '" + orgName + "'")
	}

	apps, err := cliConnection.GetApps()
	if err != nil {
		return []plugin_models.GetAppsModel{}, errors.New("Failed to get apps in organization '" + orgName + "'")
	}

	return apps, nil
}

func printAppsName(orgName string, apps []plugin_models.GetAppsModel) {
	for _, app := range apps {
		fmt.Println(orgName, ":", app.Name)
	}
}

func failed(err string) {
	fmt.Println("FAILED\n", err)
	os.Exit(1)
}
