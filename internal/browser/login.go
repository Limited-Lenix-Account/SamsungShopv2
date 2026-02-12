package browser

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"

	http "github.com/bogdanfinn/fhttp"
)

func BrowserLogin(r *rod.Browser, username, password string) ([]*http.Cookie, error) {

	r.IgnoreCertErrors(true)
	page := r.MustPage("")
	page.MustWaitDOMStable()

	// router := page.HijackRequests()
	// router.MustAdd("*", handleLogin)
	// go router.Run()

	page.MustNavigate("https://samsung.com/us/").MustWaitLoad()

	time.Sleep(1 * time.Second)
	page.MustClose()

	c, _ := r.GetCookies()

	var httpCookies []*http.Cookie

	for _, v := range c {

		if strings.Contains(v.Domain, "samsung") {
			c := http.Cookie{
				Name:  v.Name,
				Value: v.Value,
			}
			httpCookies = append(httpCookies, &c)

		}
	}

	return httpCookies, nil
}

func handleLogin(h *rod.Hijack) {

	fmt.Println(h.Request.URL())
	h.Request.Req().Header.Del("User-Agent")
	h.MustLoadResponse()

	h.ContinueRequest(&proto.FetchContinueRequest{
		RequestID: h.Response.Payload().RequestID,
	})

}
