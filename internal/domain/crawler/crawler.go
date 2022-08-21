package crawler

import (
	"errors"
	"net/http"
	localE "tag-value-finder/internal/domain/errors"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

// extracts tag value from url
func crawl(url, tag string) (tagValue string, err error) {
	resp, err := http.Get(url) //nolint:gosec,noctx
	if err != nil {
		return "", err
	}

	b := resp.Body
	defer b.Close() // close Body when the function completes

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return "", errors.New(localE.ErrorToken)
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <tag>
			isAnchor := t.Data == tag
			if !isAnchor {
				continue
			}

			// Extract the <tag> value, if there is one
			if tt = z.Next(); tt == html.TextToken {
				return z.Token().Data, nil
			}
		}
	}
}

// GetH1 extracts h1 tag value from url
func GetH1(url string) (h1Value string) {
	v, err := crawl(url, "h1")
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return v
}
