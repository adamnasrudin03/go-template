package helpers

import (
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

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	configs := configs.GetInstance()

	tokenString, err := ExtractToken(ctx)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("failed unexpected signing method : %v \n", err)
			err = NewError(ErrUnauthorized, NewResponseMultiLang(
				MultiLanguages{
					ID: fmt.Sprintf("metode penandatanganan yang tidak terduga: %v", token.Header["alg"]),
					EN: fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]),
				},
			))
			return nil, err
		}
		return []byte(configs.App.SecretKey), nil
	})

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

		return token.Claims.(*JWTClaims), nil
	}

	if _, ok := err.(*jwt.ValidationError); ok {
		errCustom := extractErrorJwtValidation(err, tokenString)
		if errCustom != nil {
			return nil, errCustom
		}
	}
	return nil, err

}

func extractErrorJwtValidation(err error, tokenString string) error {
	if ve, ok := err.(*jwt.ValidationError); ok {
		var errorCustom error
		errorCustom = NewError(ErrUnauthorized, NewResponseMultiLang(
			MultiLanguages{
				ID: fmt.Sprintf("Gagal melakukan parsing token, Error: %s", err.Error()),
				EN: fmt.Sprintf("Failed to parse token, Error: %s", err.Error()),
			},
		))
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errorCustom = NewError(ErrUnauthorized, NewResponseMultiLang(
				MultiLanguages{
					ID: fmt.Sprintf("Token format: %s tidak valid", tokenString),
					EN: fmt.Sprintf("Invalid token format: %s", tokenString),
				},
			))
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			errorCustom = NewError(ErrUnauthorized, NewResponseMultiLang(
				MultiLanguages{
					ID: "Token telah kadaluarsa",
					EN: "Token has been expired",
				},
			))
		}

		return errorCustom
	}

	log.Printf("unknown token error : %v \n", err)
	err = NewError(ErrUnauthorized, NewResponseMultiLang(
		MultiLanguages{
			ID: "Token tidak valid",
			EN: "Invalid token",
		},
	))
	return err

}
