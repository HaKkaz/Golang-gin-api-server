package test

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	// Setup the test data
	testData := []struct {
		path     string
		method   string
		body     string
		status   int
		expected string
	}{
		{
			"/api/v1/ad",
			http.MethodPost,
			`{"title": "General Case 1","startAt": "2023-12-10T03:00:00.000Z","endAt": "2025-12-31T16:00:00.000Z","conditions": {"ageStart": 20,"ageEnd": 30,"country": ["TW", "JP"],"platform": ["android", "ios"]}}`,
			200,
			`{"message":"Advertisement created successfully"}`,
		},
	}

	// Check if the api server is up
	_, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Logf("Server not up yet... try again")
		os.Exit(0)
	}

	for _, tc := range testData {
		bodyVal := tc.body
		if bodyVal == "" {
			bodyVal = "nil"
		}
		t.Logf("Sending request... Path: \"%s\"\tMethod: %s\tBody: %s", tc.path, tc.method, bodyVal)

		req, err := http.NewRequest(tc.method, "http://localhost:8080"+tc.path, bytes.NewReader([]byte(tc.body)))
		if err != nil {
			t.Fatal(err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
			return
		}

		if !assert.Equal(t, tc.status, resp.StatusCode) {
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
			return
		}
		if !assert.Equal(t, tc.expected, string(body)) {
			return
		}

		err = resp.Body.Close()
		if err != nil {
			t.Fatal(err)
			return
		}
	}
}
