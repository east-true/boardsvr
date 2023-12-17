package token

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secret []byte
var AnonymousUrls map[string]bool = map[string]bool{
	"/api/login": true,
}

func init() {
	idgen, _ := uuid.NewRandom()
	id := idgen.String()
	secret = []byte(id)
}

func JwtVerify(ctx *gin.Context) {
	if _, ok := AnonymousUrls[ctx.Request.RequestURI]; ok {
		ctx.Next()
	}

	auth := ctx.Request.Header.Get("authorization")
	if strings.HasPrefix(auth, "Bearer") {
		token := strings.Split(auth, " ")[1]
		claim := new(Claims)
		if claim.Verify(token) {
			return
		}
	}

	ctx.AbortWithStatus(http.StatusForbidden)
}

type AuthToken struct {
	Access  *Claims `json:"access_token"`
	refresh *Claims
}

func NewAuthToken(role string) *AuthToken {
	now := time.Now()
	idgen, _ := uuid.NewUUID()
	id := idgen.String()
	access := NewClaims(id, role, now, 10*time.Minute)
	refresh := NewClaims(id, role, now, 1*time.Hour)
	if err := refresh.Store(); err != nil {
		fmt.Println(err)
		return nil
	}

	return &AuthToken{
		Access:  access,
		refresh: refresh,
	}
}

func (auth *AuthToken) GetTokens() (string, string, error) {
	access, err := auth.Access.Token()
	if err != nil {
		return "", "", err
	}

	refresh, err := auth.refresh.Token()
	if err != nil {
		return access, "", err
	}

	return access, refresh, nil
}

func (auth *AuthToken) Refresh() (string, error) {
	if auth.refresh == nil {
		return "", errors.New("not issued refresh token")
	}

	if auth.refresh.Expired() {
		return "", jwt.ErrTokenExpired
	}

	auth.Access = NewClaims(auth.refresh.Subject, auth.refresh.Role, time.Now(), 10*time.Minute)
	return auth.Access.Token()
}
