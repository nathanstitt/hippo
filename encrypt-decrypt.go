package hippo

import (
	"fmt"
	"errors"
	"strings"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/nathanstitt/hippo/models"
)

type UserClaims struct {
	hm.User
	jwt.StandardClaims
}

func JWTforUser(user *hm.User) (string, error) {

	claims := UserClaims{
		*user,
		jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SessionsKeyValue)
	return tokenString, err
}

func UserforJWT(tokenString string) (*hm.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SessionsKeyValue, nil
	})

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return &claims.User, nil;
	} else {
		return nil, err;
	}
}

func Encrypt(contents jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, contents)
	tokenString, err := token.SignedString(SessionsKeyValue)
	return tokenString, err
}

func userFromRequest(r *http.Request) (*hm.User, error) {
	reqToken := r.Header.Get("Authorization")
	parts := strings.Split(reqToken, "Bearer ")
	if len(parts) != 2 {
		return nil, errors.New("failed to parse Authorization Bearer token")
	}
	user, err := UserforJWT(parts[1])
	if err != nil {
		return nil, fmt.Errorf("Failed to decode auth token: %v", err);
	}
	return user, nil
}

func Decrypt(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SessionsKeyValue, nil
	})

	if err != nil {
		return nil, err;
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("failed to extract claims from token")
	}
}

func EncryptStringProperty(property string, value string) (string, error) {
	return Encrypt(map[string]interface{}{property: value});
}

func DecryptStringProperty(tokenString string, property string) (string, error) {
	claims, err := Decrypt(tokenString)
	if err == nil {
		property, ok := claims[property].(string)
		if ok {
			return property, nil
		} else {
			return "", fmt.Errorf("Unable to decode property %s", property)
		}
	}
	return "", err
}



// // Encrypts data using 256-bit AES-GCM.  This both hides the content of
// // the data and provides a check that it hasn't been altered. Output takes the
// // form of a urlencoded base64 string
// func Encrypt(plaintext string) (string, error) {
//	//plaintext []byte
//	block, err := aes.NewCipher(SessionsKeyValue)
//	if err != nil {
//		return "", err
//	}
//	gcm, err := cipher.NewGCM(block)
//	if err != nil {
//		return "", err
//	}
//	nonce := make([]byte, gcm.NonceSize())
//	_, err = io.ReadFull(rand.Reader, nonce)
//	if err != nil {
//		return "", err
//	}
//	encrypted, err := gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
//	if err != nil {
//		return "", err
//	}
//	return base64.URLEncoding.EncodeToString(encrypted), nil
// }

// // Decrypts urlencoded b64 data using 256-bit AES-GCM and returns it as as a string.
// func Decrypt(b64text string) (string, error) {
//	block, err := aes.NewCipher(SessionsKeyValue)
//	if err != nil {
//		return "", err
//	}
//	gcm, err := cipher.NewGCM(block)
//	if err != nil {
//		return "", err
//	}
//	ciphertext, err := base64.URLEncoding.DecodeString(b64text)
//	if len(ciphertext) < gcm.NonceSize() {
//		return "", errors.New("ciphertext to short")
//	}
//	bytes, err := gcm.Open(nil,
//		ciphertext[:gcm.NonceSize()],
//		ciphertext[gcm.NonceSize():],
//		nil,
//	)
//	return string(bytes[:]), err
// }
