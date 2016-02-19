package keen

import (
	"net/url"

	"github.com/WatchBeam/keen/record"
)

// Recorder returns a new Recorder used for recording events.
//
// More information can be found in the package github.com/WatchBeam/keen/record
func Recorder(projectId, writeKey string) record.Recorder {
	url, _ := url.Parse("https://api.keen.io/")
	return record.New(url, projectId, writeKey)
}
