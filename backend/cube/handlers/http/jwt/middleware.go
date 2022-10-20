package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

const IdentityKey = "user_id"
const IdentityFlags = "user_flags"

var ErrMissingCreds = jwt.ErrMissingLoginValues
var ErrFailedAuth = jwt.ErrFailedAuthentication

type JWTMiddleware struct {
	*jwt.GinJWTMiddleware
}

type LoginResponse struct {
	Token  string `json:"token" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDc..."`
	Expire string `json:"expire" example:"2022-03-20T17:00:01Z"`
} // @name LoginResponse

type Unauthorized struct {
	Message string `json:"message" example:"user not found"`
} // @name Unauthorized

func NewJWTMiddleware(
	authenticator func(c *gin.Context) (any, error),
	authorizator func(c *gin.Context, data any) bool,
	payload func(data any) jwt.MapClaims,
	identityHandler func(c *gin.Context) any,
	privKeyFile string,
	pubKeyFile string,
) JWTMiddleware {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "neural_storage",
		SigningAlgorithm: "RS256",
		Key:              x509.MarshalPKCS1PrivateKey(key),
		Timeout:          time.Hour,
		MaxRefresh:       24 * time.Hour,
		TokenHeadName:    "Token",
		IdentityKey:      IdentityKey,
		PayloadFunc:      payload,
		IdentityHandler:  identityHandler,
		Authenticator:    authenticator,
		Authorizator: func(data any, c *gin.Context) bool {
			return authorizator(c, data)
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			if code != http.StatusOK {
				c.AbortWithStatus(code)
			}
			c.JSON(http.StatusOK, LoginResponse{token, expire.Format(time.RFC3339)})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			if code != http.StatusOK {
				c.AbortWithStatus(code)
			}
			c.JSON(http.StatusOK, LoginResponse{token, expire.Format(time.RFC3339)})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, Unauthorized{message})
		},
		PrivKeyFile: privKeyFile,
		PubKeyFile:  pubKeyFile,
		TimeFunc:    time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return JWTMiddleware{authMiddleware}
}

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	return jwt.ExtractClaims(c)
}
