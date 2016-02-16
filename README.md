# keen

[![Build Status](https://travis-ci.org/WatchBeam/keen.svg?branch=master)](https://travis-ci.org/WatchBeam/keen)
[![GoDoc](https://godoc.org/github.com/WatchBeam/keen?status.svg)](https://godoc.org/github.com/WatchBeam/keen)

Keen is a simple Golang API for communicating with the Keen.io datastorage
service. It implements only one function of the API, but this should be expected
to change in the future.

## Example

```go
import (
        "github.com/WatchBeam/keen"
        "github.com/WatchBeam/keen/record"
)

type SimpleEvent struct {
        Foo string `json:"foo"`
}

var _ record.Event = new(SimpleEvent)

func (e *SimpleEvent) Collection string { return "collection" }

recorder := keen.Recorder("project_id", "write_key")
err := recorder.Record(&SimpleEvent{
        Foo: "bar",
})
```

## License

MIT.
