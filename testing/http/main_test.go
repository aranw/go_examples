package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"testing/quick"
)

func TestAdd(t *testing.T) {
	assertion := func(a, b int) bool {
		return Add(a, b) == Add(b, a)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

func TestAddToZero(t *testing.T) {
	assertion := func(a int) bool {
		return Add(a, 0) == Add(0, a)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/add?x=2&y=3", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	z, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("expected an integer; got %s", b)
	}
	if z != 5 {
		t.Fatalf("expected z to be 5; got %v", z)
	}
}

func TestAddHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/?x=1&y=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()

	addHandler(rw, req)
	if rw.Code == 500 {
		t.Fatal("Internal server Error: " + rw.Body.String())
	}
	if rw.Body.String() != "3" {
		t.Fatal("Expected " + rw.Body.String())
	}
}

func TestAddHandlerXErrors(t *testing.T) {
	req, err := http.NewRequest("GET", "/?x=1aa&y=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()

	addHandler(rw, req)
	if rw.Code != 500 {
		t.Fatal("Expected add handler to return error")
	}
}

func TestAddHandlerYErrors(t *testing.T) {
	req, err := http.NewRequest("GET", "/?x=1&y=2b", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()

	addHandler(rw, req)
	if rw.Code != 500 {
		t.Fatal("Expected add handler to return error")
	}
}
