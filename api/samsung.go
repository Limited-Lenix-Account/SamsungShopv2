package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type login struct {
	Email            string `json:"iptLgnID"`
	Password         string `json:"iptLgnPD"`
	LoginKey         string `json:"svcIptLgnKY"`
	LoginIV          string `json:"svcIptLgnIV"`
	RememberPassword bool   `json:"remIdChkYN"`
	Assertion        any    `json:"assertion"`
}

func Login(cred map[string]string) {

	l := login{
		Email:            cred["svcIptLgnID"],
		Password:         cred["svcIptLgnPD"],
		LoginKey:         cred["svcIptLgnKY"],
		LoginIV:          cred["svcIptLgnIV"],
		RememberPassword: false,
	}

	lBody, err := json.Marshal(l)
	if err != nil {
		log.Fatal("error marshalling json")
	}

	client := &http.Client{}
	var data = strings.NewReader(string(lBody))
	req, err := http.NewRequest("POST", "https://account.samsung.com/accounts/v1/samsung_com_us/signInProc?v=1739208533942", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "account.samsung.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("X-CSRF-TOKEN", "da2f9152-0971-4cc9-8f8e-f2d0aafa6f54")
	req.Header.Set("x-recaptcha-token", "03AFcWeA6VXdqz6w9OURmNtzaxUale1noA7g5MrdVbkEuHV2BcL_4lkHVlTVbe93PdWxWzwZFcXm5F0rLZJrFm-kx43kqJ8ja_Piwhdj6GZodkYn9Hc_tgj8C2FKzPtSBReUgdwUtiWjYWPTGRuG9mf4tPHDEggPiM3uNWtnBQlIvMScgymmKcTOVkiAObJh3JicOILQZiXaKQby9rf-QyMcfOLqZakkV7Uj1uqOAkL2-rQwq82XGUiXqFW9gLXR--fsAGvrcwOSZGsM8FvEuTDxV12Yl5YN6vi-BNyseXS4bceCx4Ef87oEw4g1R_gU0xNfC1C-LzbTVxkVE7RDkaIQ51eTp4G-eKkhdl8DC_1Efrl2gCDlBlmzgbTpKjx7xb5AMSNFWieJo1VeI9nm0_lEU-ygl8NpQyY3cGJwVF-6wYPv0idNcLGoYeB7FSiuTSrXnrV4ey7AbiU4nN03VrH7utyyAzA1G5wWCxxDjDwBLJTPynPhuOsDXdwinee4bvfBUYruFVpoD8Wda8MqnXNagG2qIwfKQ0xWSNxH7G-Lqdi8b6LpFyrznICFOAdmuBz6DV96q3dQxPq3tZpa5iBFrec8sSq4sAzqm5aHeDY8NRtl4IfBbxDRewKQoBJshxumHjpGbxaul_usdELsCnBoOHHfd345PMNRkEuS-vjNQRcavY2hU6_NEKzJUdINdDBTu5CtuNcwQ3xVWjwF4V33ol-bTYxYG6SFypBbBOynif2u_tgS-VRgchEiwF2g0WOIRn5P9JwFZJNxX4_5r2zCFRHhpsEVLnLCNjjR4jhCgicS-spyJzDzG1QR_E0KO3aiQusY7ZdXSJnOt1WTfSC-DjQfFaLLsHceaYqOl4VBX1k_fZ1Mo6CvSoHCsfQA6mgym0cs1y8hK7ZFMjEf1OikbXCJoQgl3tCnnOt--yYuls3tDDlpxIRpfWg0yRZaeAIogrXhnID4drVIkd1HqoIU__nQIez7VmEshDU_-Bipci3yJ4-iwEjiAejL4N6HE5L-vrz_MB5QM6B42GFNCSgH4Lletr3ON_agKUHGxoHCdNTJ8MYIhMBTwZkyz96crzyP01M-g3qsbaP33CMcb12PE-yS0AhDXwJEOlqWUsoI6fZsoovRXLpoZJfH9RONu0NiR_5fJXdUILaHtG2875sCHQ_jd41hEWxjg3jNYYhYFsnkDy4WQeWO5I1VJFgZ09xhTurfHDeQw9k9tQUgJ_ijWhyXtYEndTHO8wrQj6FF7rNy8tdPww3Si9iHNU125Y-u_V4XULkekgT3zPwef1aYJVUKMsGk_huXGLTEbfar9oYdeK59MxrBZ7ioJdTnSIN9SG-knNU381aha5l6bj2TpK3b2hYsCHLq_GK7p1vh2I2IQLqGj3KUpRKhP5vSHkfsSWMg_HPS_4M5nFYtAStvL7wnTLovO27mPsZMhTdhZ1aAz_T8AZYFr47EvKG4E-IiT1siY_iQjgSoCgqCdqOT__aN0IqwDuMiTCP2etAD5d8ZVSg4fa-NGT1D2dwbZp6UESW5tGhwpxmyn4KeXA454o-8IPFbal0b_O0xXNnJJLz79q3OT9qFZB0w-l-WGM7EqadrbM7ZryUiEfcR2g9FbDDrZCnR0joft-FINq2G1IhouDc3HToL7EJt1gOKnNs4qMBmj-QUKE1WKXuMUGeMu0JtOK3iTdShMH7u55cXDLmmXZyAKdDKXffdUWExMiahDdBCo4TDpoFYrH856eckeFSQKNkpOBbfcesQ4cIcQth4ZHWF35P4QCesQ5uEK2fIa0VwuToyTv5YG1OYO9v5lm9lRjD622pPRSIehJAHJp6hk2tUM6DeRs8PyL1eL_DpP4Onw6ITCVr8nyBgIq4B-N0C5V-uXM-JZGUHdl2-yFAAw4bRP7SYexbb5epGkjyxbwDbaYQKsVKrFfhmfpQC1XRN9lDxNqvXuoIlCl0cm79lur-pN_FXU973ukh5-vCDvuSUBNDcLzp4-0W5kg0QZucYc76XJW27PDcT03If29PEo58MefQgWy2x3BnxXaCqe__mOZFJKs8FU7meAEoVIUe-40NloKStOCkSh2VRhzxB6vDAVwMTJrL5sUzSV1qCi0ggCJmW4ECTOWSdVyAsafeRSrkNZLVrnFSEiHgNP3wK97_O9TB3g_1KpTCRauD_f6BDynVBpAAaeLRohkzMgtQBK69QM5GYPxZG-i6tXon7lHJA7Xu-9kvl_ZLc_irxlfHA8yl4na97nALxYhoCXEvku97YGp0MituMlMHPma_jkvnoDnOT2-5gJ9YrC951wgcjJYkV647rPgMkOo9iVI-cbQD--wZigs-d2Bvlf-ay6jvhaP_xbKjNfI11MAeh6sm-I4jgbz4Q7xkBJ9SgKNAgTFn8zWH8GJKNCb2tTYEmPHl8g58QuD-6T4s5sfcumXKxMKz_0QE9vZ_P7W99joCiN3xOAix5ajX4AS8Agsa3J01rEb3A3JQDRxwTFHzwmpjgdnzwiJpeXAIpiwJJkVbW-TIk9T7g")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("x-recaptcha-action", "VERIFY_PASSWORD")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Origin", "https://account.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://account.samsung.com/accounts/v1/samsung_com_us/signInPassword")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// req.Header.Set("Cookie", `s_fpid=ab255e68-baf6-4254-b15b-199266937e1f; _common_physicalAddressText=ys4zlkcmfhk7bccfbica; country_region=CA-ON; country_codes=us; device_type=pc; at_check=true; AMCVS_48855C6655783A647F000101%40AdobeOrg=1; spr-chat-token-60c1edc94fe1cd452ceb20ba_app_950960=; iadvize-6528-vuid=0bdc300b29204db3900c50cb0451a4a13dbd48ce5a784; sa_did=DKUjpiqZfYmMasIuLLaEcdNXiPdnQbBQ; sa_773-397-549898={"US":{"status":"GRANTED","updated":"2025-02-07T19:54:31.039Z","clientId":"kv5di1wr19","deviceId":"DKUjpiqZfYmMasIuLLaEcdNXiPdnQbBQ"}}; ecom_vi_USA=NTk3MzYyYWEtYjBlZi00MmM0LTlkNjYtNjg5YTU4NjEzNWEw; ecom_session_id_USA=M2FhMTFkOGQtNzg4NC00NmU2LTk2MGMtMWQwNmYxODU1ODRm; eddzipcode=80419; s_ecom_scid=5dc303b2-508c-4d31-a057-b47fa17369aa; s_ecom_sc_cnt=1; TS0171c81e=01643b4577b76f17255c491c580d87a263fa4c49500144c52d312f7cbc0d9059349a1332468cbf8d5f7d656895ef0ad3ed4373b325; page_state=Cart - Mixed Product - EIP; AKA_A2=A; __COM_SPEED=H; OptanonConsent=isGpcEnabled=0&datestamp=Mon+Feb+10+2025+10%3A28%3A35+GMT-0700+(Mountain+Standard+Time)&version=202307.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=9877b8cb-05ab-4703-9ab3-2b397c5ccfe1&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0003%3A1%2CC0002%3A1%2CBG17%3A1%2CC0004%3A1&AwaitingReconsent=false; glbState=guhA6OWTabWWF38fhkVaOQTp97OWb9GQ; returnURL=https://www.samsung.com/us/smartphones/galaxy-z-fold6/buy/galaxy-z-fold6-256gb-unlocked-sm-f956uakaxaa/; mbox=session#9b8989073c8047f696b594b6b09ee412#1739210377; s_pers=%201%3D1%7C1739210316443%3B%20first_page_visit%3Dhttps%253A%252F%252Fwww.samsung.com%252Fus%252Fsmartphones%252Fgalaxy-z-fold6%252Fbuy%252Fgalaxy-z-fold6-256gb-unlocked-sm-f956uakaxaa%252F%7C1739210316445%3B%20s_nr%3D1739208516445-Repeat%7C1741800516445%3B%20gpv_pn%3Dhttps%253A%252F%252Fwww.samsung.com%252Fus%252Fsmartphones%252Fgalaxy-z-fold6%252Fbuy%252Fgalaxy-z-fold6-256gb-unlocked-sm-f956uakaxaa%7C1739210316446%3B%20s_fbsr%3D1%7C1739210316448%3B; AMCV_48855C6655783A647F000101%40AdobeOrg=1585540135%7CMCIDTS%7C20130%7CMCMID%7C01616498313666896292237204817809599844%7CMCAID%7CNONE%7CMCOPTOUT-1739215716s%7CNONE%7CvVersion%7C4.4.0; platform=Mac OS; utag_main=v_id:0194e1fa32fd000d4ee931b3c74205075006706d00942samsung_live$_sn:2$_se:7$_ss:0$_st:1739210316535$tapid_reset:true%3Bexp-1770494064882$dc_visit:2$ses_id:1739208513358%3Bexp-session$_pn:1%3Bexp-session$dc_event:1%3Bexp-session$_prevpage:%3Bexp-1739212116536$adobe_mcid:01616498313666896292237204817809599844%3Bexp-session$aa_vid:%3Bexp-session; s_sess=%20s_sq%3D%3B%20c_m%3DundefinedTyped%252FBookmarkedTyped%252FBookmarkedundefined%3B%20s_ppvl%3Dnew%252520ecom%252520step2%25257Ccheckout%252C26%252C26%252C1002.5%252C454%252C684%252C1440%252C900%252C2%252CP%3B%20s_cc%3Dtrue%3B%20s_ppv%3Dhttps%25253A%252F%252Fwww.samsung.com%252Fus%252Fsmartphones%252Fgalaxy-z-fold6%252Fbuy%252Fgalaxy-z-fold6-256gb-unlocked-sm-f956uakaxaa%252C7%252C7%252C684%252C1440%252C684%252C1440%252C900%252C2%252CL%3B; USAWSIAMSESSIONID=GhIGjsyK5hW2dUD2eTNyuFkccjKhfYlxbWlwwQ0RCcOVuvSr; sa_did=DKUjpiqZfYmMasIuLLaEcdNXiPdnQbBQ; sa_session_extendable=false; G_ENABLED_IDPS=google; JSESSIONID=731CC6D4C06310EA8F1BAE71EA658E06`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
