package enc

// import (
// 	"bytes"
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"crypto/rsa"
// 	"crypto/sha256"
// 	"crypto/x509"
// 	"encoding/base64"
// 	"encoding/hex"
// 	"encoding/pem"
// 	"errors"
// 	"io"
// 	"strings"

// 	"golang.org/x/crypto/pbkdf2"
// )

// type WipEnc struct {
// 	PblcKyTxt     string // RSA public key in PEM format
// 	LgnEncTp      string // Encryption type
// 	PbeKySpcIters int    // PBKDF2 iterations
// }

// type Result struct {
// 	SvcIptLgnID    string
// 	SvcIptLgnPD    string
// 	SvcIptLgnKY    string
// 	SvcIptLgnIV    string
// 	SvcIptLgnOldPD string
// }

// func wipEncTpLgnUtil(r, t, e string, wipEnc WipEnc) (Result, error) {
// 	var p Result

// 	// Normalize inputs
// 	r = strings.ToLower(strings.ReplaceAll(r, " ", ""))
// 	t = strings.ReplaceAll(t, " ", "")
// 	e = strings.ReplaceAll(e, " ", "")

// 	// Check if required inputs are empty
// 	if r == "" || t == "" {
// 		return p, errors.New("invalid input")
// 	}

// 	if wipEnc.LgnEncTp == "1" {
// 		// Parse inputs
// 		n := []byte(r)
// 		i := []byte(t)

// 		// SHA256 hash of n
// 		c := sha256.Sum256(n)

// 		// Generate random salt
// 		y := make([]byte, 16)
// 		if _, err := io.ReadFull(rand.Reader, y); err != nil {
// 			return p, err
// 		}

// 		// PBKDF2 key derivation
// 		s := pbkdf2.Key(c[:], y, wipEnc.PbeKySpcIters, 32, sha256.New)

// 		// Generate random IV (16 bytes for AES-CBC)
// 		o := make([]byte, aes.BlockSize)
// 		if _, err := io.ReadFull(rand.Reader, o); err != nil {
// 			return p, err
// 		}

// 		// AES-CBC encryption with PKCS7 padding
// 		encrypted, err := aesCBCEncryptWithPKCS7(i, s, o)
// 		if err != nil {
// 			return p, err
// 		}

// 		// Base64 encode ciphertext
// 		a := base64.StdEncoding.EncodeToString(encrypted)

// 		// Hex encode IV
// 		g := hex.EncodeToString(o)

// 		// Base64 encode derived key
// 		C := base64.StdEncoding.EncodeToString(s)

// 		// RSA encryption
// 		f, err := rsaEncrypt(C, wipEnc.PblcKyTxt)
// 		if err != nil {
// 			return p, err
// 		}

// 		// Populate result
// 		p.SvcIptLgnID = r
// 		p.SvcIptLgnPD = a
// 		p.SvcIptLgnKY = f
// 		p.SvcIptLgnIV = g

// 		// Handle optional old password
// 		if e != "" {
// 			iOld := []byte(e)
// 			encryptedOld, err := aesCBCEncryptWithPKCS7(iOld, s, o)
// 			if err != nil {
// 				return p, err
// 			}
// 			aOld := base64.StdEncoding.EncodeToString(encryptedOld)
// 			p.SvcIptLgnOldPD = aOld
// 		}
// 	}

// 	return p, nil
// }

// func aesCBCEncryptWithPKCS7(plaintext, key, iv []byte) ([]byte, error) {
// 	// Add PKCS7 padding
// 	plaintext = pkcs7Pad(plaintext, aes.BlockSize)

// 	// Create AES cipher block
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create CBC mode encrypter
// 	ciphertext := make([]byte, len(plaintext))
// 	mode := cipher.NewCBCEncrypter(block, iv)
// 	mode.CryptBlocks(ciphertext, plaintext)

// 	// Return ciphertext
// 	return ciphertext, nil
// }

// // PKCS7 padding
// func pkcs7Pad(data []byte, blockSize int) []byte {
// 	padding := blockSize - (len(data) % blockSize)
// 	padText := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(data, padText...)
// }

// // RSA encryption function
// func rsaEncrypt(data, publicKeyPEM string) (string, error) {
// 	// Decode the PEM-encoded public key
// 	block, _ := pem.Decode([]byte(publicKeyPEM))
// 	if block == nil || block.Type != "PUBLIC KEY" {
// 		return "", errors.New("failed to decode PEM block containing public key")
// 	}

// 	// Parse the public key
// 	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Type assert to *rsa.PublicKey
// 	rsaPub, ok := pub.(*rsa.PublicKey)
// 	if !ok {
// 		return "", errors.New("not an RSA public key")
// 	}

// 	// Encrypt the data using RSA-OAEP with SHA-256
// 	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPub, []byte(data), nil)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Base64 encode the encrypted data
// 	return base64.StdEncoding.EncodeToString(encrypted), nil
// }

// // "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCylnSE8ANPUNPmgYJGnApUrUPQiBmTY44Lw+fQbFOOslZZnuUasDFJuPU4287/LBQEpTtgPWLmjGftG/b2sj8eTH46mvhDtE8ijgZsMnGPMmhu/AljEvNOqU6nDZDtgGmL/pAdEBtsJ/VzClv8G9bV1kvczuZtg0gt3JTH+pagEwIDAQAB"
// func EncryptLogin(email, password string) (Result, error) {
// 	wipEnc := WipEnc{
// 		PblcKyTxt: `-----BEGIN PUBLIC KEY-----
// MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCylnSE8ANPUNPmgYJGnApUrUPQ
// iBmTY44Lw+fQbFOOslZZnuUasDFJuPU4287/LBQEpTtgPWLmjGftG/b2sj8eTH4
// 6mvhDtE8ijgZsMnGPMmhu/AljEvNOqU6nDZDtgGmL/pAdEBtsJ/VzClv8G9bV1kv
// czuZtg0gt3JTH+pagEwIDAQAB
// -----END PUBLIC KEY-----`,
// 		LgnEncTp:      "1",
// 		PbeKySpcIters: 200,
// 	}

// 	result, err := wipEncTpLgnUtil(email, password, "", wipEnc)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result, err
// }
