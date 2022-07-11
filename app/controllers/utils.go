package controllers

import (
	"todoapi/config"
)

var deployEnv = config.Config.Deploy
var serverPort = config.Config.Port
