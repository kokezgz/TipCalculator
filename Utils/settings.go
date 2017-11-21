package Utils

import (
	"encoding/json"
	"io/ioutil"
)

var setsFile = "appsettings.json"

func NewSettings(ID string) Config {
	sets := mapSettings(setsFile)
	return sets[ID]
}

func parseSettings(settingsFile string) []Config {

	f, err := ioutil.ReadFile(settingsFile)
	if err != nil {
		println(err.Error())
	}

	var configs []Config
	err = json.Unmarshal(f, &configs)

	if err != nil {
		println(err.Error())
	}

	return configs
}

func mapSettings(settingFile string) map[string]Config {

	configs := parseSettings(settingFile)
	mapSettings := make(map[string]Config)
	for _, s := range configs {
		mapSettings[s.ID] = s
	}

	return mapSettings
}
