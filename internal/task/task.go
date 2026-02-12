package task

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"samsungshop.go/internal/captcha"
	"samsungshop.go/internal/database"
	"samsungshop.go/internal/enc"
	"samsungshop.go/internal/jwt"
)

// TODO: make it so only app login is a thing, there's no reason to do browser login... like 0

type Task struct {
	DB         *database.DB
	Email      string
	Password   string
	Client     tls_client.HttpClient
	CsrfToken  string
	CsrfHeader string

	ReCaptchaToken  string
	RecaptchaAction string
	ReCaptchaValue  string
	AssessmentId    string
	Url             *url.URL
	ReDirURL        string

	PromoJWT     string
	CouponCode   string
	EncPayload   BillingPayload
	ShippingInfo ShippingPayload
	CardInfo     PaymentInfo

	UpdateTerms bool

	AppLogin     bool
	Login        bool
	LoginJWT     string
	BrowserLogin bool

	AuthStruct    *jwt.AccessResp
	CartID        string
	DeliveryGroup string
	ExchangeID    string
	EncLogin      map[string]string

	Log *logrus.Entry
}

func (t *Task) Start(log *logrus.Logger) {
	log.Info("Bot Started")

	if err := t.Initalize(log); err != nil {
		fmt.Printf("error intializing task: %s\n", err)
	}

	// getUser() will determine if a login is needed, if a accounts' jwt or access token is empty:
	// t.Login = true
	if err := t.getUser(); err != nil {
		fmt.Printf("error getting user: %s\n", err)
	}

	if t.Login {
		t.Log.Info("logging in...")
		// gets xsrf token
		if err := t.getHome(); err != nil {
			fmt.Printf("error getting homepage: %s\n", err)
		}

		// sends login -> gets access token struct n stuff
		if err := t.sendLogin(); err != nil {
			fmt.Printf("error logging in: %s\n", err)
		}

		if t.UpdateTerms {
			if err := t.updateTerms(); err != nil {
				fmt.Printf("error updating terms: %s\n", err)
			}
		}

		if err := t.completeSignIn(); err != nil {
			fmt.Printf("error completing sign-in: %s\n", err)
		}

		if err := t.getJWT(); err != nil {
			fmt.Printf("error getting jwt: %s\n", err)
		}

	}

	// if err := t.getExchange(); err != nil {
	// 	fmt.Printf("error getting cart: %s\n", err)
	// }

	if err := t.addToCart(); err != nil {
		fmt.Printf("error getting cart: %s\n", err)
	}

	if err := t.getCart(); err != nil {
		log.Fatal(err)
	}

	if err := t.addShipping(); err != nil {
		fmt.Printf("error adding shipping: %s\n", err)
	}

	if err := t.addBilling(); err != nil {
		fmt.Printf("error adding billing: %s\n", err)
	}

	// if err := t.applyCoupon(); err != nil {
	// 	fmt.Printf("error adding coupon code: %s\n", err)
	// }

	if err := t.getShippingRates(); err != nil {
		fmt.Printf("error getting shipping rate: %s\n", err)
	}

	if err := t.setShippingRate(); err != nil {
		fmt.Printf("error setting shipping rate: %s\n", err)
	}

	if err := t.submitOrder(); err != nil {
		fmt.Printf("error submitting order: %s\n", err)
	}

}

func (t *Task) Initalize(log *logrus.Logger) error {
	t.Email = "esliva+1@mines.edu"
	t.Password = "L3nixLovesYou!"
	t.CouponCode = "ref-fi8hum"
	t.Log = log.WithField("email", t.Email)
	t.Log.Info("Initalizing...")

	// t.CartID = "616a9c47-95b3-4e34-9b5e-a7159269dd73"

	p, err := t.DB.GetProfile(t.Email)
	if err != nil {
		t.Log.Fatalf("cannot get profile: %s", err)
	}

	s := ShippingPayload{
		Email:      *p[0],
		FirstName:  *p[1],
		LastName:   *p[2],
		Phone:      *p[3],
		Line1:      *p[4],
		Line2:      *p[5],
		City:       *p[6],
		State:      *p[7],
		PostalCode: *p[8],
		Country:    "US",
	}

	c := PaymentInfo{
		Number:   *p[9],
		ExpMonth: *p[10],
		ExpYear:  *p[11],
		CVV:      *p[12],
	}

	t.ShippingInfo = s
	t.CardInfo = c

	// get login JWT from database
	jwt, err := t.DB.GetUserJWT(t.Email)
	if err != nil {
		return err
	}
	t.LoginJWT = jwt

	jar := tls_client.NewCookieJar()
	opts := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.NikeIosMobile),
		tls_client.WithCookieJar(jar),
		tls_client.WithNotFollowRedirects(),
		// tls_client.WithDebug(),
	}

	u, _ := url.Parse("https://account.samsung.com")
	t.Url = u

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), opts...)
	if err != nil {
		return fmt.Errorf("error creating session: %w", err)
	}
	t.Client = client

	return nil
}

// Main Function: get xsrf token and xsrf header name
//
// Returns: sets tasks' xsrf header and value
func (t *Task) getHome() error {

	glb := "GLBwx1hzy719qa"
	var c []*http.Cookie
	glbcookie := http.Cookie{
		Name:   "glbState",
		Value:  glb,
		Domain: "samsung.com",
	}

	c = append(c, &glbcookie)

	// GLB8xgqid5evht
	// GLB3tqgy9weo3p

	url := ("https://account.samsung.com/accounts/eCommerce_US/signInGate?clientId=gx9zz84e3x&countryCode=US&tokenType=TOKEN&deviceType=APP&deviceUniqueID=A029FA4D-35E6-4149-96E6-F92BF137BFB7&devicePhysicalAddressText=A029FA4D-35E6-4149-96E6-F92BF137BFB7&deviceOSVersion=18.2.1&competitorDeviceYNFlag=N&locale=en&redirect_uri=https://ecommapi.ecom-mobile-samsung.com/sso/callback&state=58694312-c822-429e-a1e4-a868342eda34&")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	req.Header = map[string][]string{
		"User-Agent": {"Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"},
		"Host":       {"account.samsung.com"},
		"Accept":     {"application/json, text/plain, */*"},
		"Referer":    {"https://www.samsung.com/"},
	}

	t.Client.SetCookies(t.Url, c)
	resp, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err)
	}

	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	t.Client.SetCookies(t.Url, resp.Cookies())

	// XSRF is parsed from the login page in the app, not the xhr endpoint

	val, key, err := parseXsrfInHtml(string(body))
	if err != nil {
		return fmt.Errorf("cannot parse xsrf")
	}

	fmt.Printf("csrf key: %s\n", key)
	fmt.Printf("csrf val: %s\n", val)

	t.CsrfHeader = key
	t.CsrfToken = val

	t.Log.Info("Created Session")

	return nil

}

