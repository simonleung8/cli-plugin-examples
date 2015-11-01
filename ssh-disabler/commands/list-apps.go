package commands

import (
	"errors"
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/cloudfoundry/cli/plugin/models"
)

func ListAllApps(cliConnection plugin.CliConnection) error {
	orgs, err := cliConnection.GetOrgs()
	if err != nil {
		return errors.New("Error getting list of organizations: " + err.Error())
	}

	for _, org := range orgs {
		apps, err := GetAppsInOneOrg(cliConnection, org.Name)
		if err != nil {
			fmt.Println("Warning: Failed to get apps in organization '" + org.Name + "'")
			continue
		}

		PrintAppsName(org.Name, apps)
	}

	return nil
}

func ListAppsInOneOrg(cliConnection plugin.CliConnection, orgName string) error {
	apps, err := GetAppsInOneOrg(cliConnection, orgName)
	if err != nil {
		return errors.New("Error getting list of organizations: " + err.Error())
	}

	PrintAppsName(orgName, apps)

	return nil
}

func GetAppsInOneOrg(cliConnection plugin.CliConnection, orgName string) ([]plugin_models.GetAppsModel, error) {
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

func PrintAppsName(orgName string, apps []plugin_models.GetAppsModel) {
	for _, app := range apps {
		fmt.Println(orgName, ":", app.Name)
	}
}
