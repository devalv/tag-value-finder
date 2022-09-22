package crawler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	localE "github.com/devalv/tag-value-finder/internal/domain/errors"
)

const mockedH1Value = "BIBA"
const mockedH2Value = "Boba"

func mockH1Endpoint(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<h1>" + mockedH1Value + "</h1>"))
}

func mockH2Endpoint(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<h1><H2>\t" + mockedH2Value + "   </H2></h1>"))
}

func mockedHttpServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimSpace(r.URL.Path) {
		case "/h1":
			mockH1Endpoint(w, r)
		case "/h2":
			mockH2Endpoint(w, r)
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))
}

func TestCrawl(t *testing.T) {
	testSrv := mockedHttpServer()

	cases := []struct {
		url      string
		tag      string
		tagValue string
	}{
		{
			url:      testSrv.URL + "/h1",
			tag:      "h1",
			tagValue: mockedH1Value,
		},
		{
			url:      testSrv.URL + "/h2",
			tag:      "h2",
			tagValue: mockedH2Value,
		},
		{
			url:      testSrv.URL,
			tag:      "h1",
			tagValue: "",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s:%s", c.url, c.tag), func(t *testing.T) {
			tagValue, err := crawl(c.url, c.tag)
			if err != nil && err.Error() != localE.ErrorToken {
				t.Fail()
			}
			if tagValue != c.tagValue {
				t.Fatalf("Expected %s but was `%s`", c.tagValue, tagValue)
			}

		})
	}

}

func TestGetH1(t *testing.T) {
	testSrv := mockedHttpServer()

	cases := []struct {
		url      string
		tag      string
		tagValue string
	}{
		{
			url:      testSrv.URL + "/h1",
			tag:      "h1",
			tagValue: mockedH1Value,
		},
		{
			url:      testSrv.URL,
			tag:      "h1",
			tagValue: "",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s:%s", c.url, c.tag), func(t *testing.T) {
			tagValue := GetH1(c.url)
			if tagValue != c.tagValue {
				t.Fatalf("Expected %s but was `%s`", c.tagValue, tagValue)
			}

		})
	}
}
