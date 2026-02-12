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
    '_common_physicalAddressText': 'o5qzo2bsr9o3acbabjce',
    'sa_did': 'WIC6QEtk7I1rPEkFBmDLTUFHVtFd5bli',
    'JSESSIONID': '42AB98754D7299DF766BDA699A724FEC',
    'RT': '"z=1&dm=samsung.com&si=d3e6f527-ecae-4ec2-9fed-39d17359e9c7&ss=m6zfyncs&sl=2&tt=4ny&bcn=%2F%2F17de4c15.akstat.io%2F&ld=3qm&nu=kpaxjfo&cl=6m6&hd=6yp"',
    'sa_session_extendable': 'false',
    'G_ENABLED_IDPS': 'google',
    '_dd_s': 'rum=0&expire=1739216346025',
}

headers = {
    'Host': 'account.samsung.com',
    'Connection': 'keep-alive',
    'sec-ch-ua-platform': '"macOS"',
    'X-CSRF-TOKEN': '172fd670-f357-47a9-ab8c-03f030c76223',
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36',
    'Accept': 'application/json, text/plain, */*',
    'sec-ch-ua': '"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"',
    'sec-ch-ua-mobile': '?0',
    'Origin': 'https://account.samsung.com',
    'Sec-Fetch-Site': 'same-origin',
    'Sec-Fetch-Mode': 'cors',
    'Sec-Fetch-Dest': 'empty',
    'Referer': 'https://account.samsung.com/accounts/v1/samsung_com_us/signInGate?response_type=code&client_id=kv5di1wr19&locale=en_US&countryCode=US&redirect_uri=https:%2F%2Fwww.samsung.com%2Faemapi%2Fv6%2Fdata-login%2FafterLogin.us.json&state=GLBk3s3jprergd&goBackURL=https:%2F%2Fwww.samsung.com%2Fus%2F&scope=',
    'Accept-Language': 'en-US,en;q=0.9',
    'Content-Type': 'application/json; charset=UTF-8',
}

