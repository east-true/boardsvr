package token

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func NewClaims(uuid, role string, now time.Time, dur time.Duration) *Claims {
	return &Claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   uuid,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(dur)),
		},
	}
}

func (c *Claims) Token() (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString(secret)
}

func (c *Claims) Verify(token string) bool {
	jwtToken, err := jwt.ParseWithClaims(token, c, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return secret, nil
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

func (c *Claims) Store() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	conn := rdb.Conn()
	defer conn.Close()

	ctx := context.Background()
	token, err := c.Token()
	if err != nil {
		fmt.Println(err)
	}

	err = conn.Set(ctx, c.Subject, token, c.ExpiresAt.Sub(time.Now())).Err()
	if err != nil {
		return err
	}

	return nil
}
