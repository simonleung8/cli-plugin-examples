package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type TokenRefresher struct{}

func (c *TokenRefresher) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "TokenRefresher",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "show-token",
				HelpText: "call CLI to refresh auth token",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "cf show-token",
				},
			},
		},
	}
}
func (c *TokenRefresher) Run(cliConnection plugin.CliConnection, args []string) {
	token, err := cliConnection.AccessToken()
	if err != nil {
		fmt.Println("Error refreshing token:", err)
	}

	fmt.Println("Otained refreshed token from CLI:")
	fmt.Println("")
	fmt.Println(token)
}

func main() {
	plugin.Start(new(TokenRefresher))
}
