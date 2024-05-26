package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/database"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, helpers.APIResponse(err.Error(), http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}

func AuthorizationMustBe(role []string) gin.HandlerFunc {
	isRoleValid := map[string]bool{}
	isAllRole := false
	for _, v := range role {
		isRoleValid[strings.TrimSpace(v)] = true
		if v == models.ALL {
			isAllRole = true
		}
	}

	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userEmail := userData["email"].(string)
		user := models.User{}

		err := db.Select("id", "role").Where("id = ? AND email = ?", userID, userEmail).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, helpers.APIResponse("Log in again with registered user", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		if err != nil {
			c.JSON(http.StatusUnauthorized, helpers.APIResponse("Failed to check user log in", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		if !isAllRole && !isRoleValid[user.Role] {
			c.JSON(http.StatusForbidden, helpers.APIResponse("You are not allowed to access this resources", http.StatusForbidden, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
