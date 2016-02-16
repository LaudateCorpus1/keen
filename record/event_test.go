package record_test

import "github.com/WatchBeam/keen/record"

type StubEvent struct {
	Woot int `json:"woot"`
}

var _ record.Event = new(StubEvent)

func (s *StubEvent) Collection() string { return "foo" }
