package main

import "github.com/cloudfoundry/cli/plugin"

func main() {
	appLister := &AppLister{}
	plugin.Start(appLister)
}
