package lib

import (
	"errors"
	"github.com/pascaldekloe/jwt"
	"strconv"
	"time"
)

type JWT struct {
	Claims jwt.Claims
	Secret string
}

func NewJWT(config Config) JWT {
	var claims jwt.Claims
	//claims.Subject = strconv.FormatInt(user.ID, 10)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(
		time.Now().Add(time.Duration(config.JWT.TokenLifeTime) * time.Minute),
	)

	return JWT{
		Claims: claims,
		Secret: config.SecretKey,
	}
}

func (l JWT) GenerateToken(id int64) (string, error) {
	l.Claims.Subject = strconv.FormatInt(id, 10)

	jwtBytes, err := l.Claims.HMACSign(jwt.HS256, []byte(l.Secret))
	if err != nil {
		return "", err
	}

	return string(jwtBytes), nil
}

func (l JWT) GetSubjectFromToken(token string) (int64, error) {
	claims, err := jwt.HMACCheck([]byte(token), []byte(l.Secret))
	if err != nil {
		return 0, errors.New("invalid token")
	}

	if !claims.Valid(time.Now()) {
		return 0, errors.New("invalid time")
	}

	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
