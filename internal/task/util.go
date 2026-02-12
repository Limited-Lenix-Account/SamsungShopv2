package task

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"samsungshop.go/internal/jwt"
)

func parseXsrfInHtml(body string) (string, string, error) {

	pattern := `'token'\s*:\s*'([^']+)',\s*'headerName'\s*:\s*'([^']+)'`

	// Compile the regex pattern
	re := regexp.MustCompile(pattern)

	// Find the first match
	match := re.FindStringSubmatch(body)

	// fmt.Println(match)
	if len(match) == 3 {
		return match[1], match[2], nil
	} else {
		return "", "", fmt.Errorf("cannot parse csrf header")
	}

}

func parseLoginToken(body *goquery.Document) (*jwt.AccessResp, error) {
	b, ok := body.Find(`input[name='code']`).First().Attr("value")
	if ok {
		fmt.Println("value found!")
	} else {
		return nil, fmt.Errorf("cannot find access struct")
	}
	var access jwt.AccessResp
	err := json.Unmarshal([]byte(b), &access)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling access json: %w", err)
	}
	return &access, nil
}
