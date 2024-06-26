package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/database"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	cfg    = configs.GetInstance()
	logger = driver.Logger(cfg)
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := VerifyToken(c)
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
		authorizationMustBeValidation(c, isRoleValid, isAllRole)
		c.Next()
	}
}

func authorizationMustBeValidation(c *gin.Context, isRoleValid map[string]bool, isAllRole bool) {
	var (
		db        = database.GetDB()
		userID    = c.MustGet("id").(uint64)
		userEmail = c.MustGet("email").(string)
		user      = models.User{}
	)

	err := db.Select("id", "role").Where("id = ? AND email = ?", userID, userEmail).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || user.ID == 0 {
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
		logger.Errorf("Failed to check user log in: %v ", err)
		helpers.RenderJSON(c.Writer, http.StatusUnauthorized, helpers.NewError(helpers.ErrUnauthorized, helpers.NewResponseMultiLang(
			helpers.MultiLanguages{
				ID: "Gagal mengecek user log in",
				EN: "Failed to check user log in",
			},
		)))
		c.Abort()
		return
	}

	if !isAllRole && !isRoleValid[user.Role] {
		err = helpers.ErrCannotHaveAccessResources()
		helpers.RenderJSON(c.Writer, http.StatusForbidden, err)
		c.Abort()
		return
	}

	c.Set("role", user.Role)
}
