package keen_test

import (
	"testing"

	"github.com/WatchBeam/keen"
	"github.com/WatchBeam/keen/record"
	"github.com/stretchr/testify/assert"
)

func TestRecordReturnsARecorder(t *testing.T) {
	r := keen.Recorder("project_id", "write_key")

	assert.IsType(t, &record.KeenRecorder{}, r)
}
