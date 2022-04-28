package middlewares

import (
	"errors"
	"strings"

	"github.com/SermoDigital/jose/jws"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gitlab.com/trunglen/iam-api/helpers"
)

var identityKey = "id"
var jwtLog = log.WithFields(logrus.Fields{
	"middleware": "jwt",
})

func JwtMiddleware(e *casbin.Enforcer) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/health" {
			c.String(200, "I'm fine")
			c.Abort()
			return
		}
		var custoCD, err = subjectFromJWT(c)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}
		var requestURI = c.GetHeader("X-Forwarded-Uri")
		jwtLog.Info("forwarding to uri ", requestURI)
		if ok, _ := e.Enforce(custoCD, requestURI, c.Request.Method); !ok {
			jwtLog.Info("checking policy no rule apply =>> successfully, forwading to ", requestURI)
			c.Next()
			return
		}
		enforceContext := helpers.NewEnforceContext("2")
		if ok, _ := e.Enforce(enforceContext, custoCD, requestURI, c.Request.Method); ok {
			jwtLog.Info("checking policy detected rule =>> successfully, forwading to ", requestURI)
			c.Next()
		} else {
			jwtLog.Info("checking policy =>> fail, abort to ", requestURI)
			c.AbortWithError(403, errors.New("ForBidden"))
		}
	}
}

type Payload struct {
	Sub       string `json:"sub"`
	Biztype   string `json:"bizType"`
	Iss       string `json:"iss"`
	Mobile    string `json:"mobile"`
	Fullname  string `json:"fullName"`
	Bondcode  string `json:"bondCode"`
	ClientID  string `json:"client_id"`
	Custodycd string `json:"custodyCD"`
	Exp       int64  `json:"exp"`
	Iat       int64  `json:"iat"`
	Jti       string `json:"jti"`
	Email     string `json:"email"`
	Key       string `json:"key"`
	Role      string `json:"role"`
}

var (
	TokenHeadName        = "Bearer"
	ErrEmptyAuthHeader   = errors.New("auth header is empty")
	ErrInvalidAuthHeader = errors.New("auth header is invalid")
	ErrEmptyQueryToken   = errors.New("query token is empty")
	ErrCannotParseToken  = errors.New("can not parse token")
)

func jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func jwtFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)

	if token == "" {
		return "", ErrEmptyQueryToken
	}

	return token, nil
}

// var jwtKey = jwt.NewHS256([]byte("your-256-bit-secret"))

// var jwtKey = jwt.NewRS256(jwt.RSAPublicKey())

func subjectFromJWT(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		// Incorrect Authorization header format.
		return "", ErrEmptyQueryToken
	}
	token := authHeader[strings.Index(authHeader, prefix)+len(prefix):]
	// var token, _ = jwtFromQuery(c, "token")
	if token == "" {
		return "", ErrEmptyQueryToken
	}

	newToken, err := jws.ParseJWT([]byte(token))
	if err != nil {
		return "", ErrCannotParseToken
	}
	claim := newToken.Claims()
	return claim.Get("custodyCD").(string), nil
}
