package env

import "github.com/chubin518/kestrel-layout-basic/buildinfo"

const (
	DEVELOPMENT = "dev"
	TESTING     = "test"
	STAGING     = "stage"
	PRODUCTION  = "prod"
)

var active *environment

type environment struct {
	value string
}

func init() {
	active = &environment{
		value: buildinfo.Environment,
	}
}

func Active() string {
	return active.value
}

func IsDevelopment() bool {
	return Active() == DEVELOPMENT
}

func IsTesting() bool {
	return Active() == TESTING
}

func IsStaging() bool {
	return Active() == STAGING
}

func IsProduction() bool {
	return Active() == PRODUCTION
}
