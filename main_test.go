package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {
	now := func() time.Time {
		return time.Date(2011, 3, 4, 10, 20, 0, 0, time.UTC)
	}
	s := httptest.NewServer(http.HandlerFunc(index(now)))
	defer s.Close()
	rsp, err := http.Get(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()
	b, _ := ioutil.ReadAll(rsp.Body)
	if rsp.Status[0] != '2' {
		t.Fatalf("got status %s but expected 2x. body=%s", rsp.Status, string(b))
	}
	tests := []struct{ Has, Want interface{} }{
		{string(b), "{\"current_time\":\"2011-03-04T10:20:00Z\"}\n"},
	}
	for i, tc := range tests {
		if tc.Has != tc.Want {
			t.Errorf("%d: want=%#v has=%#v", i+1, tc.Want, tc.Has)
		}
	}
}
