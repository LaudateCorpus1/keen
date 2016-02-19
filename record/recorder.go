package record

// A Recorder represents an interface to a concrete implementation capable of
// recording events into a Keen.io datastore.
type Recorder interface {
	// Record records an event to the Keen datastore, logging any error that
	// occurs along the way. All errors should be returned immediately,
	// stopping execution of the method.
	Record(e Event) error
}
