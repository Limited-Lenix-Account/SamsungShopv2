const crypto = require('crypto'); 

var e = {
    "payment_info": {
        "option": "AdyenCards",
        "credit_card_details": {
            "expiration_year": "2029",
            "expiration_month": "09",
            "card_number": "4444555555555555",
            "card_type": "Visa",
            "cvc": "999"
        }
    },
    "auth_info": {
        "save_payment": false
    }
}

var t = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3d3/Eodx/pdPkuYJ5Jjl\nMdzzXp6mywDLX9rUycilWPYwkefSQ2TDW2y0rrDTWHY+S6ToDKcdOdeZoBuA0wxy\neFGnqkO77xFE848/JQZ613qPQHE/bq7f/fZNLctvjuZ5ADJ17PHygc4YX6GaczKb\nHytIfBtkhSItC1faB5gl7psNFa7vSLHEQMeYX1nZI/S90DxDfk4CqY9lBOOzxEr6\nZjYbyfcQSQmh2Wfstz5ZIzpJcRnKtkrp/bX1OkBL3WPT+JESG/Sm/d0FRcvmwXUU\nTerL6q+yskKrFYUyTJvW7rhnyDLlcfEkRQo/K9GsJEKX/H8QE3qeeDOSqiRgfq57\nfQIDAQAB\n-----END PUBLIC KEY-----\n"
var n = "adyen_cards"

const o = Buffer;

function a(e, t, n) {
    let r = {
        key: t
    }
      , a = crypto.randomBytes(32)
      , s = "";
    try {
        "object" == typeof e && (s = JSON.stringify(e));
        let t = crypto.publicEncrypt(r, o.from(a)).toString("base64")
          , u = crypto.randomBytes(16)
          , c = crypto.createCipheriv("aes-256-cbc", a, u)
          , l = c.update(s)
          , d = {
            encrypted_payload: o.concat([u, l, c.final()]).toString("base64"),
            encrypted_password: t
        };
        return {
            payment_method: n,
            encrypted_payment_context: d
        }
    } catch (e) {
        console.error(e)
    }
}

console.log(a(e, t, n))