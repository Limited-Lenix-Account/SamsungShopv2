import requests

cookies = {
    'country_region': 'CA-ON',
    's_fpid': '97042c2b-ec38-4fd1-a279-4da3e9fe4640',
    'AKA_A2': 'A',
    'country_codes': 'us',
    'device_type': 'pc',
    '__COM_SPEED': 'H',
    'cookie_country': 'us',
    's_ips_aep': '684',
    's_ppv_aep': 'https%253A%252F%252Fwww.samsung.com%252Fus%252F%2C5%2C5%2C684%2C1%2C18',
    'kndctr_48855C6655783A647F000101_AdobeOrg_cluster': 'or2',
    'kndctr_48855C6655783A647F000101_AdobeOrg_identity': 'CiY2MTgzMjA1MDY5NDUxNTc4MDA1MTg5NDg0MDcwMjAxOTU3NDMxOFISCI_0xYrPMhABGAEqA09SMjAA8AGP9MWKzzI=',
    'AMCV_48855C6655783A647F000101%40AdobeOrg': 'MCMID|61832050694515780051894840702019574318',
    'mboxEdgeCluster': '35',
    'mbox': 'session%2361832050694515780051894840702019574318%2DaUReFx%231739217297',
    'TEAL': 'v:0194f1517beb001b63087c6b4d3405075001b06d00942samsung_live$t:1739217237478$s:1739215437476%3Bexp-sess$sn:1$en:1',
    '_uetsid': '93562430e7e411ef8610c134ff0bbd77',
    '_uetvid': '935627c0e7e411efaf7027d6e966b67c',
    '_gcl_au': '1.1.1251132948.1739215438',
    'tapid_reset': 'true',
    'tfpsi': '9daba1e6-32fe-4018-bfdc-53929493b3a4',
    'spr-chat-token-60c1edc94fe1cd452ceb20ba_app_950960': '',
    'da_sid': 'EE1AC57A8E94AE8A1407AA13A0A7299769.0|4|0|3',
    'da_lid': 'DD29F6499AD4EA118156BB99E2A5639CDA|0|0|0',
    'da_intState': '',
    'iadvize-6528-vuid': 'e4d89f3d7abf4f049fe1d278e9954921a29dbbdb14464',
    'glbState': 'GLBk3s3jprergd',
    'returnURL': 'https%3A%2F%2Fwww.samsung.com%2Fus%2F',
    's_tp_aep': '12805',
    'OptanonConsent': 'isGpcEnabled=0&datestamp=Mon+Feb+10+2025+12%3A23%3A59+GMT-0700+(Mountain+Standard+Time)&version=202307.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=f33ff229-49cf-49e6-a4da-6e7ed0d8b6d1&interactionCount=1&landingPath=https%3A%2F%2Fwww.samsung.com%2Fus%2F&groups=C0001%3A1%2CC0003%3A1%2CC0002%3A1%2CBG17%3A1%2CC0004%3A1',
    'sa_773-397-549898': '{"US":{"status":"GRANTED","updated":"2025-02-10T19:23:59.978Z","clientId":"kv5di1wr19","deviceId":"XSRpKCMvJyDhRwTxOwuYOCEIeOBkkxzF"}}',
    '_ga_0JTZHYKZ5Z': 'GS1.1.1739215440.1.0.1739215440.60.0.0',
    '_ga': 'GA1.1.183618100.1739215440',
    'USAWSIAMSESSIONID': 'mTdsEwy38Pbv4tYNE0hgZLdz8Cf6SCD2CDZwFfyoSg14YGad',
    '_common_physicalAddressText': 'o5qzo2bsr9o3acbabjce',
    'sa_did': 'WIC6QEtk7I1rPEkFBmDLTUFHVtFd5bli',
    'JSESSIONID': '42AB98754D7299DF766BDA699A724FEC',
    'RT': '"z=1&dm=samsung.com&si=d3e6f527-ecae-4ec2-9fed-39d17359e9c7&ss=m6zfyncs&sl=2&tt=4ny&bcn=%2F%2F17de4c15.akstat.io%2F&ld=3qm&nu=kpaxjfo&cl=6m6&hd=6yp"',
    'sa_session_extendable': 'false',
    'G_ENABLED_IDPS': 'google',
    'datadome': 'mu2QXpE34GeUnhjWHUJ4i35ur5us_LWY~zhq~SBat0~lSSOjyu9KYSK_9Ed9MN2pYv15Syou7foI5sDpsWRH8EbjQ0XVdqznKztxzT539aQ0olj_uuQudNqiCRpTyJ0E',
    '_dd_s': 'rum=0&expire=1739216347029',
}

