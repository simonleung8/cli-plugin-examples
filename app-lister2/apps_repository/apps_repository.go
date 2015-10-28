package appsRepository

import (
	"encoding/json"

	"github.com/cloudfoundry/cli/plugin"
)

type AppsModel struct {
	NextUrl   string     `json:"next_url,omitempty"`
	Resources []AppModel `json:"resources"`
}

type AppModel struct {
	Entity EntityModel `json:"entity"`
}

type EntityModel struct {
	Name      string `json:"name"`
	StackGuid string `json:"stack_guid"`
	State     string `json:"state"`
}

func GetAllApps(cliConnection plugin.CliConnection) ([]AppModel, error) {
	response, err := cliConnection.CliCommandWithoutTerminalOutput("curl", "v2/apps")

	apps := AppsModel{}
	err = json.Unmarshal([]byte(response[0]), &apps)

	return apps.Resources, err
}