// Main function: solves captcha and sends login information - Returns JSON indicating if credentials were valid
//
// Returns: Sets UpdateTerms flag
func (t *Task) sendLogin() error {

	c, err := captcha.SolveV2()
	if err != nil {
		return fmt.Errorf("cannot solve captcha")
	}

	l := LoginPayload{
		Email:    t.Email,
		Password: t.Password,
		Captcha:  c["gRecaptchaResponse"],
	}

	lPayload, err := json.Marshal(l)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://account.samsung.com/accounts/eCommerce_US/signInProc?v=1740765221436", strings.NewReader(string(lPayload)))
	if err != nil {
		return err
	}

	req.Header = map[string][]string{
		"Host":       {"account.samsung.com"},
		t.CsrfHeader: {t.CsrfToken},
		// "x-recaptcha-token": {t.ReCaptchaToken},
		"User-Agent":   {"Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"},
		"Referer":      {"https://account.samsung.com/accounts/eCommerce_US/signInPassword"},
		"Content-Type": {"application/json; charset=UTF-8"},
	}

	resp, err := t.Client.Do(req)
	if err != nil {
		return (err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err)
	}

	var loginResp EmailResp

	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return err
	}

	if loginResp.RtnCd == "FAILED" {
		return fmt.Errorf("login failed: %s", loginResp.RtnMsg)
	}

	t.Log.Info("Logged In")

	if loginResp.RtnCd == "REQUIRE_TERMS_UPDATE" {
		t.UpdateTerms = true
	}

	var cookies []*http.Cookie

	clientCookies := t.Client.GetCookies(t.Url)
	cookies = append(cookies, clientCookies...)
	// for _, v := range cookies {
	// 	fmt.Println(v.Name, v.Value)
	// }
	t.Client.SetCookies(t.Url, cookies)

	return nil
}

// Main Function: Sends request to update terms on an account if needed
//
// Returns: Nothing
func (t *Task) updateTerms() error {
	t.Log.Info("Updating terms")
	var data = strings.NewReader(`{"tncAccepted":"Y","privacyAccepted":"Y","cscAccepted":"Y","customizationServiceAccepted":"N","newsAndSpecialOffersAccepted":"N","personalizedAdsSADataAccepted":"N","personalizedAdsLocDataAccepted":"N","financialIncentivesAccepted":"Y"}`)
	req, err := http.NewRequest("POST", "https://account.samsung.com/accounts/v1/samsung_com_us/termsProc?v=1739232480520", data)
	if err != nil {
		return (err)
	}

	req.Header = map[string][]string{
		"Host":              {"account.samsung.com"},
		t.CsrfHeader:        {t.CsrfToken},
		"x-recaptcha-token": {t.ReCaptchaToken},
		"User-Agent":        {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36"},
		"Referer":           {"https://account.samsung.com/accounts/v1/samsung_com_us/signIn"},
		"Content-Type":      {"application/json; charset=UTF-8"},
	}

	resp, err := t.Client.Do(req)
	if err != nil {
		return (err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err)
	}

	var loginResp EmailResp
	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return err
	}

	return nil
}

// Main Function: GETs signInComplete page and parses access token struct if successful
//
// Returns: Sets the task's AuthStruct object
func (t *Task) completeSignIn() error {

	t.Log.Info("Completing Signin...")

	// desktop endpoint https://account.samsung.com/accounts/v1/samsung_com_us/signInComplete
	req, err := http.NewRequest(http.MethodGet, "https://account.samsung.com/accounts/eCommerce_US/signInComplete", nil)
	if err != nil {
		return err
	}

	req.Header = map[string][]string{
		"Host":       {"www.samsung.com"},
		t.CsrfHeader: {t.CsrfToken},
		// "x-recaptcha-token": {t.ReCaptchaToken},
		"User-Agent": {"Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"},
		"Referer":    {"https://account.samsung.com/accounts/eCommerce_US/signInPassword"},
	}

	resp, err := t.Client.Do(req)
	if err != nil {
		return (err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot parse login body: %w", err)
	}

	access, err := parseLoginToken(doc)
	if err != nil {
		return fmt.Errorf("cannot find access token: %w", err)
	}

	t.AuthStruct = access

	resp.Body.Close()
	return nil
}

// will implement later to use jwt token instead of v1
// will be used to determine if a re-login is needed
func (t *Task) getUser() error {

	if t.LoginJWT == "" {
		t.Login = true
		t.Log.Error("login missing")
		return nil
	}

	u := make(map[string]string)
	u["jwt"] = t.LoginJWT

	uPayload, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error marshalling json: %w", err)
	}

	var data = strings.NewReader(string(uPayload))
	req, err := http.NewRequest("POST", "https://www.samsung.com/us/api/v1/sso/jwt/details", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("x-client-request-id", "pwa_common_919caba1-8fb2-4a77-b2fd-3c8e28d08874")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/cart/?appId=samsung-mobile-app-ios")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Login = true
		return fmt.Errorf("invalid status code: %w", err)
	}

	var user SamsungUser
	err = json.Unmarshal(body, &user)
	// if account is invalid relogin i guess
	if user.UserInfo.Firstname == "" {
		t.Login = true
		return fmt.Errorf("missing name on account")
	}

	t.Log.Info("Logged In!")
	t.Login = false
	return nil
}

