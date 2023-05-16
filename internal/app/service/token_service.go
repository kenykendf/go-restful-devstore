// nolint
package service

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type GenerateToken struct {
	AccessTokenKey       string
	RefreshTokenKey      string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

func NewGenerateToken(
	accessTokenKey string,
	refreshTokenKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) *GenerateToken {
	return &GenerateToken{
		AccessTokenKey:       accessTokenKey,
		RefreshTokenKey:      refreshTokenKey,
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
	}
}

func (gt *GenerateToken) GenerateAccessToken(userID int) (string, time.Time, error) {
	exp := time.Now().Add(gt.AccessTokenDuration)
	key := []byte(gt.AccessTokenKey)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", exp, err
	}

	return tokenString, exp, nil
}

func (gt *GenerateToken) GenerateRefreshToken(userID int) (string, time.Time, error) {
	exp := time.Now().Add(gt.RefreshTokenDuration)
	key := []byte(gt.RefreshTokenKey)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", exp, err
	}

	return tokenString, exp, nil

}

func (gt *GenerateToken) VerifyAccessToken(tokenString string) (string, error) {
	sub, err := gt.verify(tokenString, gt.AccessTokenKey)
	return sub, err
}

func (gt *GenerateToken) VerifyRefreshToken(tokenString string, tokenKey string) (string, error) {
	sub, err := gt.verify(tokenString, gt.RefreshTokenKey)
	return sub, err
}

func (gt *GenerateToken) verify(tokenString string, tokenKey string) (string, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(tokenKey), nil
	}

	token, err := jwt.Parse(tokenString, keyfunc)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		sub := fmt.Sprint(claims["sub"])
		return sub, nil
	}

	return "", err
}
