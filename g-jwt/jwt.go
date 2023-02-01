package g_jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtClass struct {
}

var Jwt = JwtClass{}

func (j *JwtClass) GetJwt(privKey string, expireDuration time.Duration, payload map[string]interface{}) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privKey))
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(expireDuration).Unix()
	claims["iat"] = time.Now().Unix() // 颁发时间
	claims["payload"] = payload
	token.Claims = claims
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf("jwt generate error:%v", err)
	}
	return tokenString, nil
}

func (j *JwtClass) VerifyJwt(pubKey string, tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
		if err != nil {
			return nil, err
		}
		return verifyKey, nil
	})
	if err != nil {
		return false, fmt.Errorf("jwt verify error:%v", err)
	}
	return token.Valid, nil
}

func (j *JwtClass) VerifyJwtSkipClaimsValidation(pubKey string, tokenStr string) bool {
	parser := jwt.Parser{
		SkipClaimsValidation: true,
	}
	token, err := parser.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
		if err != nil {
			return nil, fmt.Errorf("VerifyJwtSkipClaimsValidation,err:%v", err)
		}
		return verifyKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func (j *JwtClass) DecodeBodyOfJwt(tokenStr string) map[string]interface{} {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenStr, claims, nil)
	return claims
}

func (j *JwtClass) DecodePayloadOfJwtBody(tokenStr string) map[string]interface{} {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenStr, claims, nil)
	return claims[`payload`].(map[string]interface{})
}