json_data = {
    'token': '03AFcWeA7Z63f4mF8DqhAN3Mg2AXzvXbvfzBy2aupEc-ns9uUv8J2lv8tE-aLrQ-tfIaFLbsHJTcA8avlhMnAmf3EHaNeIG9cfi2MUYdiaBf4PdVo6A7Y5chpUFbpWJM-k82-qUyVVIJzZveYxhgkpirGP3x1LyIqiUdDn5TOx-H--6vktzMYd4Pepx8_g9PwaJrfZb589YvHLFZiRa6tK4mApWZ91SKT_rxwhTEDUkZYShciGcj2r1oDJ4jBc2PkVDMKyZiL7atlIk9Zs_1VLfmOxmMQT-jh57KYUPq6D1G69qWQntzvqCEdAw0qqNPcgWuhq2Vw08C6kfbI01urOxEGHNuaB7F4iqZ5nnZ9SnsNSiWc3Rql2_PThKE-hyhm5bPNV1ZQXipjw8z5tT_6mcRNh_RxR4tpPgiQ4X3as6nnVIL4ALYcFzOIbugv-TIVs7nPpQJCKKjvitBd3A9g8YwPkT5moDANE1UU0QLD_PdeSwtRqtkwcsW8bMHqYP46U57GBxMcnrMSJJLvSvNVUgFuvlLkR6bSdLXWpnT9m_s6DO0E2WV3Ce0_ai6XrpdIzpiuJSOgTT2vc06_5Hl_Yzi-jYbuXBtuBsTmbdHDs2F3sHKTSqth4LCn7qWpFbIEzRGHMisTB0H6sy0F6eAK5ppC0w4_fDhSeouLSVcxf78Ask5J4C2Vb7e0JpqzGDS4zYQkWHbeXbu-QNN6XY0588EcG3p_zJZNSnxdR5hP1vRc8pfQu_czQP840-mx7qBsVxlUFZMareKuEHjCVnjo3xfF4fWUVyMzvUn3A2_NPHIHBy5JcBSaPq2rTxxIFvDGJDjOHzo0eGDdfZWTCy38lT4YC9QZQ5yjMF5qs-7XdneWTUUkMezc0G1xOS8PSBjGrDbqks0geeE7nqm58OCCzoqLLDhBQTrmXGJUZ5oc24S9DDxlTVyN0EwAn43xo37S-vw12GIZj-IaXCxffDLdMEfUnQBi0sfTIfxheykQGltqXvk_QTg7A0A5_Pz4Gf54V-MpZpCynbvuHqsE1GHrqAovA96ZrhfO2TMIEQLoB-YcnFVvgq7HicSmX4JRYFIE7TaLrHXwgg28_5HvABTBgJ04yhEUqvGGOF0eWqyuwNi3YdshjGrhRA75yt78KbkMwh-OBoIQ8oCIVfjUnvKQQwayQhfGRVguUYa0XcG4VEWH0MPfCQVr4QRSlbUMOniOKZ0xuPkQdIb98M9fFP-lZ4nmYLx5L8s8c6W21mwsOsfuX2WDfYi6hysBVOvHQ4Nhe3ZXZIEV2vpdFKmFkIRmYNB8IsX_54tW98zL1Vm1TKXCDPssD1Ugyz5ET2q7kNyEp2cKNvdHJ8rz0BziRW-aW3heSHVIzhRSjsPZlcCWDHbdd7YPzDnLaiD22wgu2LYYMf_mBxaRCsYFnV5sHkgiQRTt-WpLfMDtAWuyNfuHSYN-R1dc_qY3W7jA4aKYbW3FCIq3xjJewj2s7nQU2OL5KdidSZwDdhk15rRZVj_hsGXj2ALayAugOfaKXVlBKpw--Pz5wqhxTgIIsTHGdNXm4WwQgeXnnIGnH0Wlj2dk-FnWNPVJfTgxSe7Xawzt3H9LbMjeMtR9rT3jrTKJesIel0885HNkpAz7z8NbRG8aLx3x87hO86rkIRb3ujWLTtXbojbx-KJeU-7MUD4kr_UgWDOGxWrumIS3UaHa-D7-DIlXLccivsD_goq_i-hgQGLJNsB6mVzGVtt_dvmAVZB019M_TO0uFkBncTA-aEL2Hcpb-tnAMMEMiEGCLUZMO7SIWutlkri_Yb0k2iwVC-wFsd4Ql3ndVyz1QIvIqfTYtTW6ILrkn1uMEJBB10hYnjhvAnSx5rA_NoTxlEdJilNuc65qGydIu7EwMbrgqp69b1Yo-aZuND8SRKFWqKAVTxGCgJrCELnUcaniBDI3P5iaAgrnzq6MgqO7IYXYQFBl3kmYuq354vRlFXFWaAlRiWSNLPKCSeeyg5f8J7o2AFk9-t54YEpiErvj2aWDO6hbQZwoVSEyL4570FhCQx_D9vpS4NFQD7LQSHe5cQOI2Nqhz7sRO3HVIozmpHIYbGZWMmqNdDorNE_uoYtQRpdHXlm7HeVKiK2_-NnWnv2SMRp7Dw7q59dsw6UbZJouQ109MCnpX8QfMUf9lUe042rLehInl_emi-NXZYHH6pi3lsIpEfwZhXMGONqqVJccBRkuYGZOMmXVvWMHVeQwUBUEXIGDeiC3ZdO1_l8sZXdmlvKA0sprtawh_XTvEaaWf4pRogCruEOwL72Ck5fC-H1u3ac-V9R1gYSrQEhDr_xRr3HH20IcOjdVGnTFeXNvh0OfG2CtJ1jdjnBZmiyPn4BvIGCfXmD-jw2YFU6uFGUkGHiQI2DQr2r4YWstkJ66RWH5jJ1nRJRuag29WHVTztAPVSEqbgXvT29p6YRIZWbD4PCi2GTwo3vsIDIaSlvQfFsfg-LDH5vjvcghOVYnUc9QOpdTTt4ORISZERxWr-fRS5bVA64s6AWY305LKHHwGv6ImvNe3Nt7bI3kulgH32fBUehFes78PKGHBWJ4zLrAC57oiTeJqsaRRADe_5ElP1AOlVzXQIo5vhBnV2b4Fo_d1QfVIy0FLaiq3KxuvZNpZaTz_o3VPGSzIR8dcsV4XYqAzRvkTL_bciST_CqzEZDTJ3CeaaofaN81tiQeeMkYNdtgv1KjgeHy_pqBdsElH3yuwzd7xnN6kViPqVDjAt5Aqf9IM9a8gbIwhjS6-1p-xotcpNsmeKyCqU2OuU0Y-EnA9yMeMwpUjGEoGZ-VTQSpfGSp7LS2-kk2l2SrPb1GalPNdDbPwjXfDm7JNdltJ73kFOLV8C1y-7Qy1YKpzTqNi9Fh6LdfncjMnR5LbnZdStoh_h4NQBv-Iczv2tQ',
    'action': 'login',
    'username': 'lenix15@iibmp.com',
}

response = requests.post(
    'https://account.samsung.com/accounts/v1/samsung_com_us/recaptcha/assessment',
    cookies=cookies,
    headers=headers,
    json=json_data,
    verify=False
)

# Note: json_data will not be serialized by requests
# exactly as it was in the original request.
#response = requests.post(
#    'https://account.samsung.com/accounts/v1/samsung_com_us/recaptcha/assessment',
#    params=params,
#    cookies=cookies,
#    headers=headers,
#    data=data,
#    proxies=proxies,
#)

print(response.text)