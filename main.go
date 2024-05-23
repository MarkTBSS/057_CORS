package main

import (
	"os"

	_pkgConfig "github.com/MarkTBSS/057_CORS/config"
	_pkgModulesServers "github.com/MarkTBSS/057_CORS/modules/servers"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := _pkgConfig.LoadConfig(envPath())
	_pkgModulesServers.NewServer(cfg).Start()
}
