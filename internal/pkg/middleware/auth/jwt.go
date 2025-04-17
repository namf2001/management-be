package auth

import (
	"github.com/gin-gonic/gin"
	"management-be/internal/model"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
)

var (
	UserNameKey = "username"
	EmailKey    = "email"
)

func NewJWTAuth() (*jwt.GinJWTMiddleware, error) {
	// Create a new JWT middleware instance
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
		Authenticator:   Authorize,
		Unauthorized:    Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	if err != nil {
		return nil, err
	}
	return jwtMiddleware, nil
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	userData, ok := data.(map[string]interface{})
	if !ok {
		return jwt.MapClaims{}
	}

	user, ok := userData["user"].(model.User)
	if !ok {
		return jwt.MapClaims{}
	}

	return jwt.MapClaims{
		jwt.IdentityKey: user.ID,
		UserNameKey:     user.Username,
		EmailKey:        user.Email,
	}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"Username":    claims["username"],
		"Email":       claims["email"],
		"UserID":      claims["identity"],
	}
}

func Authorize(c *gin.Context) (interface{}, error) {
	// Here you should implement your actual authentication logic
	return nil, jwt.ErrFailedAuthentication
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
