// main_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig(t *testing.T) {
	_, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		t.Errorf("ERROR: Test - Missing config file")
	}
}

func TestSingleResult(t *testing.T) {
	t.Run("Must return single result", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/zipcode?key=alcoy", nil)
		response := httptest.NewRecorder()

		handleZipcode(response, request)

		got := response.Body.String()
		want := `[{"zipcode":"6023","area":"Alcoy","provinceCity":"Cebu"}]`

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})
}
