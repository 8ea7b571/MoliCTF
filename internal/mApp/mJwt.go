package mApp

import (
	"time"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/golang-jwt/jwt/v5"
)

type JwtUser struct {
	ID uint `json:"id"`

	Name   string `json:"name"`
	Gender uint   `json:"gender"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`

	Username string `json:"username"`
	Active   bool   `json:"active"`

	jwt.RegisteredClaims
}

func (mapp *MApp) GenerateJwt(user *JwtUser) (string, error) {
	claims := JwtUser{
		ID:       user.ID,
		Name:     user.Name,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Username: user.Username,
		Active:   user.Active,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.MConfig.MApp.Expire) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(mapp.secret))
}

func (mapp *MApp) ParseJwt(tokenStr string) (*JwtUser, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtUser{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mapp.secret), nil
	})

	if claims, ok := token.Claims.(*JwtUser); ok && token.Valid {
		user := &JwtUser{
			ID:       claims.ID,
			Name:     claims.Name,
			Gender:   claims.Gender,
			Phone:    claims.Phone,
			Email:    claims.Email,
			Avatar:   claims.Avatar,
			Username: claims.Username,
			Active:   claims.Active,
		}
		return user, nil
	} else {
		return nil, err
	}
}
