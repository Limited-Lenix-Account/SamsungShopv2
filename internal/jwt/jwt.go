package jwt

import (
	"fmt"
	"time"
)

func MakeEcommJWT(appid string) (string, error) {

	// var jwt string

	// defaultHeaders := DefaultHeaders{
	// 	Alg:  "ES256",
	// 	Type: `v3\\/mobile\\/auth`,
	// }

	// jsonHeaders, err := json.Marshal(defaultHeaders)
	// if err != nil {
	// 	return "", fmt.Errorf("error mashaling headers: %w", err)
	// }

	// b64Headers := base64.StdEncoding.EncodeToString(jsonHeaders)
	// fmt.Println("Target jwt header: eyJhbGciOiJFUzI1NiIsInR5cGUiOiJ2M1xcXFwvbW9iaWxlXFxcXC9hdXRoIn0=")
	// fmt.Println("Actual jwt header:", b64Headers)
	// jwt += b64Headers + "."

	// payload := JwtPayload{
	// 	AppVersion:       "2.0.35059",
	// 	UniqueUserID:     "cream",
	// 	UniqueAppId:      "0c371750-dbf1-42a9-824d-1e59e852d6ee",
	// 	StoreArbitration: true,
	// 	SignedAt:         1740859994307,
	// 	KeyVersion:       "v1",
	// }
	t := time.Now().UnixMilli()
	payload := JwtPayload{
		AppVersion:       "2.0.35059",
		UniqueUserID:     "cream",
		UniqueAppId:      appid,
		StoreArbitration: true,
		SignedAt:         t,
		KeyVersion:       "v1",
	}

	// jsonPayload, err := json.Marshal(payload)
	// if err != nil {
	// 	return "", fmt.Errorf("cannot marshal jwt payload: %w", err)
	// }
	// b64Payload := base64.StdEncoding.EncodeToString(jsonPayload)
	// // fmt.Println("Target jwt header: eyJhcHBfdmVyc2lvbiI6IjIuMC4zNTA1OSIsInVuaXF1ZV91c2VyX2lkIjoiY3JlYW0iLCJ1bmlxdWVfYXBwX2lkIjoiMGMzNzE3NTAtZGJmMS00MmE5LTgyNGQtMWU1OWU4NTJkNmVlIiwic3RvcmVfYXJiaXRyYXRpb24iOnRydWUsInNpZ25lZF9hdCI6MTc0MDg1OTk5NDMwNywia2V5X3ZlcnNpb24iOiJ2MSJ9")
	// // fmt.Println("Actual jwt header:", b64Payload)
	// jwt += b64Payload + "."

	key := getECKey()
	privateKey, err := readECprivateKey(key)
	if err != nil {
		return "", fmt.Errorf("cannot read EC Private Key")
	}

	p, err := signPayload(payload, privateKey)
	if err != nil {
		return "", err
	}
	return p, nil
}
