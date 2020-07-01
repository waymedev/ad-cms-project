package jwt

import (
	"cwm.wiki/ad-CMS/initStep"
	"cwm.wiki/ad-CMS/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserStdClaims struct {
	jwt.StandardClaims
	*model.Users
}

func (u UserStdClaims) Valid() (err error) {

	if u.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}

	return nil

}

func JwtGenerateToken(user *model.Users, d time.Duration) (string, error) {
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
	tokenString, err := token.SignedString([]byte(initStep.Appkey))

	if err != nil {
		fmt.Println("It's error")
		return "", err
	}

	return tokenString, nil

}

func JwtParseUser(tokenString string) (*model.Users, error) {

	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}

	claims := UserStdClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(initStep.Appkey), nil
	})

	if err != nil {
		return nil, err
	}

	if err = claims.Valid(); err != nil {
		return nil, err
	}
	return claims.Users, err
}
