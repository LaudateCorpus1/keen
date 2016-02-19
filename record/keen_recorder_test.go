package record_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/WatchBeam/keen/record"
	"github.com/stretchr/testify/assert"
)

func TestNewReturnsRecorder(t *testing.T) {
	url, _ := url.Parse("https://api.keen.io/")
	r := record.New(url, "project_id", "write_key")

	assert.IsType(t, &record.KeenRecorder{}, r)
}

func TestRecordMakesSuccessfulRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/3.0/projects/1/events/foo?api_key=bar", r.URL.String())
		assert.Equal(t, []byte("{\"woot\":6}"), payload)

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	url, _ := url.Parse(server.URL)
	recorder := record.New(url, "1", "bar")

	err := recorder.Record(&StubEvent{Woot: 6})

	assert.Nil(t, err)
}

func TestRecordPropogatesNon200Errors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}))
	defer server.Close()

	url, _ := url.Parse(server.URL)
	recorder := record.New(url, "1", "bar")

	err := recorder.Record(&StubEvent{Woot: 6})

	assert.Equal(t, "keen: errorful HTTP response: bad request", err.Error())
}
