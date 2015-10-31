package appsRepository

import (
	"encoding/json"
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/cloudfoundry/cli/plugin/models"
)

type AppModel struct {
	Entity EntityModel `json:"entity"`
}

type EntityModel struct {
	Name      string `json:"name"`
	StackGuid string `json:"stack_guid"`
	EnableSSH bool   `json:"enable_ssh"`
}

func DisableAppsSSH(cliConnection plugin.CliConnection, apps []plugin_models.GetAppsModel) error {
	for _, app := range apps {
		response, err := cliConnection.CliCommandWithoutTerminalOutput("curl", "v2/apps/"+app.Guid, "-X", "PUT", "-d", `{"enable_ssh":false}`)
		if err != nil {
			fmt.Println("Error curling v2/apps endpoint")
		}

		app := AppModel{}
		err = json.Unmarshal([]byte(response[0]), &app)

		fmt.Println("enable_ssh", app.Entity.Name, app.Entity.EnableSSH)
	}

	return nil
}
