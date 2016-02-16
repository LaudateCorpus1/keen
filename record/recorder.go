package record

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// A Recorder is responsible for recording events into Keen.io's data collection
// service. It uses the given projectId and writeKey, meaning that a Recorder is
// scoped to the project context. All methods on it are blocking, so it may be
// best to spawn go-routines for high-throughput event generators.
type Recorder struct {
	// projectId and writeKey are used for communicating to Keen.io's API
	// which project to write to.
	projectId, writeKey string
	// baseUrl is the base URL of the Keen API.
	baseUrl *url.URL
	// httpClient is responsible for round-tripping the http requests to and
	// from Keen.io's API.
	httpClient http.Client
}

// New returns a pointer to a new instance of a Recorder. It passes along the
// given projectId and writeKey, which will be used in all future requests.
func New(baseUrl *url.URL, projectId, writeKey string) *Recorder {
	return &Recorder{
		baseUrl:   baseUrl,
		projectId: projectId,
		writeKey:  writeKey,
	}
}

// Record records a particular event into Keen.io's data-collection library
// using the "events" API. If any error is encountered within the runtime, it
// will be returned immediately, and execution will be halted. If the request
// returns an error-ful response code, then it will be turned into an error, and
// returned as well. If the returned error is nil, then the event posting has
// gone successfully.
//
// The event passed to this method is stored in an event "collection" with the
// name corresponding to the return value of Event.Collection().
//
// Note: since this method is blocking, it may be best to run it within its own
// goroutine, but this is not necessary.
func (r *Recorder) Record(e Event) error {
	resp, err := r.request(e)
	if err != nil {
		return err
	}

	return r.responseToError(resp)
}

// responseToError converts non 200-level response codes into errors, containing
// the body of the response.
func (r *Recorder) responseToError(resp *http.Response) error {
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return fmt.Errorf("keen: errorful HTTP response: %s", data)
}

// request prepares an API request to post an event to Keen's API. It marshals
// the event interface into JSON, and sends that as the body of the request.
func (r *Recorder) request(e Event) (*http.Response, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	url, err := r.url(e)
	if err != nil {
		return nil, err
	}

	return r.httpClient.Post(url.String(), "application/json", bytes.NewReader(body))
}

// url returns the appropriate url to use for the request, filling in the
// details pertaining to the project name and the event collection. It also
// authenticates the request.
func (r *Recorder) url(e Event) (*url.URL, error) {
	url, err := url.Parse(
		fmt.Sprintf("/3.0/projects/%s/events/%s?api_key=%s",
			r.projectId, e.Collection(), r.writeKey))
	if err != nil {
		return nil, err
	}

	return r.baseUrl.ResolveReference(url), nil
}
