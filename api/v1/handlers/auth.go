package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"safer.com/m/internal/env"
	"strings"
	. "safer.com/m/models"
)

func AuthenticateUserJWT(jwtToken string,roles []string) error{
	claims := jwt.MapClaims{}
	_,err:=jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Cfg["SECRET_KEY"]), nil
	})

	if err!=nil{
		return err
	}
	// ... error handling
	if len(roles) == 0{
		return nil
	}
	for _,role := range roles {
		if claims["role"] == role{
			return nil
		}
	}
	return errors.New("unauthorized")
}


func AuthMiddleWare(roles []string,next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
		if AuthenticateUserJWT(authToken,roles) !=nil{
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(HttpError{Message: "access denied"})
		}else{
			claims := jwt.MapClaims{}
			jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(env.Cfg["SECRET_KEY"]), nil
			})
			ctx := context.WithValue(r.Context(), "user",claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}