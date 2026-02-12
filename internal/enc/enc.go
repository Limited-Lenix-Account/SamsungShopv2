package enc

// import (
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
// 	"fmt"
// 	"strings"

// 	"golang.org/x/crypto/pbkdf2"
// )

// func getKeys() (string, string) {
// 	// Generate a keypair
// 	keypair, err := rsa.GenerateKey(rand.Reader, 2048)
// 	if err != nil {
// 		fmt.Println("Error generating keypair:", err)
// 		return "", ""
// 	}

// 	// Convert the private key to PEM format
// 	privateKeyBytes := x509.MarshalPKCS1PrivateKey(keypair)
// 	privateKeyPem := pem.EncodeToMemory(&pem.Block{
// 		Type:  "RSA PRIVATE KEY",
// 		Bytes: privateKeyBytes,
// 	})
// 	// fmt.Println("Private Key:\n", string(privateKeyPem))

// 	// Convert the public key to PEM format
// 	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&keypair.PublicKey)
// 	if err != nil {
// 		fmt.Println("Error marshalling public key:", err)
// 		return "", ""
// 	}
// 	publicKeyPem := pem.EncodeToMemory(&pem.Block{
// 		Type:  "PUBLIC KEY",
// 		Bytes: publicKeyBytes,
// 	})

// 	return string(privateKeyPem), string(publicKeyPem)
// 	// fmt.Println("Public Key:\n", string(publicKeyPem))
// }

// func sanitize(input string) string {
// 	return strings.TrimSpace(strings.ToLower(input))

// }

// func randomBytes(length int) []byte {
// 	bytes := make([]byte, length)
// 	if _, err := rand.Read(bytes); err != nil {
// 		panic(err) // should never err
// 	}
// 	return bytes
// }

// func aesEncrypt(data, key, iv []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ciphertext := make([]byte, len(data))
// 	stream := cipher.NewCFBEncrypter(block, iv)
// 	stream.XORKeyStream(ciphertext, data)

// 	return ciphertext, nil
// }

// func rsaEncrypt(publicKeyPEM string, data []byte) ([]byte, error) {
// 	publicKeyBlock, _ := pem.Decode([]byte(publicKeyPEM))
// 	if publicKeyBlock == nil {
// 		return nil, errors.New("failed to parse public key PEM")
// 	}

// 	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse public key: %v", err)
// 	}

// 	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
// 	if !ok {
// 		return nil, errors.New("not a valid RSA public key")
// 	}

// 	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, data)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to encrypt data: %v", err)
// 	}

// 	return encryptedData, nil
// }

// func EncryptLogin(username, password, oldpassword string, iterations int) (map[string]string, error) {
// 	_, aesPubKey := getKeys()
// 	result := make(map[string]string)

// 	cleanUsername := sanitize(username)
// 	cleanPassword := sanitize(password)

// 	if username == "" || password == "" {
// 		return nil, errors.New("username and password are required")
// 	}

// 	usernameHash := sha256.Sum256([]byte(cleanUsername))
// 	iv := randomBytes(16)
// 	salt := randomBytes(16)

// 	encryptionKey := pbkdf2.Key(usernameHash[:], salt, iterations, 16, sha256.New)

// 	encryptedPassword, err := aesEncrypt([]byte(cleanPassword), encryptionKey, iv)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to encrypt password: %v", err)
// 	}

// 	encryptedKey, err := rsaEncrypt(aesPubKey, encryptionKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to encrypt key: %v", err)
// 	}

// 	result["svcIptLgnID"] = username
// 	result["svcIptLgnPD"] = base64.StdEncoding.EncodeToString(encryptedPassword)
// 	result["svcIptLgnKY"] = base64.StdEncoding.EncodeToString(encryptedKey)
// 	result["svcIptLgnIV"] = hex.EncodeToString(iv)

// 	return result, nil

// }
