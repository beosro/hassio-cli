package command

import (
	"fmt"
	"os"

	"github.com/home-assistant/hassio-cli/command/helpers"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// CmdHost All host endpoints for hass.io
func CmdHost(c *cli.Context) {
	const HassioBasePath = "host"
	action := ""
	endpoint := ""
	serverOverride := ""
	get := false
	Options := c.String("options")
	RawJSON := c.Bool("rawjson")
	Filter := c.String("filter")
	if c.NArg() > 0 {
		action = c.Args()[0]
	}

	switch action {
	case "info": // GET
		get = true
		endpoint = action
	case "reboot", // POST
		"shutdown",
		"options":
		endpoint = action
	default:
		fmt.Fprintf(os.Stderr, "No valid action detected.\n")
		os.Exit(3)
	}

	log.WithFields(log.Fields{
		"action":         action,
		"endpoint":       endpoint,
		"serveroverride": serverOverride,
		"get":            get,
		"options":        Options,
		"rawjson":        RawJSON,
		"filter":         Filter,
	}).Debug("[CmdHost]")

	if endpoint != "" {
		helpers.ExecCommand(HassioBasePath, endpoint, serverOverride, get, Options, Filter, RawJSON)
	}
}
