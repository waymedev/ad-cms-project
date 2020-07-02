package jwt

import (
	"cwm.wiki/ad-CMS/initStep/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserInfo struct {
	Username string `json:"username"`
	Type     int32  `json:"type"`
}

type UserStdClaims struct {
	jwt.StandardClaims
	*UserInfo
}

func (u UserStdClaims) Valid() (err error) {

	if u.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}

	return nil

}

func GenerateToken(user *UserInfo, d time.Duration) (string, error) {
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	uClaims := UserStdClaims{
		stdClaims,
		user,
	}

	// 注意加密方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	tokenString, err := token.SignedString([]byte(global.Appkey))

	if err != nil {
		fmt.Println("It's error")
		return "", err
	}

	return tokenString, nil

}

func ParseUser(tokenString string) (*UserInfo, error) {

	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}

	claims := UserStdClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.Appkey), nil
	})

	if err != nil {
		return nil, err
	}

	if err = claims.Valid(); err != nil {
		return nil, err
	}
	return claims.UserInfo, err
}
