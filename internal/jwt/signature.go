package jwt

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	USER_NO_LONGER_VALID = 32
	POW_2_WIDTH          = 16
)

func getECKey() string {
	// sA = {68, 68, 68, 68, 68, 43, 44, 46, USER_NO_LONGER_VALID, 39, 73, 44, 42, 73, 57, 59, USER_NO_LONGER_VALID, 63, 40, 61, 44, 73, 34, 44, 48, 68, 68, 68, 68, 68, 99, 36, 33, 10, 42, 40, 56, 44, 44, USER_NO_LONGER_VALID, 33, 34, 1, 2, 95, 19, 0, 95, 58, 47, 30, 36, 61, 36, 0, 6, 56, 92, 5, 44, 61, 14, 0, 45, 88, 90, 49, 57, 38, 66, 25, 4, 44, 95, 36, 93, 0, 90, 92, 12, 39, 30, USER_NO_LONGER_VALID, 6, 40, 6, 46, 42, 42, 24, 46, 58, 36, 93, 80, 99, 40, 30, 44, 33, 6, 60, 56, 45, 56, 14, 40, 44, 15, 19, 31, 15, USER_NO_LONGER_VALID, 35, 8, 15, 31, 11, 8, 60, 28, 14, 4, 35, 38, 1, 24, 26, 8, 13, 48, USER_NO_LONGER_VALID, 66, 62, 38, 45, 26, 37, 11, 28, 8, 19, 56, USER_NO_LONGER_VALID, 28, 70, 29, 15, 94, 8, 29, 0, 61, 81, 13, 29, 1, 94, 59, POW_2_WIDTH, 99, 62, 24, 47, POW_2_WIDTH, POW_2_WIDTH, 46, 0, 5, 95, 35, 56, 17, 89, 59, 40, 11, 93, 15, 95, 28, 88, 51, 30, 43, 90, 31, 49, 0, 80, 42, 95, 36, 94, 30, 84, 84, 99, 68, 68, 68, 68, 68, 44, 39, 45, 73, 44, 42, 73, 57, 59, USER_NO_LONGER_VALID, 63, 40, 61, 44, 73, 34, 44, 48, 68, 68, 68, 68, 68};
	// sB = {4, 4, 4, 4, 4, 2, 4, 6, 0, 6, 0, 4, 2, 0, POW_2_WIDTH, 18, 0, 22, 0, 20, 4, 0, 2, 4, POW_2_WIDTH, 4, 4, 4, 4, 4, 2, 4, 0, 2, 2, 0, POW_2_WIDTH, 4, 4, 0, 0, 2, 0, 2, 22, 18, 0, 22, 18, 6, 22, 4, 20, 4, 0, 6, POW_2_WIDTH, 20, 4, 4, 20, 6, 0, 4, POW_2_WIDTH, 18, POW_2_WIDTH, POW_2_WIDTH, 6, 2, POW_2_WIDTH, 4, 4, 22, 4, 20, 0, 18, 20, 4, 6, 22, 0, 6, 0, 6, 6, 2, 2, POW_2_WIDTH, 6, 18, 4, 20, POW_2_WIDTH, 2, 0, 22, 4, 0, 6, 20, POW_2_WIDTH, 4, POW_2_WIDTH, 6, 0, 4, 6, 18, 22, 6, 0, 2, 0, 6, 22, 2, 0, 20, 20, 6, 4, 2, 6, 0, POW_2_WIDTH, 18, 0, 4, POW_2_WIDTH, 0, 2, 22, 6, 4, 18, 4, 2, 20, 0, 18, POW_2_WIDTH, 0, 20, 6, 20, 6, 22, 0, 20, 0, 20, POW_2_WIDTH, 4, 20, 0, 22, 18, POW_2_WIDTH, 2, 22, POW_2_WIDTH, 6, POW_2_WIDTH, POW_2_WIDTH, 6, 0, 4, 22, 2, POW_2_WIDTH, POW_2_WIDTH, POW_2_WIDTH, 18, 0, 2, 20, 6, 22, 20, POW_2_WIDTH, 18, 22, 2, 18, 22, POW_2_WIDTH, 0, POW_2_WIDTH, 2, 22, 4, 22, 22, 20, 20, 2, 4, 4, 4, 4, 4, 4, 6, 4, 0, 4, 2, 0, POW_2_WIDTH, 18, 0, 22, 0, 20, 4, 0, 2, 4, POW_2_WIDTH, 4, 4, 4, 4, 4};
	var sA = []byte{68, 68, 68, 68, 68, 43, 44, 46, USER_NO_LONGER_VALID, 39, 73, 44, 42, 73, 57, 59,
		USER_NO_LONGER_VALID, 63, 40, 61, 44, 73, 34, 44, 48, 68, 68, 68, 68, 68, 99, 36,
		33, 10, 42, 40, 56, 44, 44, USER_NO_LONGER_VALID, 33, 34, 1, 2, 95, 19, 0, 95, 58,
		47, 30, 36, 61, 36, 0, 6, 56, 92, 5, 44, 61, 14, 0, 45, 88, 90, 49, 57, 38, 66,
		25, 4, 44, 95, 36, 93, 0, 90, 92, 12, 39, 30, USER_NO_LONGER_VALID, 6, 40, 6, 46, 42,
		42, 24, 46, 58, 36, 93, 80, 99, 40, 30, 44, 33, 6, 60, 56, 45, 56, 14, 40, 44, 15,
		19, 31, 15, USER_NO_LONGER_VALID, 35, 8, 15, 31, 11, 8, 60, 28, 14, 4, 35, 38, 1, 24,
		26, 8, 13, 48, USER_NO_LONGER_VALID, 66, 62, 38, 45, 26, 37, 11, 28, 8, 19, 56,
		USER_NO_LONGER_VALID, 28, 70, 29, 15, 94, 8, 29, 0, 61, 81, 13, 29, 1, 94, 59, POW_2_WIDTH,
		99, 62, 24, 47, POW_2_WIDTH, POW_2_WIDTH, 46, 0, 5, 95, 35, 56, 17, 89, 59, 40, 11, 93,
		15, 95, 28, 88, 51, 30, 43, 90, 31, 49, 0, 80, 42, 95, 36, 94, 30, 84, 84, 99, 68,
		68, 68, 68, 68, 44, 39, 45, 73, 44, 42, 73, 57, 59, USER_NO_LONGER_VALID, 63, 40, 61,
		44, 73, 34, 44, 48, 68, 68, 68, 68, 68,
	}

	var sB = []byte{
		4, 4, 4, 4, 4, 2, 4, 6, 0, 6, 0, 4, 2, 0, POW_2_WIDTH, 18, 0, 22, 0, 20, 4, 0, 2, 4,
		POW_2_WIDTH, 4, 4, 4, 4, 4, 2, 4, 0, 2, 2, 0, POW_2_WIDTH, 4, 4, 0, 0, 2, 0, 2, 22, 18,
		0, 22, 18, 6, 22, 4, 20, 4, 0, 6, POW_2_WIDTH, 20, 4, 4, 20, 6, 0, 4, POW_2_WIDTH, 18,
		POW_2_WIDTH, POW_2_WIDTH, 6, 2, POW_2_WIDTH, 4, 4, 22, 4, 20, 0, 18, 20, 4, 6, 22, 0, 6,
		0, 6, 6, 2, 2, POW_2_WIDTH, 6, 18, 4, 20, POW_2_WIDTH, 2, 0, 22, 4, 0, 6, 20, POW_2_WIDTH,
		4, POW_2_WIDTH, 6, 0, 4, 6, 18, 22, 6, 0, 2, 0, 6, 22, 2, 0, 20, 20, 6, 4, 2, 6, 0,
		POW_2_WIDTH, 18, 0, 4, POW_2_WIDTH, 0, 2, 22, 6, 4, 18, 4, 2, 20, 0, 18, POW_2_WIDTH, 0,
		20, 6, 20, 6, 22, 0, 20, 0, 20, POW_2_WIDTH, 4, 20, 0, 22, 18, POW_2_WIDTH, 2, 22, POW_2_WIDTH,
		6, POW_2_WIDTH, POW_2_WIDTH, 6, 0, 4, 22, 2, POW_2_WIDTH, POW_2_WIDTH, POW_2_WIDTH, 18, 0, 2,
		20, 6, 22, 20, POW_2_WIDTH, 18, 22, 2, 18, 22, POW_2_WIDTH, 0, POW_2_WIDTH, 2, 22, 4, 22,
		22, 20, 20, 2, 4, 4, 4, 4, 4, 4, 6, 4, 0, 4, 2, 0, POW_2_WIDTH, 18, 0, 22, 0, 20, 4,
		0, 2, 4, POW_2_WIDTH, 4, 4, 4, 4, 4,
	}

	convArr := convolve(sA, sB)

	return string(convArr)

}

func convolve(bArr []byte, bArr2 []byte) []byte {
	if bArr == nil || bArr2 == nil || len(bArr) != len(bArr2) {
		return nil
	}
	length := len(bArr)
	bArr3 := make([]byte, length)

	for i := 0; i < length; i++ {
		bArr3[i] = (bArr[i] ^ 105) | (bArr2[i] & 150)
	}
	return bArr3
}

func readECprivateKey(str string) (*ecdsa.PrivateKey, error) {

	block, _ := pem.Decode([]byte(str))

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %w", err)
	}

	return privateKey, nil
}

func signPayload(payload JwtPayload, key *ecdsa.PrivateKey) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	token.Header = map[string]interface{}{
		"alg":  "ES256",
		"type": `v3/mobile/auth`,
	}
	signedToken, _ := token.SignedString(key)

	return signedToken, nil
}
