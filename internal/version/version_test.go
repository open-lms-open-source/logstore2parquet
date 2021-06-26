package version

import (
	"testing"
)

func TestString(t *testing.T) {
	App = "MyApp"
	Version = "v1.0.0"
	BuildDate = "1970-01-01T09:30:00+0930"

	want := "MyApp v1.0.0, built 1970-01-01T09:30:00+0930, "
	got := String()
	if got != want {
		t.Errorf("version.String() got '%s', want '%s'", got, want)
	}
}

func TestJSON(t *testing.T) {
	App = "MyApp"
	Version = "v1.0.0"
	BuildDate = "1970-01-01T09:30:00+0930"

	want := `{"app":"MyApp","version":"v1.0.0","build_date":"1970-01-01T09:30:00+0930","go_version":""}`
	got := JSON()
	if got != want {
		t.Errorf("version.JSON() got '%s', want '%s'", got, want)
	}
}
