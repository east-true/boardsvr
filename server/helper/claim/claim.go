package claim

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pborman/uuid"
)

var secret []byte

func init() {
	idgen := uuid.NewRandom()
	id := idgen.String()
	secret = []byte(id)
}

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func New(uuid, role string, dur time.Duration) *Claims {
	return &Claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dur)),
		},
	}
}

func (c *Claims) GetToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return ""
	}

	return signedToken
}

func (c *Claims) Verify(token string) bool {
	jwtToken, err := jwt.ParseWithClaims(token, c, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, jwt.ErrInvalidKeyType
	})
	if err != nil {
		return false
	}

	return jwtToken.Valid
}

func (c *Claims) Expired() bool {
	return c.ExpiresAt.Time.After(time.Now())
}
