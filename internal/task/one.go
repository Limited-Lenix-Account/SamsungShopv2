package task

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func One() {
	client := &http.Client{}
	var data = strings.NewReader(``)
	req, err := http.NewRequest("POST", "https://www.samsung.com/us/api/v4.1/shopping-carts/d6a55db3-8407-4d76-a0e1-4c4f4764ead2/payment/apply-payment", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.samsung.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("x-ecom-app-id", "cart-spa")
	req.Header.Set("x-ecom-locale", "en-US")
	req.Header.Set("x-client-request-id", "pwa_common_33369e60-ef42-44c8-99b6-c40a587cfbcd")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	req.Header.Set("Origin", "https://www.samsung.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.samsung.com/us/web/app/checkout/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", `country_region=CA-ON; s_fpid=c634dc4b-1bb8-4294-a6f4-87501fa33472; AKA_A2=A; country_codes=us; device_type=pc; at_check=true; s_ecom_scid=d6a55db3-8407-4d76-a0e1-4c4f4764ead2; s_ecom_sc_cnt=1; ecom_vi_USA=ZGZmYjM2OTctOGY1OC00NzYyLWJmMmMtM2Y2YWE5YjhkZmEy; ecom_session_id_USA=OWFiYmE1OGEtZTRkZC00NjljLWE5MmItZDhmNjc2MGI3NWM3; s_ecid=MCMID%7C08233243831789355651510597816113923661; AMCVS_48855C6655783A647F000101%40AdobeOrg=1; __COM_SPEED=H; tracker_device_is_opt_in=true; mboxEdgeCluster=35; ab_test_show_buyLink=true; tracker_device=960b2ea4-6fd9-4bbe-ba44-64dbb46703e2; page_state=Cart - Mixed Product - EIP; kampyle_userid=3dd6-c870-9f86-8204-ca0a-926a-f534-e517; _ga=GA1.1.652919136.1739992635; _fbp=fb.1.1739992635076.154121339781625852; _gcl_au=1.1.470962812.1739992635; __attentive_session_id=b2bd9ad199d641de8c1c24a834061c1b; __attentive_id=4b1a32fc957d4e4990f9cc86299f6d5a; __attentive_cco=1739992635485; tfpsi=78c85cee-49eb-42c5-929b-b6d340f34b91; __attentive_ss_referrer=ORGANIC; __attentive_dv=1; _aeaid=78907008-02b6-479d-a1dd-a1ce86f653e3; aelastsite=IqynIxmTpWLTmucVE6VJcc93GvEXZoAScAF4P8JKEh9cCxYtr%2FMAw5Qj62MwuOMg; aelreadersettings=%7B%22c_big%22%3A0%2C%22rg%22%3A0%2C%22memph%22%3A0%2C%22contrast_setting%22%3A0%2C%22colorshift_setting%22%3A0%2C%22text_size_setting%22%3A0%2C%22space_setting%22%3A0%2C%22font_setting%22%3A0%2C%22k%22%3A0%2C%22k_disable_default%22%3A0%2C%22hlt%22%3A0%2C%22disable_animations%22%3A0%2C%22display_alt_desc%22%3A0%7D; spr-chat-token-60c1edc94fe1cd452ceb20ba_app_950960=; iadvize-6528-vuid=7a1b035c2a88480ca2137dd52ec43536e729d058b32d4; __idcontext=eyJjb29raWVJRCI6IjJ0R3lLMTFPdVBmd09IZ1lndW5iQ1BMUzRsZiIsImRldmljZUlEIjoiMnRHeUsxWE1xTzdQd0xsRjQ3TXUzQkxRNkRCIiwiaXYiOiIiLCJ2IjoiIn0%3D; sa_did=ZtCfnBzYABSCVrbxHjKlSaOUxnqRvclj; sa_773-397-549898={"US":{"status":"GRANTED","updated":"2025-02-19T19:17:20.260Z","clientId":"kv5di1wr19","deviceId":"ZtCfnBzYABSCVrbxHjKlSaOUxnqRvclj"}}; AAMC_samsungelectronicsamericainc_0=REGION%7C9; aam_test=segs%3D7431031; aam_sc=aamsc%3D4718718; aam_uuid=08205780020363145131511229012232379089; AMCV_48855C6655783A647F000101%40AdobeOrg=1585540135%7CMCIDTS%7C20139%7CMCMID%7C08233243831789355651510597816113923661%7CMCAID%7CNONE%7CMCOPTOUT-1739999906s%7CNONE%7CMCAAMLH-1740597506%7C9%7CMCAAMB-1740597506%7Cj8Odv6LonN4r3an7LhD3WZrU1bUpAkFkkiY1ncBR96t2PTI%7CvVersion%7C4.4.0%7CMCCIDH%7C-638464310; TS011dc0a2=011ef6f086f38e7d19caef2e913e3477410b8aca823ce27707ddb0287a67a63a2f7c45316a4d7bc2e2c7e0434914867bf0c0832559; TS0171c81e=011ef6f086f38e7d19caef2e913e3477410b8aca823ce27707ddb0287a67a63a2f7c45316a4d7bc2e2c7e0434914867bf0c0832559; mbox=session#cfc25a7d9f7c4fffa56ca7aa310b2f00#1739994574|PC#cfc25a7d9f7c4fffa56ca7aa310b2f00.35_0#1803237514; OptanonConsent=isGpcEnabled=0&datestamp=Wed+Feb+19+2025+12%3A18%3A35+GMT-0700+(Mountain+Standard+Time)&version=202307.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=03cb5bb5-6d5c-4c91-af34-89b53234e9d6&interactionCount=1&landingPath=https%3A%2F%2Fwww.samsung.com%2Fus%2Fweb%2Fapp%2Fcheckout%2F&groups=C0001%3A1%2CC0003%3A1%2CC0002%3A1%2CBG17%3A1%2CC0004%3A1; _ga_0JTZHYKZ5Z=GS1.1.1739992634.1.1.1739992717.53.0.0; _uetsid=20bcbc70eef611efac40c7744e0d7eb8; _uetvid=20bcf6f0eef611ef955b4191323bb5ef; _attn_=eyJ1Ijoie1wiY29cIjoxNzM5OTkyNjM1NDgxLFwidW9cIjoxNzM5OTkyNjM1NDgxLFwibWFcIjoyMTkwMCxcImluXCI6ZmFsc2UsXCJ2YWxcIjpcIjRiMWEzMmZjOTU3ZDRlNDk5MGY5Y2M4NjI5OWY2ZDVhXCJ9Iiwic2VzIjoie1widmFsXCI6XCJiMmJkOWFkMTk5ZDY0MWRlOGMxYzI0YTgzNDA2MWMxYlwiLFwidW9cIjoxNzM5OTkyNzE3MDY2LFwiY29cIjoxNzM5OTkyNzE3MDY2LFwibWFcIjowLjAyMDgzMzMzMzMzMzMzMzMzMn0ifQ==; da_sid=7C70531E8E3AAE8AD50DAA13A149DC858C.1|3|0|3; da_lid=4F43602D9A7AEA11405CBB99E34B968E3F|0|0|0; da_intState=; __attentive_pv=3; kampyleUserSession=1739992718582; kampyleUserSessionsCount=3; kampyleSessionPageCounter=1; s_pers=%201%3D1%7C1739994615964%3B%20first_page_visit%3Dhttps%253A%252F%252Fwww.samsung.com%252Fus%252Fweb%252Fapp%252Fcheckout%252F%7C1739994615964%3B%20s_nr%3D1739992815965-New%7C1742584815965%3B%20gpv_pn%3Dnew%2520ecom%2520step2%257Ccheckout%7C1739994615965%3B%20s_fbsr%3D1%7C1739994615966%3B; attntv_mstore_phone=7203819035:0; utag_main=v_id:01951fa486f5000e1b4cb141432505075004006d00942samsung_live$_sn:1$_se:22$_ss:0$_st:1739995144527$ses_id:1739992631030%3Bexp-session$_pn:3%3Bexp-session$_prevpage:%3Bexp-1739996944530$adobe_mcid:08233243831789355651510597816113923661%3Bexp-session$aa_vid:%3Bexp-session$tapid_reset:true%3Bexp-1771528634243$dc_visit:1$dc_event:5%3Bexp-session$dc_region:us-west-2%3Bexp-session; s_sess=%20c_m%3DundefinedTyped%252FBookmarkedTyped%252FBookmarkedundefined%3B%20s_ppvl%3Dnew%252520ecom%252520step2%25257Ccheckout%252C16%252C16%252C684%252C714%252C684%252C1440%252C900%252C2%252CL%3B%20s_cc%3Dtrue%3B%20s_sq%3Dsssamsungnewus%253D%252526c.%252526a.%252526activitymap.%252526page%25253Dhttps%2525253A%2525252F%2525252Fwww.samsung.com%2525252Fus%2525252Fweb%2525252Fapp%2525252Fcheckout%2525252F%252526link%25253DSamsung%25252520Homepage%25252520Become%25252520a%25252520member%25252520and%25252520earn%252525201410%25252520Samsung%25252520Rewards%25252520points%25252520on%25252520this%25252520order.%25252520Sign%25252520in%25252520%2525252F%25252520Sign%25252520up%25252520Guest%25252520checkout%25252520Zip%25252520code%2525253A%25252520%252526region%25253DBODY%252526.activitymap%252526.a%252526.c%252526pid%25253Dhttps%2525253A%2525252F%2525252Fwww.samsung.com%2525252Fus%2525252Fweb%2525252Fapp%2525252Fcheckout%2525252F%252526oid%25253Dfunctionrg%25252528%25252529%2525257B%2525257D%252526oidt%25253D2%252526ot%25253DDIV%3B%20s_ppv%3Dnew%252520ecom%252520step2%25257Ccheckout%252C27%252C67%252C2998%252C714%252C684%252C1440%252C900%252C2%252CL%3B`)
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
