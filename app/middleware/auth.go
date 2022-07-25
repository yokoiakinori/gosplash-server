package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginUser := session.Get("loginUser")

		if loginUser == nil {
			c.String(http.StatusUnauthorized, "ログインしていません。")
			c.Abort()
		} else {
			c.Set("loginUser", loginUser)
			c.Next()
		}
	}
}