package  utils

import (
	"crypto/md5"
	"encoding/hex"
	"safer.com/m/internal/env"
	"time"

	//	"encoding/json"
//	"fmt"
	"github.com/golang-jwt/jwt"
//	"safer.com/m/internal/env"
	. "safer.com/m/models"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateJWT(user User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":user.Id,
		"role":user.Role,
		"expires": time.Now().Local().Add(time.Hour * time.Duration(24*5)),
	})
	tokenString, err := token.SignedString([]byte(string(env.Cfg["SECRET_KEY"])))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}