package http2

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealth(t *testing.T) {
	_ = HealthCheckHandler(context.TODO(), ":3000")
	//Time to run-up http server
	time.Sleep(1 * time.Second)

	cases := []struct {
		url     string
		expCode int
		expBody string
	}{
		{
			url:     "/health",
			expCode: http.StatusOK,
			expBody: "ok",
		},
		{
			url:     "/bad-health-url",
			expCode: http.StatusNotFound,
			expBody: "404 page not found\n",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s:%d", c.url, c.expCode), func(t *testing.T) {
			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, c.url, nil)
			http.DefaultServeMux.ServeHTTP(resp, req)

			r := resp.Result()
			if r.StatusCode != c.expCode {
				t.Fatalf("Expected %d, but was %d", c.expCode, r.StatusCode)
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fail()
			}

			if string(body) != c.expBody {
				t.Fatalf("Expected %s but was `%s`", c.expBody, body)
			}

		})
	}

}
