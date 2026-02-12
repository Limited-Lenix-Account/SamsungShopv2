package task

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

func (t *Task) appLogin() error {

	req, err := http.NewRequest("GET", "https://account.samsung.com/accounts/eCommerce_US/signInGate?clientId=gx9zz84e3x&countryCode=US&tokenType=TOKEN&deviceType=APP&deviceUniqueID=&devicePhysicalAddressText=&deviceOSVersion=17.2&competitorDeviceYNFlag=N&locale=en&redirect_uri=https://ecommapi.ecom-mobile-samsung.com/sso/callback&state=7c658120-af78-4d12-8521-b2f7ec3aa57b&", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "account.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko)")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Dest", "document")
	// req.Header.Set("Cookie", `device_type=pc; spr-chat-token-60c1edc94fe1cd452ceb20ba_app_950960=; da_intState=0; da_lid=F8AB07409A7DEA11E5CFBB99E51BAA26A1|0|0|0; da_sid=FD6AA4808E33AE8E68A3AA13A11F4895C7.0|4|0|3; mbox=PC#05146f25e58a42ad92beb6d87a6711c2.35_0#1803934468|session#6e083aae0a1848ee9f8b9f9140c0ffa9#1740691533; OptanonConsent=isGpcEnabled=0&datestamp=Thu+Feb+27+2025+13%3A54%3A29+GMT-0700+(Mountain+Standard+Time)&version=202307.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=33e9112d-5f2a-48e5-bc8e-7463891be171&interactionCount=0&landingPath=NotLandingPage&groups=C0001%3A1%2CC0003%3A1%2CC0002%3A1%2CBG17%3A1%2CC0004%3A1&AwaitingReconsent=false; s_ecom_sc_cnt=0; AMCVS_48855C6655783A647F000101%40AdobeOrg=1; AMCV_48855C6655783A647F000101%40AdobeOrg=1585540135%7CMCIDTS%7C20146%7CMCMID%7C11520140382857642123403813666805772112%7CMCAID%7CNONE%7CMCOPTOUT-1740696867s%7CNONE%7CMCAAMLH-1741294467%7C9%7CMCAAMB-1741294467%7Cj8Odv6LonN4r3an7LhD3WZrU1bUpAkFkkiY1ncBR96t2PTI%7CvVersion%7C4.4.0%7CMCCIDH%7C-166383409; TS0171c81e=015e3f34b926e7ee05531f53035dc0decfe077b301b263be926f55bfc3d0adfd14271183d81af3bec41521ed35da3b232e8309e370; at_check=true; ecom_session_id_USA=ZTk3MmM4M2MtM2JhMi00NWU4LWI2MjktYzU0ODViZGNkYzRk; eddzipcode=10001; mboxEdgeCluster=35; s_ecom_scid=b8147ecb-e7e1-4c63-8ebb-6dc62db88a5d; utag_main=v_id:01934f982f980015da3337750b7905058004505000718samsung_live$_sn:5$_se:27$_ss:0$_st:1740691467188$tapid_reset:true%3Bexp-1763743197132$dc_visit:4$ses_id:1740686726208%3Bexp-session$_pn:22%3Bexp-session$_prevpage:undefined%3Bexp-1740692625373$adobe_mcid:11520140382857642123403813666805772112%3Bexp-session$aa_vid:%3Bexp-session$dc_event:5%3Bexp-session$dc_region:us-west-2%3Bexp-session; FLUTTER_CART_ID=b8147ecb-e7e1-4c63-8ebb-6dc62db88a5d; country_codes=us; country_region=CA-ON; jwt_USA=ZXlKaGJHY2lPaUpTVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnpZMmhsYldGZmRtVnljMmx2YmlJNklqSXVNQ0lzSW1WNGNDSTZNVGMwTXpFd09EZzBNQzQwTnpNc0ltMWhlRjlsZUhCcGNua2lPakUzTkRBM056WXdOREF1TkRjekxDSjJaWEp6YVc5dUlqb2lNUzR3SWl3aVkybHdhR1Z5SWpvaVZUSkdjMlJIVm10WU1UbEhSRUZqYWpSa1FqaFJiR2g2ZDBOUk1VcFlhbU1yUTI1V1dEUTNNR1pOZW1NMU1reElkalpSVDJ0RGQyUXlObVZCVW1sVVRXbDVja2RNU0ZSM2FsZFBiVTA0UmpGUFVsRkpObTF0Y2tKc2JHRjBSM2RhTDBsVU1WZE9jekkxVUhCMFpqWnNOak5CVjA0NVdWTm1USE0yZUVOVFpUWm1kMkpJWVRaTlkySjNaV3M0ZW1kTFJEY3laemhzV25OWldrUlNaVlZoY1VVeFdVcHFlVEpITjB4M1pVY3haR1ZOTkRKNVFWYzFkMUpTV1VkUE1uWm1UV1JsVTNweGRXaGtaVU4zYUVRclN6Vm5NRzk2TUc1eVJVVnBhRWhJTUhKT1FqbElZbGMxVUUxVloxcHhTV3MzWlV4V04zQnhabEI1VVZkeFdtaHBRM0ZoYVhsdmFVZ3dUbmhVTTFSTmEydzFVQ3RyVTJ0UFZISnhORXMyZW5kaGRsSnRZV1ZUTkVKbGIxUm1jV3g2VW5ad2RUUmlUVlZYTUhFeWMxQnhUMlIwY0cweWFWVnRZMWxXT0ZSb1lUVTFURTVETHpsdVQxZzRLMUJvTWtwc2RVRXJaVE0zTjJ4eWJYUndSVmhWWm14RFprcHVXRXN2UlVSRGJERk1NRmx3Y1VKS1YzZEJTMnRhTmxWSmJVWnFNR1VyT1doeWRrbGhjbXBWTUUwNVZFRkxWMHd2TjJrNE1VNU9jVGN5ZEU5R09YQjRXR0phUWpSM2RuUlpSVFZVWldGelVrTnBTbE5aVG5saGFVRlJTbGhyUjJwcGQweHZOVTB4YlZKalVYSlNUREUzZEc4ME1FUk9kRGhxWmtwTFJ6Z3JSbGREWWpOdmJYZGxPRmh4U1dodmQyZGpUalY0VVVkVkwzVTFiRTlPUW1KRk1EZEhORVZ0TTNabk1TOTVielF3YTBaSE0ySmhSa1pYTkVGbFFuZzRjWFpWVGxCYVowRjRlbTgyS3pOdFR6bEdibHBJYW10RlkxQjZjR1l5UlRkclQyMW1TVzlFV204NWRrWkxWaTlyWVdnNWFIaDBiMkpFUTJOVWJqWTFSVVF3Y1RCRU0zSTRSRWRMU2psdWNrNXNlRGhWU0RoMVFqSkdRekpFWlVoR1dVSTVOM0V5VUVwb05FOVZNWG95YldoTVJVdHJSV0pzVjBrOUlpd2lhV0YwSWpveE56UXdOamc1TmpRd0xDSmpiMjF3Y21WemMybHZiaUk2SW5wc2FXSWlmUS5sb3VLbkowcm5FTmx6MGU0ZGJXc1VIYTlTSnpUazBINEIxU0hJNFI5RUVKYXd1azRvTUYwWC1kNDd2VkVhSTBiTmRiSDhwUHhqUmEtZ0R3cVBKU25mNkFseDZXZHBTWXVoQmVQaHBPdjF3aVYxX0lTZmJmSkxfeGM0VGNFRm50c2JaNm9hN0pHOG9rTTZ4eWVOS2cwbWpldElGNmljc09sWXR1NTdIVGhEVmxkYkE0SzhzWDVxUW4ya0dhZGZTM2I0MjM3XzQ1ZDlCN1BZWlpuLTVWMkx6ZmlKajBLZERHUUNkR0EtNjE4Vi15NUFyRWhZcEt2T0RmR1dNdTRNenE0czF3S3F2ZjBCT2tSeUhraG9UeHdXYm5CQTZzd3paaUlYT0RGNFN4QjE1S2hZdExOdHFKRDk2UzBLUHhyRXQ0bF90eEFMR1pUampJVXlHZUJzczdSSkE=; FLUTTER_USER_AGENT=ShopSamsung-iOS/3.0.39/iPad (com.samsung.ShopSamsung); build:2025010602; iPadOS 17.2); card_carousel=A; _common_physicalAddressText=dwuimvkbjtwqaccgcdch; _fbp=fb.1.1740679117668.644835267752335349; _ga_0JTZHYKZ5Z=GS1.1.1740688531.3.1.1740689026.20.0.0; tfpsi=17858f6b-268a-4166-84ff-7009fc1c344f; AAMC_samsungelectronicsamericainc_0=REGION%7C9; __idcontext=eyJjb29raWVJRCI6IjJwQVM0c2NrZ0gzNmlLSHZ3dWlKOHRTaUNqaCIsImRldmljZUlEIjoiMnBBUzRyblRqQThud3hqTHdSWmZEQW45ckxDIiwiaXYiOiIiLCJ2IjoiIn0%3D; _gcl_au=1.1.1046515211.1740679117; _uetsid=78098e00f53411ef9ca5cdda772f4150; _uetvid=3e44b1d0a82711ef9c6e7d3f161ea51a; aam_sc=aamsc%3D4718718; aam_test=segs%3D7431031; aam_uuid=11491313183051517203400930256289297356; iadvize-6528-vuid=b4b195c8fe9b4c77a43a8eefd8d8f246f41a0dee74b64; s_ecid=MCMID%7C11520140382857642123403813666805772112; s_pers=%201%3D1%7C1740690825378%3B%20first_page_visit%3Dhttps%253A%252F%252Fwww.samsung.com%252Fus%252Fweb%252Fexpress%252Forder-confirm%252Ffc22e958-293f-48c3-a3a2-bff338988988%7C1740690825380%3B%20s_nr%3D1740689025380-Repeat%7C1743281025380%3B%20gpv_pn%3DCustomer%2520Order%2520Details%2520Page%7C1740690825381%3B%20s_fbsr%3D1%7C1740690825383%3B; RT="z=1&dm=samsung.com&si=6a97d89e-e314-49e8-bf36-c338e122f667&ss=m7nt075r&sl=3&tt=29d&bcn=%2F%2F17de4c0e.akstat.io%2F&ld=9vjw"; _ga_VZXQ7W626J=GS1.2.1740688533.2.0.1740688533.60.0.0; AKA_A2=A; ecom_vi_USA=N2U5NDhhMjItOTZlNS00MDg0LTliMjQtYjE2NTM2Njk0ZmIw; pdm=B; storeid=31_217; BVBRANDID=b2e9fb85-69c8-40d1-8e6c-095b14a2aecb; __COM_SPEED=H; sa_773-397-549898={"US":{"status":"GRANTED","updated":"2024-11-21T16:40:03.194Z","clientId":"kv5di1wr19","deviceId":"ntTzkXhnHlwSAyTxlkBOEOnZOlNccLtG"}}; sa_did=ntTzkXhnHlwSAyTxlkBOEOnZOlNccLtG; s_fpid=c2e7140e-a445-433b-9043-aae5ff7299b4`)
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", body)

	fmt.Println(resp.Cookies())

	return nil
}

func (t *Task) appCreateCart() error {

	// cardId := "c61f2583-d573-48cc-a9cf-d5ba8bf8c65e"

	e := []Exp{
		{
			SignifydPreAuthActive: true,
			CartSpa:               true,
		},
	}

	cart := AppCreateCart{
		// CartID:      cardId,
		Timezone:    "MST",
		TriggerTags: []string{"B5Q5SHOPAPP"},
		StoreID:     "31_217",
		Experiments: e,
	}

	cartPayload, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	var data = strings.NewReader(string(cartPayload))
	req, err := http.NewRequest("POST", "https://www.samsung.com/us/api/v4/shopping-carts/?auto_apply=true", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("x-ecom-app-id", "ShopSamsung-iOS/3.0.39/iPad (com.samsung.ShopSamsung)-Flutter-Flutter")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko)")
	req.Header.Set("x-ecom-jwt", t.LoginJWT)
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/cart/?appId=samsung-mobile-app-ios")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	return nil
}
