package record

// An Event represents an item that can be collected into an Event Collection by
// Keen.io. It belongs to a collection, which is represented by the return value
// of the Collection method.
type Event interface {
	// Collection returns the name of the event collection that events of
	// this particular type should belong to.
	//
	// For most implementations, it makes most sense to return a constant,
	// making this a pure-function.
	Collection() string
}
