package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
)

// request sends out a request and returns the response
func request(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// bytesBufferize expects a JSON-encodeable interface{} and returns a *bytes.Buffer
func bytesBufferize(req interface{}) *bytes.Buffer {
	ret, _ := json.Marshal(req)
	return bytes.NewBuffer(ret)
}

// compareSlice compares two slices; return true if they are equal, false otherwise
func compareSlice(a, b []interface{}) (equal bool) {
	if a == nil && b == nil {
		return true
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// compareMap compares two maps, naively
func compareMap(a, b map[string]interface{}) (equal bool) {
	for key, value := range a {
		if val, ok := b[key]; !ok || !reflect.DeepEqual(val, value) {
			return false
		}
	}
	for key, value := range b {
		if val, ok := a[key]; !ok || !reflect.DeepEqual(val, value) {
			return false
		}
	}
	return true
}
