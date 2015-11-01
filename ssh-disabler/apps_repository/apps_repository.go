package appsRepository

import (
	"encoding/json"
	"errors"

	"github.com/cloudfoundry/cli/plugin"
)

type AppModel struct {
	Entity EntityModel `json:"entity"`
}

type EntityModel struct {
	Name      string `json:"name"`
	EnableSSH bool   `json:"enable_ssh"`
}

func DisableAppSSH(cliConnection plugin.CliConnection, appGuid string) error {
	response, err := cliConnection.CliCommandWithoutTerminalOutput("curl", "v2/apps/"+appGuid, "-X", "PUT", "-d", `{"enable_ssh":false}`)
	if err != nil {
		return errors.New("Error curling v2/apps endpoint: " + err.Error())
	}

	app := AppModel{}
	err = json.Unmarshal([]byte(response[0]), &app)

	if !app.Entity.EnableSSH {
		return errors.New("Failed to disable SSH for application '" + app.Entity.Name + "'")
	}

	return nil
}
