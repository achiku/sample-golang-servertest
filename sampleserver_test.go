package sampleserver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSampleServer(t *testing.T) {
	mux := CreateMux()
	ts := httptest.NewServer(mux)
	defer ts.Close()

	urls := map[string]string{
		"/hello": "hello, world!\n",
		"/bye":   "Good bye, this world..\n",
	}

	for u, expected := range urls {
		t.Logf("url: %s%s", ts.URL, u)
		resp, err := http.Get(ts.URL + u)
		if err != nil {
			t.Error(err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()
		t.Log(string(b))
		if string(b) != expected {
			t.Errorf("want %s got %s", expected, string(b))
		}
	}
}

func TestRawHelloHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(hello))
	defer ts.Close()

	t.Logf("url: %s", ts.URL)
	resp, err := http.Get(ts.URL + "/hello")
	if err != nil {
		t.Error(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	t.Log(string(b))
}

func TestRawByeHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(bye))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/bye")
	if err != nil {
		t.Error(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	t.Log(string(b))
}
