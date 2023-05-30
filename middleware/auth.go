package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/panisin000/course/gorm"
	"github.com/panisin000/course/handler"
	"github.com/panisin000/course/util"
)

func RequireUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get header
		header := c.GetHeader("Authorization")
		header = strings.TrimSpace(header)
		min := len("Bearer ")
		if len(header) <= min {
			util.SendError(c, http.StatusUnauthorized, errors.New("token is require"))
			return
		}
		token := header[min:]
		claims, err := handler.VerifyToken(token)
		if err != nil {
			util.SendError(c, http.StatusUnauthorized, err)
			return
		}
		user, err := db.GetUserByID(claims.UserID)
		if user == nil || err != nil {
			util.SendError(c, http.StatusUnauthorized, err)
			return
		}
		handler.SetUser(c, user)
	}
}
