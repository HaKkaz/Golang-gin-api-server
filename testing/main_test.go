package test

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	// Setup the test data
	testData := initTestCases()

	// loop 5 times to wait for the server to start
	for i := 0; i < 3; i++ {
		// Check if the api server is up
		_, err := http.Get("http://localhost:8080")
		if err != nil {
			if i == 4 {
				t.Logf("Connection to server failed... exiting test...")
				os.Exit(0)
			} else {
				t.Logf("Server not up yet... try again")
				time.Sleep(3 * time.Second)
			}
		}

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
