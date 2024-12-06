package mApp

import (
	"time"

	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/internal/mModel"
	"github.com/golang-jwt/jwt/v5"
)

type JwtAdmin struct {
	Name     string    `json:"name"`
	Gender   uint      `json:"gender"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday"`

	Username string `json:"username"`
	Active   bool   `json:"active"`

	jwt.RegisteredClaims
}

func GenerateJwtForAdmin(admin *mModel.Admin, secretKey string) (string, error) {
	claims := JwtAdmin{
		Name:     admin.Name,
		Gender:   admin.Gender,
		Phone:    admin.Phone,
		Email:    admin.Email,
		Avatar:   admin.Avatar,
		Birthday: admin.Birthday,
		Username: admin.Username,
		Active:   admin.Active,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.MConfig.MApp.Expire) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseJwtForAdmin(tokenStr, secretKey string) (*mModel.Admin, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtAdmin{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(*JwtAdmin); ok && token.Valid {
		admin := &mModel.Admin{
			Name:     claims.Name,
			Gender:   claims.Gender,
			Phone:    claims.Phone,
			Email:    claims.Email,
			Avatar:   claims.Avatar,
			Birthday: claims.Birthday,
			Username: claims.Username,
			Active:   claims.Active,
		}
		return admin, nil
	} else {
		return nil, err
	}
}
