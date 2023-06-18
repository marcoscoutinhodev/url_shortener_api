package adapter

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

type URLCheckerAdapter struct{}

func NewURLCheckerAdapter() *URLCheckerAdapter {
	return &URLCheckerAdapter{}
}

func (u URLCheckerAdapter) IsURLSafe(ctx context.Context, url string) bool {
	urlEncoded := base64.RawURLEncoding.EncodeToString([]byte(url))
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf("https://www.virustotal.com/api/v3/urls/%s", urlEncoded),
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header = map[string][]string{
		"Accept":   {"application/json"},
		"x-apikey": {os.Getenv("URL_CHECKER_API_KEY")},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return res.StatusCode == 200
}
