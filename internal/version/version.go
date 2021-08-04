package version

import (
	"encoding/json"
	"fmt"
)

// Version and BuildDate contain build information populated by the
// compiler.
var (
	App       string
	Version   string
	BuildDate string
)

type versionInfo struct {
	App       string `json:"app"`
	Version   string `json:"version"`
	BuildDate string `json:"build_date"`
}

// String renders version information as a string.
func String() string {
	return fmt.Sprintf("%s %s, built %s", App, Version, BuildDate)
}

// JSON renders version information in JSON format.
func JSON() string {
	v, _ := json.Marshal(versionInfo{App, Version, BuildDate})
	return string(v)
}
