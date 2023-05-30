package handler

import (
	"net/http"
	// "time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/panisin000/course/gorm"
	"github.com/panisin000/course/model"
	"github.com/panisin000/course/util"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 12
)

type AuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(AuthReq)
		if err := c.BindJSON(req); err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), cost)
		if err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}

		user := model.User{
			Username: req.Username,
			Password: string(hash),
		}
		if err := db.CreateUser(&user); err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}

		token, err := generateToken(user.ID)
		if err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(AuthReq)
		if err := c.BindJSON(req); err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		found, err := db.GetUserByUsername(req.Username)
		if found == nil || err != nil {
			util.SendError(c, http.StatusUnauthorized, err)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(req.Password))
		if err != nil {
			util.SendError(c, http.StatusUnauthorized, err)
			return
		}
		token, err := generateToken(found.ID)
		if err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
