package auth

import "github.com/golang-jwt/jwt/v5"

func ParseToken(token string) (*jwt.MapClaims, error) {
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || claims == nil {
		return nil, err
	}

	return claims.Claims.(*jwt.MapClaims), nil
}