headers = {
    'Host': 'account.samsung.com',
    'Connection': 'keep-alive',
    'sec-ch-ua-platform': '"macOS"',
    'X-CSRF-TOKEN': '172fd670-f357-47a9-ab8c-03f030c76223',
    'x-recaptcha-token': '03AFcWeA4GO7LSXQveDF8Hmm89s7SUvZSmwRYVo4Dlb7npxQrFRfVcE3XJGUNXlztXP0S2tJMBVJ_14H2yiybHCPh9TGAR6CKCH9-2ZX-zVRjNeoFnCMSOLyYOfykZt7-a2q6pIZvd6UiQzr7qi_uFoK2DbAOMxU3qtkPZZdQvGx3IsxYVCGo6UJJpgTkbt2WXoFq4Cszd-63sC9WNJdEyumlEsWAo0ipR88FsCDLw5sUmxLdZFqobMH0jZWBsDS4KNeGXarqErZp29adem55pKJC4rUaYTDBVmiFW49bVx6O9X3iWSz4co9Qu6s1mrhcnKPX9yHqPGGBf4qkHyFY2Xq8cIo8OD6qXMpwXVS4kin0FaqRYN9kZxI834erEaHYg3mRTLgaYP7icXI3HIGHCMLSWVsSXeHuX41sOpsIXaModV2TiJk_430cD7uVSNAxS3RpzOS6quuBBvWtbhzs7wU3_HtnxznLBjpEjDgw8v9peH1VRYbzBmjsgJRi3HRYzCrrCTR09FS54ia1RcrdGA4CfxyZeY2iQFKHwXZ7Qle97THApSFl0ZlyKUAQHrSIzWJZCDg-uNn59XV57s7cvPmsFNxdtaOnI-X72CfsGvZPGTOVpZKyBT_0VCfVByreTXjkZoYdFoJOSagJmV9Dyh9N79wM5cd10F9PRsI622N8no_DI7g4ctm3RUIpX5smXD-jZpBVQNXwCB1W4s06JjofRkcqZu9A1t9AUQwFQ2VcVH5shoEZJVhev-J0u0A2EMX6QBpAiI2Cegt7dFqdjadm85CwGy8g9vbzo4xmaOnMeEdEopa2i9n69r9SYdIpekk9VXrxb6QwbImVAjtx3ZoCMuApV1b2SN3XJ7pR0Zz9lFh7r2E9qdAKpq7sYsct15On41i0VF6IMdE28DGlTE2zRwmuxJvnBqZLFijbE7mPlxkGyVR9E9NDnQPU_b7aUSx8Be6JxuH1_NP-I2D5kNr4GF8QBaCrceD6rXUE_cvajXj8cx7tKXioI-CMbrSS9BOOa_yQpz673EG9fLv6WdNsmB7B2g1Xc3nu3lNgqKMm-CSgOtY_A7yatSDtskwEkTi-N6V0DdqdTuzRr39lvByQuS5fN9p00FmTgloBonCLIshTW-0JX8O7XwXYdJ5HaYl69GwdMsNYljjjNwzIfGBQQrfyhZjyh72BvvPRREevfeofW9VHYLiI-zuNpw7JU23mzY-XFsbSPFqoxrlKI-UwPoprdxL-Uv2sOjJ6CgCrj0lGhllaWSAmSwvc15pz5w_EM6EMYN_IEfpGXiQhRb0EwWLYJiUbLPFa6NcdwAcyCAKPVYXCROUG4I87AeOkGWtXd8r0WC7RilfVgwt6_t2z8FVpt7SSQBOteUBDrNGGrwLlfKpMo_AbyT32YEivPmf8Y8DW9ysn5zuaOScxr6ksqiqT75oUm8UDOsVo_Fn2jBnYESKfOoseW-oEFC1NTEAYFRn-VUHGnj4-1MOn9ZbTBcmzxfk5CsUbE5Ai_hunmzaLC059il-4sd8-252a8WtaTA3xtPwpYg983luegkb47VGMyvaSF6AmQksa0Kg7DPzAMkUwkdQRSnVy8LeWJCAh1rUcQw6Yh28FppvVKcEim8gH8ArOT2UG8uECgeVOA9UaGmJOsimgZgtRBHKdoHuIyMf8FbDx6RmH01CkGjcG-znyWNKhFVvLsewFayqCxmWJQ-4RHBZiwwVKidrdpHtY-KoABXVImxjfNiyGBCOOUhnYigmf9SV2TAz4Gi0AvJv0OEk67oGORqsOE0NJ_gbxRzb_1sRXjOelQAj1T4wwxrpT5wIWzryAQWksmLWzzrMz3XJDsODK8OK-i8tO1lmo5_F55vx5MpSV14sALgwSqoptq2TEKUdxfbdDygWe2WXkxqtQx9x_kd2X2X3j1VETm9CWAWxtYsFDgBq8nFDkJS9rbKPZ0qv69DPhrBPb6ogDKTWWFZA5dP5umfRDiyzkvpSGjWeUcR3-nRQD5pkYmiKb7zDzaTlSsY4JeJScEfoqJ5NfbUjuigwU4jcnPNPrZMrglwle7G3ZI5EMR824I9jVwgucV_60Y4ESt1HhFua0cXSvHCbDKgMgS9gWCjthX1F0Jit1MMEpq5thQ53nOshLCJ2R7cbtGKK-dO_NxJaLWv2upciHUNo4rmiUAwm48uwZqyo6IUZ3kZt-cUICPiWg03AqrggBv-ReECoQCuKdZi33ksnP-fQAyz6KZZIu6Z7SATNR6KsVUith7uV6vC31CdtC1e1F7kBlYtoMXC9xPjliNbmTBVbd5xF2YFh8kEEOCyTZSIX10Mi6ndImqk6WkrP6QNkBHrJlw1HIXp07kYZtpsS5tJuq4I5RTpi1PWvRi-iuuWFTK8y1YIki-Z6ueSnXgkVcAhm-1k4hnTy4EhHMCnAUXv7JrdXnWLLAk_N-B9J7_Igo6u2CCTVPadUnGj8vH0i-yy8Fs0iOzMGB2JdDGuyZL_HYjkEFYIzLm636x9D5MZvLCTcMupq-MBefUF_fm_c4ChTxKpP_qdoROhjZ3KTbOOLOtrytsZ9d97paBfNro',
    'sec-ch-ua': '"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"',
    'sec-ch-ua-mobile': '?0',
    'x-recaptcha-action': 'VERIFY_ID',
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36',
    'Accept': 'application/json, text/plain, */*',
    'Origin': 'https://account.samsung.com',
    'Sec-Fetch-Site': 'same-origin',
    'Sec-Fetch-Mode': 'cors',
    'Sec-Fetch-Dest': 'empty',
    'Referer': 'https://account.samsung.com/accounts/v1/samsung_com_us/signInGate?response_type=code&client_id=kv5di1wr19&locale=en_US&countryCode=US&redirect_uri=https:%2F%2Fwww.samsung.com%2Faemapi%2Fv6%2Fdata-login%2FafterLogin.us.json&state=GLBk3s3jprergd&goBackURL=https:%2F%2Fwww.samsung.com%2Fus%2F&scope=',
    'Accept-Language': 'en-US,en;q=0.9',
    'Content-Type': 'application/json; charset=UTF-8',
}

params = {
    'v': '1739215447167',
}

json_data = {
    'loginId': '',
    'rememberId': False,
    'staySignIn': False,
}

response = requests.post(
    'https://account.samsung.com/accounts/v1/samsung_com_us/signInIdentificationProc',
    params=params,
    cookies=cookies,
    headers=headers,
    json=json_data,
    verify=False
)

# Note: json_data will not be serialized by requests
# exactly as it was in the original request.
#response = requests.post(
#    'https://account.samsung.com/accounts/v1/samsung_com_us/signInIdentificationProc',
#    params=params,
#    cookies=cookies,
#    headers=headers,
#    data=data,
#    proxies=proxies,
#)

print(response.text)