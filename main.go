package main

import (
	"log"

	"github.com/monshq/gwam/config"
	"github.com/monshq/gwam/updater"
)

func main() {
	updateAll()
}

func updateAll() {
	config := config.ReadConfig()
	for _, addon := range config.Addons {
		log.Printf("Downloading addon %s from %s", addon.Name, addon.Url)
		updater.Download(addon.Url, addon.Path)
	}
}
