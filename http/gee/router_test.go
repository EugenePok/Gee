package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestPasePattern(t *testing.T) {
	var tests = []struct {
		pattern string
		result  []string
	}{
		{"/p/:name", []string{"p", ":name"}},
		{"/p/*", []string{"p", "*"}},
		{"/p/*name/*", []string{"p", "*name"}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test %d", i)
		t.Run(testName, func(t *testing.T) {
			got := parsePattern(test.pattern)
			want := test.result
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %s want %s", got, want)
			}
		})
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()

	var tests = []struct {
		method  string
		path    string
		pattern string
		key     string
		value   string
	}{
		{"GET", "/", "/", "", ""},
		{"GET", "/hello/eugene", "/hello/:name", "name", "eugene"},
		{"GET", "/hello/b/c", "/hello/b/c", "", ""},
		{"GET", "/hi/eugene", "/hi/:name", "name", "eugene"},
		{"GET", "/assets/myFile.pdf", "/assets/*filepath", "", ""},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test %d", i)
		t.Run(testName, func(t *testing.T) {
			gotN, gotParam := r.getRoute(test.method, test.path)
			if gotN == nil {
				t.Errorf("nil shouldn't be returned")
			}
			if gotN.pattern != test.pattern {
				t.Errorf("should match %s", test.pattern)
			}
			if test.key != "" && test.value != "" {
				val, ok := gotParam[test.key]
				if !ok {
					t.Errorf("cannot get %s key from params", test.key)
				}
				if val != test.value {
					t.Errorf("want $%s, got %s value with key %s from params", val, test.value, test.key)
				}
			}
		})
	}
}
