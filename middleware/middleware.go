package middleware

import (
	"errors"
	"hacktiv8-chapter-three-session-two/helpers"
	model "hacktiv8-chapter-three-session-two/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if bearerIsExist := strings.Contains(auth, "Bearer"); !bearerIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorization",
		})
		return
	}

	token := strings.Split(auth, " ")
	if len(token) < 2 {
		err := errors.New("Must provide Authorization header with format `Bearer {token}`")

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	claims, err := helpers.VerifyAccessToken(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.Set("email", claims.User.Email)
	c.Set("role", claims.User.Role)

	c.Next()
}
