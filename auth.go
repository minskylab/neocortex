package neocortex

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type user struct {
	Username string
}

const identityKey = "id"

func getJWTAuth(engine *Engine, secretKey string) *jwt.GinJWTMiddleware {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &user{
				Username: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			data, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				return "", errors.New(err.Error())
			}
			l := new(login)
			if l == nil {
				return nil, jwt.ErrFailedAuthentication
			}
			err = json.Unmarshal(data, l)
			if err != nil {
				return "", errors.New(err.Error())
			}

			userID := l.Username
			password := l.Password

			passwordAdmin, err := engine.getAdmin(userID)
			if err != nil || passwordAdmin == "" {
				return nil, jwt.ErrFailedAuthentication
			}
			if password == passwordAdmin {
				return &user{
					Username: userID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
