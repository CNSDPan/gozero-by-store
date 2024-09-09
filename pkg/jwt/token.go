package jwt

import "github.com/golang-jwt/jwt/v4"

// GetJwtToken
// @Desc：JWT 加解密密钥
// @param：secretKey
// @param：iat 时间戳
// @param：seconds 过期时间，单位秒
// @param：payload 数据载体
// @return：string
// @return：error
func GetJwtToken(secretKey string, iat, seconds int64, payload map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payload {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
