package helper

import "github.com/golang-jwt/jwt/v4"

type GenerateJwtStruct struct {
	Id       int
	Username string
	Password string
}

func GenerateJwtToken(g *GenerateJwtStruct, secretKey string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       g.Id,
		"username": g.Username,
		"password": g.Password,
	})

	token, err := claims.SignedString([]byte(secretKey))

	return token, err

	//claims := make(jwt.MapClaims)
	//claims["exp"] = iat + seconds
	//claims["iat"] = iat
	//claims["userId"] = g.Id
	//token := jwt.New(jwt.SigningMethodHS256)
	//token.Claims = claims
	//return token.SignedString([]byte(secretKey))
}
