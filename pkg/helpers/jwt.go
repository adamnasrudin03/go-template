package helpers

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/adamnasrudin03/go-template/app/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(params JWTClaims) (string, error) {
	configs := configs.GetInstance()
	now := time.Now().Unix()
	expired := time.Now().AddDate(0, 0, configs.App.ExpiredToken).Unix()

	claims := &JWTClaims{
		Email: params.Email,
		Name:  params.Name,
		ID:    params.ID,
		StandardClaims: jwt.StandardClaims{
			Id:        fmt.Sprintf("%d", params.ID),
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: expired,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.App.SecretKey))
	if err != nil {
		log.Printf("failed generate token: %v \n", err)
		err = NewError(ErrUnauthorized, NewResponseMultiLang(
			MultiLanguages{
				ID: "Gagal melakukan signed token string",
				EN: "Failed to get signed token string",
			},
		))
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	configs := configs.GetInstance()

	tokenString, err := ExtractToken(ctx)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			err = NewError(ErrUnauthorized, NewResponseMultiLang(
				MultiLanguages{
					ID: "Gagal melakukan sign token",
					EN: "Failed to get sign token",
				},
			))

			return nil, err
		}

		return []byte(configs.App.SecretKey), nil
	})
	if err != nil {
		err = NewError(ErrUnauthorized, NewResponseMultiLang(
			MultiLanguages{
				ID: "Token tidak valid",
				EN: "Invalid token",
			},
		))
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		err = NewError(ErrUnauthorized, NewResponseMultiLang(
			MultiLanguages{
				ID: "Gagal melakukan parse claims",
				EN: "Failed to parse claims",
			},
		))
		return nil, err
	}
	if ok && token.Valid {
		ctx.Set("id", claims.ID)
		ctx.Set("name", claims.Name)
		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)
		ctx.Set("expired_at", claims.ExpiresAt)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		var errorStr string
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errorStr = fmt.Sprintf("Invalid token format: %s", tokenString)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			errorStr = "Token has been expired"
		} else {
			errorStr = fmt.Sprintf("Token Parsing Error: %s", err.Error())
		}
		return nil, errors.New(errorStr)
	} else {
		return nil, errors.New("unknown token error")
	}

	return token.Claims.(*JWTClaims), nil
}
func ExtractToken(ctx *gin.Context) (tokenString string, err error) {
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		err = NewError(ErrUnauthorized, NewResponseMultiLang(
			MultiLanguages{
				ID: "Bearer token tidak ditemukan",
				EN: "Bearer token not found",
			},
		))
		return "", err
	}

	return headerToken[7:], nil
}
