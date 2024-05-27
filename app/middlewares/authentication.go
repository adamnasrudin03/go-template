package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/database"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			helpers.RenderJSON(c.Writer, http.StatusUnauthorized, err)
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
	if len(role) == 0 {
		isAllRole = true
	}
	for _, v := range role {
		isRoleValid[strings.TrimSpace(v)] = true
		if v == models.ALL {
			isAllRole = true
		}
	}

	return func(c *gin.Context) {
		db := database.GetDB()
		userID := c.MustGet("id").(uint64)
		userEmail := c.MustGet("email").(string)
		user := models.User{}

		err := db.Select("id", "role").Where("id = ? AND email = ?", userID, userEmail).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = helpers.NewError(helpers.ErrUnauthorized, helpers.NewResponseMultiLang(
				helpers.MultiLanguages{
					ID: "Masuk kembali dengan user terdaftar",
					EN: "Log in again with registered user",
				},
			))
			helpers.RenderJSON(c.Writer, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		if err != nil {
			err = helpers.NewError(helpers.ErrUnauthorized, helpers.NewResponseMultiLang(
				helpers.MultiLanguages{
					ID: "Gagal mengecek user log in",
					EN: "Failed to check user log in",
				},
			))
			helpers.RenderJSON(c.Writer, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		if !isAllRole && !isRoleValid[user.Role] {
			err = helpers.NewError(helpers.ErrForbidden, helpers.NewResponseMultiLang(
				helpers.MultiLanguages{
					ID: "Anda tidak diizinkan untuk mengakses sumber daya ini",
					EN: "You are not allowed to access this resources",
				},
			))
			helpers.RenderJSON(c.Writer, http.StatusForbidden, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
