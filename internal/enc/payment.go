package enc

// "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3d3/Eodx/pdPkuYJ5Jjl\nMdzzXp6mywDLX9rUycilWPYwkefSQ2TDW2y0rrDTWHY+S6ToDKcdOdeZoBuA0wxy\neFGnqkO77xFE848/JQZ613qPQHE/bq7f/fZNLctvjuZ5ADJ17PHygc4YX6GaczKb\nHytIfBtkhSItC1faB5gl7psNFa7vSLHEQMeYX1nZI/S90DxDfk4CqY9lBOOzxEr6\nZjYbyfcQSQmh2Wfstz5ZIzpJcRnKtkrp/bX1OkBL3WPT+JESG/Sm/d0FRcvmwXUU\nTerL6q+yskKrFYUyTJvW7rhnyDLlcfEkRQo/K9GsJEKX/H8QE3qeeDOSqiRgfq57\nfQIDAQAB\n-----END PUBLIC KEY-----\n"
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"log"
)

type PaymentInfo struct {
	Option            string `json:"option"`
	CreditCardDetails struct {
		ExpirationYear  string `json:"expiration_year"`
		ExpirationMonth string `json:"expiration_month"`
		CardNumber      string `json:"card_number"`
		CardType        string `json:"card_type"`
		CVC             string `json:"cvc"`
	} `json:"credit_card_details"`
}

type AuthInfo struct {
	SavePayment bool `json:"save_payment"`
}

type PaymentContext struct {
	EncryptedPayload  string `json:"encrypted_payload"`
	EncryptedPassword string `json:"encrypted_password"`
}

type EncryptedPayment struct {
	PaymentMethod           string         `json:"payment_method"`
	EncryptedPaymentContext PaymentContext `json:"encrypted_payment_context"`
}

// pkcs7Pad adds PKCS7 padding to the data
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data) % blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func EncPayment(number, month, year, cvv string) EncryptedPayment {
	paymentStruct := map[string]interface{}{
		"payment_info": PaymentInfo{
			Option: "AdyenCards",
			CreditCardDetails: struct {
				ExpirationYear  string `json:"expiration_year"`
				ExpirationMonth string `json:"expiration_month"`
				CardNumber      string `json:"card_number"`
				CardType        string `json:"card_type"`
				CVC             string `json:"cvc"`
			}{
				ExpirationYear:  "20" + year,
				ExpirationMonth: month,
				CardNumber:      number,
				CardType:        "Amex",
				CVC:             cvv,
			},
		},
		"auth_info": AuthInfo{
			SavePayment: false,
		},
	}

	result, err := encryptPayment(paymentStruct, PUBLIC_KEY, PAYMENT_TYPE)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return result
}

func encryptPayment(payment_struct map[string]interface{}, pub_key string, payment_type string) (EncryptedPayment, error) {

	block, _ := pem.Decode([]byte(pub_key))
	if block == nil || block.Type != "PUBLIC KEY" {
		return EncryptedPayment{}, errors.New("failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return EncryptedPayment{}, err
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return EncryptedPayment{}, errors.New("not an RSA public key")
	}

	// Generate random AES key
	aesKey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, aesKey); err != nil {
		return EncryptedPayment{}, err
	}

	// Encrypt AES key with RSA public key
	encryptedAESKey, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, aesKey)
	if err != nil {
		return EncryptedPayment{}, err
	}

	// Convert JSON payload to string
	payload, err := json.Marshal(payment_struct)
	if err != nil {
		return EncryptedPayment{}, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return EncryptedPayment{}, err
	}

	// Encrypt payload with AES-256-CBC
	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return EncryptedPayment{}, err
	}

	mode := cipher.NewCBCEncrypter(blockCipher, iv)
	paddedPayload := pkcs7Pad(payload, aes.BlockSize)
	encryptedPayload := make([]byte, len(paddedPayload))
	mode.CryptBlocks(encryptedPayload, paddedPayload)

	finalPayload := append(iv, encryptedPayload...)


	encryptedPaymentContext := PaymentContext{
		EncryptedPayload:  base64.StdEncoding.EncodeToString(finalPayload),
		EncryptedPassword: base64.StdEncoding.EncodeToString(encryptedAESKey),
	}

	return EncryptedPayment{
		PaymentMethod:           payment_type,
		EncryptedPaymentContext: encryptedPaymentContext,
	}, nil
}
