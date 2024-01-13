package lib

import (
	"fmt"
	"time"
	"backend/src/helpers"
	"github.com/golang-jwt/jwt/v5"
	// "io/ioutil"
	// "crypto/x509"
	// "crypto/rand"
	// "crypto/ecdsa"
	// "encoding/pem"
	// "crypto/elliptic"
)

const JWT_DEFAULT = "hello"

func GenerateJWT(userName string) (string, error) {
	key := helpers.GetEnvWithDefault("JWT_KEY", JWT_DEFAULT)
	token  := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(8760)).Unix() 	// TODO: Change the expiration date
	claims["authorized"] = true
	claims["user"] = userName					// NOTE: Add all information about user here that is cacheable

	s, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println("Cannot create JWT", err)
		return "", err
	}
	return s, nil
}

func IsTokenValid(tokenString string) bool {
	key := helpers.GetEnvWithDefault("JWT_KEY", JWT_DEFAULT)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return false
	}
	return true
}

func GetUser(tokenString string) string {
	key := helpers.GetEnvWithDefault("JWT_KEY", JWT_DEFAULT)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["user"].(string)
	}
	return ""
}

/* reads from a file, and returns the pem key
func returnPK() *ecdsa.PrivateKey {
	keyBytes, err := ioutil.ReadFile("config/private_key.pem")
	if err != nil {
		fmt.Printf("Error reading private key file: %v\n", err)
		return nil
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		fmt.Println("Failed to decode PEM block")
		return nil
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("Error parsing private key: %v\n", err)
		return nil 
	}
	return privateKey
} */

/* Saves the private key to a new file
func savePrivateKeyToFile(privateKey *ecdsa.PrivateKey, filePath string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
    if err != nil {
        return err
    }

    privateKeyPEM := &pem.Block{
        Type:  "EC PRIVATE KEY",
        Bytes: privateKeyBytes,
    }
    if err := pem.Encode(file, privateKeyPEM); err != nil {
        return err
    }
    return nil
} */

/* Creates a new PEM file for JWT token
func GenerateNewKey() {
	pk, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	savePrivateKeyToFile(pk, "private_key.pem")
} */

/* 
func main() {
	t, _ := GenerateJWT("sudu")
	fmt.Println(IsTokenValid(t))
	fmt.Println(GetUser(t))
}
*/