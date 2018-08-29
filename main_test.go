package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var helloendpoints = map[string]string{
	"/hello":                 "hello",
	"/hello?uppercase=true":  "HELLO",
	"/hello?uppercase=false": "hello",
	"/hello?reverse=true":    "olleh",
	"/hello?reverse=false":   "hello",
}

func TestHellohandler(t *testing.T) {
	for key, value := range helloendpoints {
		req, err := http.NewRequest("GET", key, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Hellohandler)
		handler.ServeHTTP(rr, req)

		expected := value
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}

var hwendpoints = map[string]string{
	"/":                 "hello world",
	"/?uppercase=true":  "HELLO WORLD",
	"/?uppercase=false": "hello world",
	"/?reverse=true":    "dlrow olleh",
	"/?reverse=false":   "hello world",
}

func TestHWhandler(t *testing.T) {
	for key, value := range hwendpoints {
		req, err := http.NewRequest("GET", key, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HWhandler)
		handler.ServeHTTP(rr, req)

		expected := value
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}

var worldendpoints = map[string]string{
	"/world":                 "world",
	"/world?uppercase=true":  "WORLD",
	"/world?uppercase=false": "world",
	"/world?reverse=true":    "dlrow",
	"/world?reverse=false":   "world",
}

func TestWorldhandler(t *testing.T) {
	for key, value := range worldendpoints {
		req, err := http.NewRequest("GET", key, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Worldhandler)
		handler.ServeHTTP(rr, req)

		expected := value
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}
