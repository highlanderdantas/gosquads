package conf

import (
	"log"

	"gopkg.in/ini.v1"
	"gopkg.in/macaron.v1"
)

//Cfg represents the pointer to configuration file
var Cfg *ini.File

// find configuration file
func init() {
	var err error
	Cfg, err = macaron.SetConfig("conf/app.ini")
	if err != nil {
		log.Fatalf("[conf/Init] Error during app.ini reading. Error: %s\n", err.Error())
	}
}
