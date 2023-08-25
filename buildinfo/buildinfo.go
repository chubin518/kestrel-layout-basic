package buildinfo

import (
	"os"
	"path"
)

var (
	// build version
	Version string = "0.0.1-beta1"
	// build name
	Name string = path.Base(os.Args[0])
	// build env: dev/test/stage/prod
	Environment string = "dev"
	// build time
	BuildTime string
	// build go version
	BuildVersion string
)