// Lowkey the Most important part of the login
//
// Main Function: Uses access token struct to validate login
//
// Returns: Sets task's LoginJWT and (will also set the JWT inside of the db too)
func (t *Task) getJWT() error {
	fmt.Println("getting jwt")

	// appUuid := "0c371750-dbf1-42a9-824d-1e59e852d6ee"
	// userUuid := "5adc6df9-ca43-4158-9648-17c147affef8"

	appUuid := uuid.New().String()
	userUuid := "5ebe2b9c-6089-4214-957d-7633b45f4b51"
	ecommsig, err := jwt.MakeEcommJWT(appUuid)
	if err != nil {
		return fmt.Errorf("error making jwt: %w", err)
	}
	// os.Exit(1)
	auth := AuthLogin{
		AccessToken:   t.AuthStruct.AccessToken,
		APIServerURL:  t.AuthStruct.APIServerURL,
		AppID:         "gx9zz84e3x",
		UserID:        t.AuthStruct.UserID,
		UniqueAppID:   appUuid,
		SignedAt:      time.Now().Unix(),
		KeyVersion:    "v1",
		UniqueUserID:  userUuid,
		EcomSignature: ecommsig, // eyJhbGciOiJFUzI1NiIsInR5cGUiOiJ2My9tb2JpbGUvYXV0aCJ9.eyJhcHBfdmVyc2lvbiI6IjIuMC4zNTA1OSIsInVuaXF1ZV91c2VyX2lkIjoiY3JlYW0iLCJ1bmlxdWVfYXBwX2lkIjoiMGMzNzE3NTAtZGJmMS00MmE5LTgyNGQtMWU1OWU4NTJkNmVlIiwic3RvcmVfYXJiaXRyYXRpb24iOnRydWUsInNpZ25lZF9hdCI6MTc0MDg1OTk5NDQ4OSwia2V5X3ZlcnNpb24iOiJ2MSJ9.8hPslzdishLIwjHXlUVq5s9E2TUiPgRCw66JpPt7BuEc8z1kPPiWJnS-QLuPiiGCPenj6UT5jmvsVTsjgmxcRA
		AppVersion:    "2.0.35059",
	}

	authPayload, err := json.Marshal(auth)
	if err != nil {
		return fmt.Errorf("error marshalling auth payload")
	}

	req, err := http.NewRequest(http.MethodPost, "https://us.ecom.samsung.com/v3/sso/mobile/auth", strings.NewReader(string(authPayload)))
	if err != nil {
		return err
	}

	// TODO: Add better headers here hahahahahhaha
	req.Header = map[string][]string{
		"x-ecom-app-id": {"samsung-mobile-app-ios"},
	}

	resp, err := t.Client.Do(req)
	if err != nil {
		return (err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jwt JwtResp
	err = json.Unmarshal(body, &jwt)
	if err != nil {
		return fmt.Errorf("cannot unmarshal jwt repsonse: %w", err)
	}
	t.LoginJWT = jwt.Jwt
	fmt.Println(t.LoginJWT)

	err = t.DB.UpdateAuthToken(t.Email, t.AuthStruct.AccessToken, jwt.Jwt)
	if err != nil {
		fmt.Printf("error updating account for %s: %s\n", t.Email, err)
	}

	return nil
}

func (t *Task) getExchange() error {
	var data = strings.NewReader(`{"devices":[{"sku":"SM-F741UZYEXAA","offer_id":"1738182176334","device_id":"fed8eafe-c6dd-4928-8308-914259a80668","device_info":{"state":{"data_wiped":"Yes","device_working":"Yes","screen_working":"Yes"}}}]}`)
	req, err := http.NewRequest("POST", "https://www.samsung.com/us/api/v1/exchange", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"`)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.samsung.com/us/smartphones/galaxy-z-flip6/buy/galaxy-z-flip6-512gb-unlocked-sm-f741uzsexaa/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var tradein map[string]interface{}
	err = json.Unmarshal(body, &tradein)
	if err != nil {
		return fmt.Errorf("cannot unmarshal trade-in: %w", err)
	}

	if tradein["id"] == nil {
		return fmt.Errorf("trade in id not found")
	}

	t.ExchangeID = tradein["id"].(string)

	return nil
}
func (t *Task) getCart() error {

	t.Log.Info("Getting Cart...")

	e := []Exp{
		{
			SignifydPreAuthActive: true,
			CartSpa:               true,
		},
	}

	c := AppCreateCart{
		CartID:      t.CartID,
		Timezone:    "MST",
		TriggerTags: []string{"B10Q10SHOPAPP"},
		StoreID:     "31_217",
		Experiments: e,
	}

	cPayload, _ := json.Marshal(c)

	var data = strings.NewReader(string(cPayload))
	req, err := http.NewRequest("POST", "https://www.samsung.com/us/api/v4/shopping-carts/?auto_apply=true", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	// req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	// req.Header.Set("sec-ch-ua", `"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`)
	// req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/cart/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "close")
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var cart Response
	err = json.Unmarshal(body, &cart)

	t.CartID = cart.CartID

	return nil
}

// https://www.samsung.com/us/web/app/cart/#/cart-info/?addItem[]=SM-S936UZKEXAA,1&exchange_id=171f6284-6b00-4467-ac22-7f687b11872a,SM-S936UZKEXAA,1&addChildItem[]=3MOS-PCKPREM-TAB-OFFER,SM-S936UZKEXAA,1&addChildItem[]=2-MO-ADOBE-LR,SM-S936UZKEXAA,1&addChildItem[]=PC-OFFER-SIRIUSXM6MOS,SM-S936UZKEXAA,1&step=CART&paymentPlanId=PayinFull
// https://www.samsung.com/us/web/app/cart/#/cart-info/?addItem[]=SM-S938UZKAXAA,1&exchange_id=c5b218f2-61da-475f-98df-10f4750cc152,SM-S938UZKAXAA,1&addChildItem[]=SM-L300NZGAXAA,SM-S938UZKAXAA,1&addChildItem[]=3MOS-PCKPREM-TAB-OFFER,SM-S938UZKAXAA,1&addChildItem[]=YOUTUBE_PREMIUM_3MONTH,SM-S938UZKAXAA,1&addChildItem[]=PC-OFFER-SIRIUSXM6MOS,SM-S938UZKAXAA,1&step=CART&paymentPlanId=PayinFull

func (t *Task) addToCart() error {
	// i := []Item{
	// 	{
	// 		SkuID:    "SM-R630NZAAXAR",
	// 		Quantity: "1",
	// 	},
	// 	// {
	// 	// 	SkuID:    "YOUTUBE_PREMIUM_3MONTH",
	// 	// 	Quantity: "1",
	// 	// },
	// 	// {
	// 	// 	SkuID:    "2-MO-ADOBE-LR",
	// 	// 	Quantity: "1",
	// 	// },
	// 	// {
	// 	// 	SkuID:    "ARCSITE_PREMIUM_30DAY",
	// 	// 	Quantity: "1",
	// 	// },
	// }
	c := []CartItem{
		{
			SkuID:    "SM-R400NZAAXAR",
			Quantity: "1",
			// ExchangeID: t.ExchangeID,
			// LineItems: i,
		},
	}

	e := []Exp{
		{
			SignifydPreAuthActive: true,
			CartSpa:               true,
		},
	}

	cart_struct := CartPayload{
		Experiments: e,
		PostalCode:  "33837",
		StoreId:     "31_1184",
		LineItems:   c,
	}

	payload, err := json.Marshal(cart_struct)
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}

	url := "https://www.samsung.com/us/api/v4/shopping-carts/?auto_apply=true"

	var data = strings.NewReader(string(payload))
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	// req.Header.Set("x-ecom-app-id")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/cart/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "close")
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cart Response
	err = json.Unmarshal(body, &cart)
	if err != nil {
		return fmt.Errorf("error unmarshalling json: %w", err)
	}

	// fmt.Println(cart.UserAlerts.LineItemsAlerts...)
	if len(cart.UserAlerts.LineItemsAlerts) > 0 {
		alert := cart.UserAlerts.LineItemsAlerts[0].(map[string]interface{})["alert_data"].(map[string]interface{})["alert_type"]
		switch alert {
		case ITEM_OOS:
			t.Log.Fatal("Item OOS!")
		case INVALID_EXCHANGE:
			t.Log.Fatal("Invalid Exchange!")
		default:
			fmt.Println(string(body))
			t.Log.Fatalf("unknown alert: %s", cart.UserAlerts.LineItemsAlerts[0].(map[string]interface{})["alert_data"].(map[string]interface{})["alert_type"])
		}
	}

	t.Log.Info("Added To Cart")
	t.CartID = cart.CartID

	return nil

}

func (t *Task) addShipping() error {
	t.Log.Info("Adding Shipping...")

	sPayload, err := json.Marshal(t.ShippingInfo)
	if err != nil {
		return fmt.Errorf("cannot marshal json payload: %w", err)
	}

	var data = strings.NewReader(string(sPayload))

	url := fmt.Sprintf("https://www.samsung.com/us/api/v4.1/shopping-carts/%s/shipping-info?skip_validation=true", t.CartID)

	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	req.Header.Set("x-ecom-locale", "en-US")
	// req.Header.Set("x-client-request-id", "pwa_common_8acc0cc6-fb4b-42f6-9b3a-096d17349635")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/checkout/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	switch resp.StatusCode {
	case 200:
		t.Log.Info("added shipping")
		return nil
	case 401:
		return fmt.Errorf("expried login session")
	default:
		fmt.Println(string(body))
		return fmt.Errorf("unknown error")
	}

}

// also not really needed im p sure
func (t *Task) getShippingRates() error {
	url := fmt.Sprintf("https://www.samsung.com/us/api/v4/shopping-carts/%s/delivery-modes", t.CartID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	// req.Header.Set("x-client-request-id", "pwa_common_e2b9dc72-d64a-4618-9ef3-9d9e25c4cafc")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/checkout/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	// req.Header.Set("Cookie", `country_region=CA-ON; s_fpid=c634dc4b-1bb8-4294-a6f4-87501fa33472; AKA_A2=A; country_codes=us; device_type=pc; at_check=true; s_ecom_scid=d6a55db3-8407-4d76-a0e1-4c4f4764ead2; s_ecom_sc_cnt=1; ecom_vi_USA=ZGZmYjM2OTctOGY1OC00NzYyLWJmMmMtM2Y2YWE5YjhkZmEy; ecom_session_id_USA=OWFiYmE1OGEtZTRkZC00NjljLWE5MmItZDhmNjc2MGI3NWM3; s_ecid=MCMID%7C08233243831789355651510597816113923661; AMCVS_48855C6655783A647F000101%40AdobeOrg=1; __COM_SPEED=H; tracker_device_is_opt_in=true; mboxEdgeCluster=35; ab_test_show_buyLink=true; tracker_device=960b2ea4-6fd9-4bbe-ba44-64dbb46703e2; page_state=Cart - Mixed Product - EIP; kampyle_userid=3dd6-c870-9f86-8204-ca0a-926a-f534-e517; _ga=GA1.1.652919136.1739992635; _fbp=fb.1.1739992635076.154121339781625852; _gcl_au=1.1.470962812.1739992635; __attentive_session_id=b2bd9ad199d641de8c1c24a834061c1b; __attentive_id=4b1a32fc957d4e4990f9cc86299f6d5a; __attentive_cco=1739992635485; tfpsi=78c85cee-49eb-42c5-929b-b6d340f34b91; __attentive_ss_referrer=ORGANIC; __attentive_dv=1; _aeaid=78907008-02b6-479d-a1dd-a1ce86f653e3; aelastsite=IqynIxmTpWLTmucVE6VJcc93GvEXZoAScAF4P8JKEh9cCxYtr%2FMAw5Qj62MwuOMg; aelreadersettings=%7B%22c_big%22%3A0%2C%22rg%22%3A0%2C%22memph%22%3A0%2C%22contrast_setting%22%3A0%2C%22colorshift_setting%22%3A0%2C%22text_size_setting%22%3A0%2C%22space_setting%22%3A0%2C%22font_setting%22%3A0%2C%22k%22%3A0%2C%22k_disable_default%22%3A0%2C%22hlt%22%3A0%2C%22disable_animations%22%3A0%2C%22display_alt_desc%22%3A0%7D; spr-chat-token-60c1edc94fe1cd452ceb20ba_app_950960=; iadvize-6528-vuid=7a1b035c2a88480ca2137dd52ec43536e729d058b32d4; __idcontext=eyJjb29raWVJRCI6IjJ0R3lLMTFPdVBmd09IZ1lndW5iQ1BMUzRsZiIsImRldmljZUlEIjoiMnRHeUsxWE1xTzdQd0xsRjQ3TXUzQkxRNkRCIiwiaXYiOiIiLCJ2IjoiIn0%3D; sa_did=ZtCfnBzYABSCVrbxHjKlSaOUxnqRvclj; sa_773-397-549898={"US":{"status":"GRANTED","updated":"2025-02-19T19:17:20.260Z","clientId":"kv5di1wr19","deviceId":"ZtCfnBzYABSCVrbxHjKlSaOUxnqRvclj"}}; AAMC_samsungelectronicsamericainc_0=REGION%7C9; aam_test=segs%3D7431031; aam_sc=aamsc%3D4718718; aam_uuid=08205780020363145131511229012232379089; AMCV_48855C6655783A647F000101%40AdobeOrg=1585540135%7CMCIDTS%7C20139%7CMCMID%7C08233243831789355651510597816113923661%7CMCAID%7CNONE%7CMCOPTOUT-1739999906s%7CNONE%7CMCAAMLH-1740597506%7C9%7CMCAAMB-1740597506%7Cj8Odv6LonN4r3an7LhD3WZrU1bUpAkFkkiY1ncBR96t2PTI%7CvVersion%7C4.4.0%7CMCCIDH%7C-638464310; TS011dc0a2=011ef6f086f38e7d19caef2e913e3477410b8aca823ce27707ddb0287a67a63a2f7c45316a4d7bc2e2c7e0434914867bf0c0832559; TS0171c81e=011ef6f086f38e7d19caef2e913e3477410b8aca823ce27707ddb0287a67a63a2f7c45316a4d7bc2e2c7e0434914867bf0c0832559; mbox=session#cfc25a7d9f7c4fffa56ca7aa310b2f00#1739994574|PC#cfc25a7d9f7c4fffa56ca7aa310b2f00.35_0#1803237514; OptanonConsent=isGpcEnabled=0&datestamp=Wed+Feb+19+2025+12%3A18%3A35+GMT-0700+(Mountain+Standard+Time)&version=202307.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=03cb5bb5-6d5c-4c91-af34-89b53234e9d6&interactionCount=1&landingPath=https%3A%2F%2Fwww.samsung.com%2Fus%2Fweb%2Fapp%2Fcheckout%2F&groups=C0001%3A1%2CC0003%3A1%2CC0002%3A1%2CBG17%3A1%2CC0004%3A1; _ga_0JTZHYKZ5Z=GS1.1.1739992634.1.1.1739992717.53.0.0; _uetsid=20bcbc70eef611efac40c7744e0d7eb8; _uetvid=20bcf6f0eef611ef955b4191323bb5ef; _attn_=eyJ1Ijoie1wiY29cIjoxNzM5OTkyNjM1NDgxLFwidW9cIjoxNzM5OTkyNjM1NDgxLFwibWFcIjoyMTkwMCxcImluXCI6ZmFsc2UsXCJ2YWxcIjpcIjRiMWEzMmZjOTU3ZDRlNDk5MGY5Y2M4NjI5OWY2ZDVhXCJ9Iiwic2VzIjoie1widmFsXCI6XCJiMmJkOWFkMTk5ZDY0MWRlOGMxYzI0YTgzNDA2MWMxYlwiLFwidW9cIjoxNzM5OTkyNzE3MDY2LFwiY29cIjoxNzM5OTkyNzE3MDY2LFwibWFcIjowLjAyMDgzMzMzMzMzMzMzMzMzMn0ifQ==; da_sid=7C70531E8E3AAE8AD50DAA13A149DC858C.1|3|0|3; da_lid=4F43602D9A7AEA11405CBB99E34B968E3F|0|0|0; da_intState=; __attentive_pv=3; kampyleUserSession=1739992718582; kampyleUserSessionsCount=3; kampyleSessionPageCounter=1; s_pers=%201%3D1%7C1739994615964%3B%20first_page_visit%3Dhttps%253A%252F%252Fwww.samsung.com%252Fus%252Fweb%252Fapp%252Fcheckout%252F%7C1739994615964%3B%20s_nr%3D1739992815965-New%7C1742584815965%3B%20gpv_pn%3Dnew%2520ecom%2520step2%257Ccheckout%7C1739994615965%3B%20s_fbsr%3D1%7C1739994615966%3B; s_sess=%20c_m%3DundefinedTyped%252FBookmarkedTyped%252FBookmarkedundefined%3B%20s_ppvl%3Dnew%252520ecom%252520step2%25257Ccheckout%252C16%252C16%252C684%252C714%252C684%252C1440%252C900%252C2%252CL%3B%20s_cc%3Dtrue%3B%20s_sq%3D%3B%20s_ppv%3Dnew%252520ecom%252520step2%25257Ccheckout%252C27%252C67%252C2998%252C714%252C684%252C1440%252C900%252C2%252CL%3B; utag_main=v_id:01951fa486f5000e1b4cb141432505075004006d00942samsung_live$_sn:1$_se:16$_ss:0$_st:1739994624710$ses_id:1739992631030%3Bexp-session$_pn:3%3Bexp-session$_prevpage:%3Bexp-1739996424714$adobe_mcid:08233243831789355651510597816113923661%3Bexp-session$aa_vid:%3Bexp-session$tapid_reset:true%3Bexp-1771528634243$dc_visit:1$dc_event:4%3Bexp-session$dc_region:us-west-2%3Bexp-session`)
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	switch resp.StatusCode {
	case 200:
		var delivery DeliveryGroups
		err = json.Unmarshal(body, &delivery)
		if err != nil {
			return fmt.Errorf("error unmarshalling delivery group json: %w", err)
		}

		keys := make([]string, 0, len(delivery.CartPayload.DeliveryGroups))
		for k := range delivery.CartPayload.DeliveryGroups {
			keys = append(keys, k)
		}

		if len(keys) == 0 {
			return fmt.Errorf("no delivery groups found in shipping rates")
		}

		t.DeliveryGroup = keys[0]
	default:
		fmt.Println(string(body))
		return fmt.Errorf("unknown error")
	}

	return nil
}

func (t *Task) setShippingRate() error {
	type M map[string]string
	var shippingRates []M
	shippingGroup := M{"delivery_group_id": t.DeliveryGroup, "delivery_sku": "DM_IM_FS_EXPEDITED"}
	shippingRates = append(shippingRates, shippingGroup)
	shippingRatePayload, err := json.Marshal(shippingRates)
	if err != nil {
		return fmt.Errorf("cannot marshal shipping rate payload: %w", err)
	}

	var data = strings.NewReader(string(shippingRatePayload))
	url := fmt.Sprintf("https://www.samsung.com/us/api/v4/shopping-carts/%s/delivery-modes", t.CartID)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("x-client-request-id", "pwa_common_e289377d-c3bb-4003-ad47-ca267ec4aa3c")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/checkout/?appId=samsung-mobile-app-ios")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", `TS0c171c81e=011fc4fd04b7f22f393631b665a9bc84cbdec77b486c325736af79ff57cbe8347a276ff1d594fa354a753bc8a21c33cae97a9827d1; device_type=pc; eddzipcode=33837; s_ecom_sc_cnt=1; TS011dc0a2=011fc4fd04b7f22f393631b665a9bc84cbdec77b486c325736af79ff57cbe8347a276ff1d594fa354a753bc8a21c33cae97a9827d1; s_ecom_scid=616a9c47-95b3-4e34-9b5e-a7159269dd73; s_sq=sssamsungnewusdev%3D%2526c.%2526a.%2526activitymap.%2526page%253Dhttps%25253A%25252F%25252Fwww.samsung.com%25252Fus%25252Fweb%25252Fapp%25252Fcart%25252F%25253FappId%25253Dsamsung-mobile-app-ios%2526link%253DCheckout%2526region%253Dpdp-page%2526.activitymap%2526.a%2526.c; spr-chat-token-60c1edc94fe1cd452ceb20ba_app_950960=; da_intState=0; da_lid=F8AB07409A7DEA11E5CFBB99E51BAA26A1|0|0|0; da_sid=35D086B68E3AAE809FEEAA13A134FAFA6F.0|4|0|3; mbox=PC#05146f25e58a42ad92beb6d87a6711c2.35_0#1804363120|session#8940dd087eac454d9fb56eced1bcf011#1741120184; OptanonConsent=isGpcEnabled=0&datestamp=Tue+Mar+04+2025+12%3A58%3A42+GMT-0700+(Mountain+Standard+Time)&version=202307.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=33e9112d-5f2a-48e5-bc8e-7463891be171&interactionCount=0&landingPath=NotLandingPage&groups=C0001%3A1%2CC0003%3A1%2CC0002%3A1%2CBG17%3A1%2CC0004%3A1&AwaitingReconsent=false; utag_main=v_id:01934f982f980015da3337750b7905058004505000718samsung_live$_sn:19$_se:1$_ss:1$_st:1741120120308$tapid_reset:true%3Bexp-1763743197132$dc_visit:6$_prevpage:undefined%3Bexp-1741119255937$ses_id:1741118320308%3Bexp-session$_pn:1%3Bexp-session; at_check=true; epp_store_segment=Education; mboxEdgeCluster=35; checkout_continuity_service=faaccc75-411e-410f-9798-2dfb687b6000; tracker_device=faaccc75-411e-410f-9798-2dfb687b6000; FLUTTER_CART_ID=616a9c47-95b3-4e34-9b5e-a7159269dd73; jwt_USA=ZXlKaGJHY2lPaUpTVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnpZMmhsYldGZmRtVnljMmx2YmlJNklqSXVNQ0lzSW1WNGNDSTZNVGMwTXpVek5USXdNaTQ1Tmpjc0ltMWhlRjlsZUhCcGNua2lPakUzTkRFeU1ESXpPVEV1T1RZM0xDSjJaWEp6YVc5dUlqb2lNUzR3SWl3aVkybHdhR1Z5SWpvaVZUSkdjMlJIVm10WU1UbDRUVTB3VTFCRFlsUmFiRFJ2YjJkeVpHbHlSVEJXVEdkVmNIRk5hVTB6UVhKWll5czFVSGNyTW5Oak5GTlZNRXBXZDB0MlpEVldhMjFGUnpsbVZscDRNa0p0UmtkdlJWaGxVamxzVFcwd1REZHRaV0V4Y0RRd1pHbDJOREp4UzBWcVF6WktjbkZwZFdFcmFuTnRWMnAxUkVrM05VVnFjVU5UZUhFeWVTdEJVMWhRVjFaWFR6SmxhR2R5VldsV1JGUjBRWEp4ZFhoTGNESkVVa05ETnpSaFRpdEJLMHRWVUVnMVIzZFNVSG94VVVsTVlVOWpaV1ZsUjBOU2IwbG1iM2g2YzBNeFpWUTFkbWxRY1hneVpHUlJTRmRGWjFJelFVZE9OV05YV0hkdVUyUlFabUk1Ym1aRFRWcERhM1JHUkdaTVkzZEpaSFEyUVhOUVV6SjJlRkUxVUV4RksxRjJaM1JZYms4dmIyZHNORkU0ZGxwVGIzbFVlR3hEWnpsRGJGTXJXWFpxTldSdWJXOW5lVkZJWkVkcWVVVlZla2RySzFGemVWbHdRVGhoYlVkalZIaHhjVk01S3poQ1FUbEpTblJWVVRKMVExVnlPRTlJTkhCS2JuZGlkWGhxU0V4Mk56ZzJiU3RYZFVGUlpYbFljSE4zTVhoS2IyNUtlR3h4YkZaWGRrUmthVTR2WW1zMmFrTlpkekIzWjIxYVNUTk1SekJIV205b2RVOTVWVFoyYTJaSVlqZzRhMUozZDI1cWNXWnRhM0pYVm01MVZuRjJSbEJxUm1wSFVURkJhRE5NY0VOdmMyaDNkWFZTWmpSNmEyaGxXbXB1TldwVFRWWkxkMDFXTWsxbE0wVnpia0pCTlRkemNFdEtla2RaUW5ZMVRWUllOMjgwYm1SdVVVdzBLMXBITlN0WmJ6TjBSSEkzYjJWamVVcFBTVEZaVjBVMFdrNVVWVTVzUmpWbk0wdDFOSFo1VGxkV2NYVnpNek5aT1ZScVIyUktSMFpaU21ka1dVSm9URGRMYWxoeVJIaEZPRXBsUW5BclJYZHdTRkpIU0VkalMweHphVTlzT0ZSek9WRkJaRFZKWWs1bFVGWlpUblI0VWpoNFFVdFlSV3gwYUd0elRVRlNVR3RuSzNsVmNTOVRjbkpNVTNkdVkyTlJkRmd4WTFBd2NtaFVLemN6V2toaGJFSkhURkZDY1hGQlVuUmxRbGRNVlRGYVJUVnhMM2RITVZwUE1DOTRVRWx0UzJKSWNGTXJURzQ1WTNoWlZ6RnVOMmhhTDBndk1GRjBMemRZYVhoVWRYaE9OMFYxYzFGblVuTmtZbUZqVDJaMVQwRlFaVkpRVjFoWmVWWmljVGhrWVhwRmNtWk5WMk5tVFN0dk5WWlVaMk5MSzNBMU9FTjVjVGxRUTNWRWVHRXhVSEpTVW1sQlRWWTBZakpWVFd3d1RFdDBPWGczZUhwM2JYQkdPVVZXVGtGVFVYazRUa1ZaZDFwdVpDdEdZVXN3UmtKR09YWjFlblpFYm1KM1RGZzVWWEpoYm10aE9HbGlVWFV6VXpsa1JXeG5VMHBEVm5BM1NWWTRiVFZXVTNaYVpWSktNWGR0VVZVeldVdDZTblp5YUdZME0zVk9VMEpPUjJoamJXSlZSWEZIY0c5T0x6Vk1lVEl2VW10U1JHcEZjSG94VkZWaVExaDJjVTFwTmpCVE1VNUJabHBISzNOTkswUnZUbEZaZG5WU1RXZEdVMHN6ZVVka00zbFBhMDVyT0VkMVJUVnlWRVpTWVhFMlJ6TjZZVGRxYlRSeVNIQnRkemhyUW1SMGVVbFVNbkV6TTBkMUwzQklWMjVyWlROWlduWlZWbmhCTWxabWJTdHBha2gwY2xjM01sUXpaakZpU2tKMGMwaFBabFZKTjJkUlVWSjVMemR0VnpKbmVsVlFNV0pvTVcxWWNtOHlRa1JEYVN0VFEycHhabFZ1ZFdSR1VIWlRNR3RLZDJaR1RFeFVZV3BWZEhOaFJXSkhNM1pGU1hSTk1WWTVablJWZGtwbVkxaDJja0p1WTFkVVdIUjNWR28yTldsYVRtcGlUVEo1YTFOSUt6TlFNV3BUUlQwaUxDSnBZWFFpT2pFM05ERXhNVFl3TURJc0ltTnZiWEJ5WlhOemFXOXVJam9pZW14cFlpSjkuV2p3U1ZCcnR5QU5IYklwN29kMWZUZy1fX2oySEs5OEJ1NjdZaXNlVVdtR3pOX2Q3T1poWElpSXZYcGJkeWFyYklCWF9aSnRJTmJ4Uk1NTlNtWGRoYkVuRm94UE5vOTRtVFdrMEpmdng2a0lIbFB5d1R0TDU3amFpQ3BRRGlUOVFqcVNLa2ZkbTczckIzbFNYUkZUeVRfcmRKbUUxWGxGZ2Y3VEtrWEtRTVhmZ3BFV3k5Y0NBdGpOUGZuLXo4U2lDanoyM0FOVjlxX0FsYkNSR1RqcXZYbk43UjFJbGhjYXRVOVRuLXRoWWZnRF84dVI1NllhNUlURGhZam11MGp0Wkp2YzFjamlpaTdzVWpGV0hBRU1vUkQ3RmhkNnloTTdQR1hqb2NGakZRbmE4M0ZCVVZJTHJZQWpDWXI0MGFyTzFPRENXSkVOOGxUNGhfenJKRFBrX01n; ecom_session_id_USA=MjYyYWNjODQtOGYwNS00YTk2LThiZDQtNTI3OTkzYzdmM2Iz; rewards_tier=Silver; FLUTTER_USER_AGENT=ShopSamsung-iOS/3.0.39/iPad (com.samsung.ShopSamsung); build:2025010602; iPadOS 17.2); s_sess=%20s_ppvl%3D%3B%20s_cc%3Dtrue%3B%20s_ppv%3Dnew%252520ecom%252520step3%25257Cthankyou%252C85%252C100%252C1479%252C768%252C948%252C768%252C1024%252C2%252CP%3B; __idcontext=eyJjb29raWVJRCI6IjJwQVM0c2NrZ0gzNmlLSHZ3dWlKOHRTaUNqaCIsImRldmljZUlEIjoiMnBBUzRyblRqQThud3hqTHdSWmZEQW45ckxDIiwiaXYiOiIiLCJ2IjoiIn0%3D; _ga_VZXQ7W626J=GS1.2.1741115659.3.0.1741115659.60.0.0; aelastsite=IqynIxmTpWLTmucVE6VJcc93GvEXZoAScAF4P8JKEh9cCxYtr%2FMAw5Qj62MwuOMg; aelreadersettings=%7B%22c_big%22%3A0%2C%22rg%22%3A0%2C%22memph%22%3A0%2C%22contrast_setting%22%3A0%2C%22colorshift_setting%22%3A0%2C%22text_size_setting%22%3A0%2C%22space_setting%22%3A0%2C%22font_setting%22%3A0%2C%22k%22%3A0%2C%22k_disable_default%22%3A0%2C%22hlt%22%3A0%2C%22disable_animations%22%3A0%2C%22display_alt_desc%22%3A0%7D; _aeaid=6a1d68e0-4a34-4503-8bcf-e0230d0d7a52; _fbp=fb.1.1740679117668.644835267752335349; _gcl_au=1.1.1046515211.1740679117; __attentive_id=8561cb60f54347d3a28aa209e6a426ce; AAMC_samsungelectronicsamericainc_0=REGION%7C9; _ga_0JTZHYKZ5Z=GS1.1.1741115656.2.0.1741115656.60.0.0; _uetsid=0515fb10f87b11efab0735ce22920672; _uetvid=3e44b1d0a82711ef9c6e7d3f161ea51a; aam_sc=aamsc%3D4718718; aam_test=segs%3D7431031; aam_uuid=11491313183051517203400930256289297356; _attn_=eyJ1Ijoie1wiY29cIjoxNzMyMjA3MTk3NDEyLFwidW9cIjoxNzMyMjA3MTk3NDEyLFwibWFcIjoyMTkwMCxcImluXCI6ZmFsc2UsXCJ2YWxcIjpcIjg1NjFjYjYwZjU0MzQ3ZDNhMjhhYTIwOWU2YTQyNmNlXCJ9Iiwic2VzIjoie1widmFsXCI6XCI2NzQyNDRkNzU0YWE0MTQ3ODhmMjQxNTRiMDlhZTkwMlwiLFwidW9cIjoxNzQxMTE1NjU2NDk2LFwiY29cIjoxNzQxMTE1NjU2NDk2LFwibWFcIjowLjAyMDgzMzMzMzMzMzMzMzMzMn0ifQ==; page_state=state1.4|Cart - Minimum Not Met Cart - Mixed Product - EIP; s_pers=%201%3D1%7C1741117455946%3B%20first_page_visit%3Dhttps%253A%252F%252Fwww.samsung.com%252Fus%252Fweb%252Fexpress%252Forder-confirm%252Fa340ae07-0c76-42fd-9d6d-87278e87948d%7C1741117455947%3B%20s_nr%3D1741115655947-Repeat%7C1743707655947%3B%20gpv_pn%3Dnew%2520ecom%2520step3%257Cthankyou%7C1741117455948%3B%20s_fbsr%3D1%7C1741117455949%3B; DECLINED_DATE=1741115654636; RT="z=1&dm=samsung.com&si=6a97d89e-e314-49e8-bf36-c338e122f667&ss=m7uvaymc&sl=1&tt=1cj&bcn=%2F%2F17de4c16.akstat.io%2F&ld=1ct"; kampyleInvitePresented=true; kampyleSessionPageCounter=2; kampyleUserPercentile=77.35429754556344; AB_test_HideHeaderLinks=false; DATradein=false; ab_test_show_buyLink=true; signifyd_pre_auth_active=true; tgt_bpno=0329622399; AMCVS_48855C6655783A647F000101%40AdobeOrg=1; AMCV_48855C6655783A647F000101%40AdobeOrg=1585540135%7CMCIDTS%7C20151%7CMCMID%7C11520140382857642123403813666805772112%7CMCAID%7CNONE%7CMCOPTOUT-1741121696s%7CNONE%7CMCAAMLH-1741719296%7C9%7CMCAAMB-1741719296%7Cj8Odv6LonN4r3an7LhD3WZrU1bUpAkFkkiY1ncBR96t2PTI%7CvVersion%7C4.4.0%7CMCCIDH%7C-166383409; __COM_SPEED=H; country_codes=us; country_region=CA-ON; myOrderPZN=true; pdm=D; storeid=31_482; tmktid=4789741110; tmktname=Samsung Education Offer Program; iadvize-6528-vuid=b4b195c8fe9b4c77a43a8eefd8d8f246f41a0dee74b64; s_ecid=MCMID%7C11520140382857642123403813666805772112; kampyleUserSession=1740775040135; kampyleUserSessionsCount=6; apevwkcart=Wunderkind; pznapewkcartckie_575496=Wunderkind; sssuscust=true; apay-session-set=Gp5tQ2lv6viG8e%2BlMt2%2B6l1zbblM45t%2BLeJz75jSj5yMtIb%2BwIfDYWa%2FoVTr2%2FY%3D; ecom_vi_USA=N2U5NDhhMjItOTZlNS00MDg0LTliMjQtYjE2NTM2Njk0ZmIw; BVBRANDID=b2e9fb85-69c8-40d1-8e6c-095b14a2aecb; sa_773-397-549898={"US":{"status":"GRANTED","updated":"2024-11-21T16:40:03.194Z","clientId":"kv5di1wr19","deviceId":"ntTzkXhnHlwSAyTxlkBOEOnZOlNccLtG"}}; sa_did=ntTzkXhnHlwSAyTxlkBOEOnZOlNccLtG; __attentive_cco=1732207197413; __attentive_dv=1; __attentive_ss_referrer=https://www.samsung.com/us/web/express/order-confirm/d4cd25c6-0310-4f80-9899-f769e25d069b?appId=samsung-mobile-app-ios; kampyle_userid=519e-78fc-d246-c172-75ba-e735-3529-8755; tracker_device_is_opt_in=true; s_fpid=c2e7140e-a445-433b-9043-aae5ff7299b4`)
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return fmt.Errorf("cannot unmarshall json into interface: %w", err)
	}

	switch resp.StatusCode {
	case 200:
		t.Log.Info("Applied Shipping Rate!")
		break
	default:
		// fmt.Println(string(body))
		t.Log.Errorf("Error: %s", res["message"].(string))
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	return nil
}

func (t *Task) applyCoupon() error {

	// var data = strings.NewReader(`{"coupon_code":"woahaaaaa"}`)

	c := make(map[string]string)

	c["coupon_code"] = t.CouponCode
	cPayload, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshalling code")
	}

	url := fmt.Sprintf("https://www.samsung.com/us/api/v4/shopping-carts/%s/coupon-codes", t.CartID)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(cPayload)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("x-client-request-id", "pwa_common_8ea50cb4-56e3-423f-97b2-bf616e9c2b41")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/cart/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cart Response
	err = json.Unmarshal(body, &cart)

	switch resp.StatusCode {
	case 200:
		t.Log.Info("Coupon Applied Successfully")
		t.Log.Infof("Cart Total: %.02f", cart.Cost.Total)
		return nil
	case 409:
		return fmt.Errorf("invalid coupon")
	default:
		fmt.Println(string(body))
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

}

func (t *Task) addBilling() error {

	t.Log.Info("Adding Billing...")

	encResult := enc.EncPayment(t.CardInfo.Number, t.CardInfo.ExpMonth, t.CardInfo.ExpYear, t.CardInfo.CVV)
	url := fmt.Sprintf("https://www.samsung.com/us/api/v4.1/shopping-carts/%s/payment/apply-payment", t.CartID)

	// gRecaptchaResponse

	t.Log.Info("Solving v3 captcha...")
	cap, err := captcha.SolveV3()
	if err != nil {
		return fmt.Errorf("error solving captcha: %w", err)
	}

	t.ReCaptchaToken = fmt.Sprintf("%v", cap["gRecaptchaResponse"])

	// if err := t.solveReCaptcha("https://www.samsung.com/us/web/app/checkout/"); err != nil {
	// 	fmt.Printf("error getting recaptcha: %s\n", err)
	// }

	e := EncPayment{
		EncryptedPayload:  encResult.EncryptedPaymentContext.EncryptedPayload,
		EncryptedPassword: encResult.EncryptedPaymentContext.EncryptedPassword,
	}

	b := BillingInfo{}

	billing := BillingPayload{
		PaymentMethod:           encResult.PaymentMethod,
		EncryptedPaymentContext: e,
		BillingInfo:             b,
		Token:                   t.ReCaptchaToken,
		CaptchaProvider:         "recaptcha_enterprise",
	}

	t.EncPayload = billing

	billingPayload, err := json.Marshal(billing)
	if err != nil {
		return fmt.Errorf("error marshalling billing info: %w", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(billingPayload)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("x-client-request-id", "pwa_common_a53917ac-52a6-4363-992d-89e193df97ba")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/checkout/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	switch resp.StatusCode {
	case 200:
		t.Log.Info("Billing Information Added")
		return nil
	default:
		fmt.Println(string(body))
		return fmt.Errorf("invalid status code: %w", err)
	}

}

func (t *Task) submitOrder() error {

	t.Log.Info("Submitting Order...")

	url := fmt.Sprintf("https://www.samsung.com/us/api/v4.1/shopping-carts/%s/payment/one-step-transaction", t.CartID)

	e := EncPayment{
		EncryptedPayload:  t.EncPayload.EncryptedPaymentContext.EncryptedPayload,
		EncryptedPassword: t.EncPayload.EncryptedPaymentContext.EncryptedPassword,
	}

	b := OrderPayload{
		PaymentMethod:           "adyen_cards",
		EncryptedPaymentContext: e,
	}

	bData, err := json.Marshal(b)
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}

	var data = strings.NewReader(string(bData))
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPhone (com.samsung.ShopSamsung)-Flutter-Flutter")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("x-client-request-id", "pwa_common_2b41a30d-9026-486c-8b61-6648b82dd69f")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/checkout/?appId=samsung-mobile-app-ios")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var order map[string]interface{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return fmt.Errorf("cannot unmarshal json to interface: %w", err)
	}

	switch resp.StatusCode {
	case 200:
		t.Log.Info("Order Placed Successfully!")
		break

	default:
		t.Log.Errorf("invalid status code: %d", resp.StatusCode)
		t.Log.Errorf("Error: %s", order["message"].(string))

	}

	return nil
}

func (t *Task) checkOrder() error {

	t.Log.Info("Checking Order...")

	return nil
}
