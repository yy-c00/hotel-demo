package authorization

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secret []byte
	once sync.Once
)

//LoadCertificate to signed json web tokens
func LoadCertificate(url string) error {
	var err error
	once.Do(func() {secret, err = ioutil.ReadFile(url)})
	return err
}

//CustomClaim is a custom claim to jwt
type CustomClaim struct {
	Data []byte `json:"data"`
	jwt.StandardClaims `json:"claim"`
}

//GenerateToken generate token that contains any struct
func GenerateToken(data interface{}, duration time.Duration) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	claim := &CustomClaim {
		Data: bytes,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secret)
}

//Secret secret
func Secret() []byte {
	return secret
}