package utils

import (
	"github.com/dgrijalva/jwt-go"
	"encoding/base64"
	"crypto/md5"
	"time"
	"math/rand"
	"apiproject/database/myredis"
	"github.com/garyburd/redigo/redis"
)




//base64解码
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}


//md5值
func To_md5(encode string) (decode string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(encode))
	cipherStr := md5Ctx.Sum(nil)
	return string(base64Encode(cipherStr))
}


type Claims struct {
	Appid string `json:"appid"`
	// recommended having
	jwt.StandardClaims
}

func CreateToken(appid string) (string) {

	secret := GenerateKey(10)
	//过期时间
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claims := Claims{
		appid,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    appid,
		},
	}

	// Create the token using your claims
	c_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := c_token.SignedString([]byte(secret))

	//redis存储token
	myredis.Conn().Do("SET", signedToken, secret)

	return signedToken
}

func TokenAuth(signedToken string) (string, error) {

	secret, _ := redis.String(myredis.Conn().Do("GET", signedToken))

	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		//fmt.Println(reflect.TypeOf(claims.StandardClaims.ExpiresAt))
		//return claims.Appid, err
		return claims.Appid, err
	}
	return "", err
}


//生成随机秘钥
//@num int 生成字符串位数
func GenerateKey(mun int)(string){
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < mun; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}