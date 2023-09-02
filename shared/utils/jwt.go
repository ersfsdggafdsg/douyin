package utils
import (
	"time"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("鸡你太美")
var outOfDateTime = time.Hour * 1000

type Claims struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

// 签发用户token
func GenerateToken(id int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(outOfDateTime)
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "1122233",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证用户token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &Claims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return jwtSecret, nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims);
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

