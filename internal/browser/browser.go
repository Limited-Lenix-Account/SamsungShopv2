package browser

import (
	"log"

	http "github.com/bogdanfinn/fhttp"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
)

func MakeBrowser(username, password string) ([]*http.Cookie, error) {

	pool := rod.NewBrowserPool(1)

	u := launcher.NewUserMode().
		Set("disable-default-apps").
		Set("no-first-run").
		// Set("ignore-certificate-errors", "").
		// Set(flags.Headless).MustLaunch()
		Headless(false).
		MustLaunch()

	bPool, err := pool.Get(func() (*rod.Browser, error) {
		browser := rod.New().ControlURL(u).MustConnect().DefaultDevice(devices.IPadPro)
		return browser, nil
	})

	if err != nil {
		log.Fatalf("error creating pool")
	}

	cookies, err := BrowserLogin(bPool, username, password)
	if err != nil {
		return nil, err
	}

	return cookies, nil

}
