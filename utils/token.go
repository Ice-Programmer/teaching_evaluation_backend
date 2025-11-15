package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
	"time"
)

var JwtKey = []byte("hdu_itmo_teaching_evaluation")

type Claims struct {
	Username string                     `json:"username"`
	ID       int64                      `json:"id"`
	Role     teaching_evaluate.UserRole `json:"role"`
	CreateAt int64                      `json:"create_at"`
	jwt.RegisteredClaims
}

func GenerateToken(expireTime time.Time, userInfo *teaching_evaluate.UserInfo) (string, error) {
	claims := &Claims{
		Username: userInfo.Name,
		ID:       userInfo.Id,
		Role:     userInfo.Role,
		CreateAt: userInfo.CreateAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "teaching_evaluation",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("token signing error: %s", err.Error())
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
