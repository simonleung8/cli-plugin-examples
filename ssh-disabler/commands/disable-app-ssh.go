package commands

import (
	"errors"
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/simonleung8/cli-plugin-examples/ssh-disabler/apps_repository"
)

func DisableAppsInAllOrg(cliConnection plugin.CliConnection) error {
	orgs, err := cliConnection.GetOrgs()
	if err != nil {
		return errors.New("Error getting list of organizations: " + err.Error())
	}

	for _, org := range orgs {
		apps, err := GetAppsInOneOrg(cliConnection, org.Name)
		if err != nil {
			fmt.Println("Failed to get apps in organization '" + org.Name + "'")
			continue
		}

		for _, app := range apps {
			err = appsRepository.DisableAppSSH(cliConnection, app.Guid)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	return nil
}

func DisableAppsInOneOrg(cliConnection plugin.CliConnection, orgName string) error {
	apps, err := GetAppsInOneOrg(cliConnection, orgName)
	if err != nil {
		fmt.Println("exitFail to get apps in organization '" + orgName + "'")
	}

	for _, app := range apps {
		err = appsRepository.DisableAppSSH(cliConnection, app.Guid)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}
