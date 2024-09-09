package util

import (
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net"
)

// EnterExchange
// @Desc：统一入库金额换算,将金额 * 10000
// @param：price 只支持 int | int8 | int32 | int64 | float32 | float64
// @return：dbPrice
func EnterExchange(price interface{}) (dbPrice int64) {
	var rate int64 = 10000
	switch price.(type) {
	case int:
		newPrice, _ := price.(int)
		dbPrice = decimal.NewFromInt(int64(newPrice)).Mul(decimal.NewFromInt(rate)).IntPart()
	case int8:
		newPrice, _ := price.(int8)
		dbPrice = decimal.NewFromInt(int64(newPrice)).Mul(decimal.NewFromInt(rate)).IntPart()
	case int32:
		newPrice, _ := price.(int32)
		dbPrice = decimal.NewFromInt(int64(newPrice)).Mul(decimal.NewFromInt(rate)).IntPart()
	case int64:
		newPrice, _ := price.(int64)
		dbPrice = decimal.NewFromInt(newPrice).Mul(decimal.NewFromInt(rate)).IntPart()
	case float32:
		newPrice, _ := price.(float32)
		dbPrice = decimal.NewFromFloat(float64(newPrice)).Mul(decimal.NewFromInt(rate)).IntPart()
	case float64:
		newPrice, _ := price.(float64)
		dbPrice = decimal.NewFromFloat(newPrice).Mul(decimal.NewFromInt(rate)).IntPart()
	}
	return dbPrice
}

// OutExchange
// @Desc：统一出库金额显示转换，将 数据表 的 金额 / 10000
// @param：price
func OutExchange(dbPrice int64) (price float64) {
	var rate int64 = 10000
	price = decimal.NewFromInt(dbPrice).Sub(decimal.NewFromInt(rate)).InexactFloat64()
	return price
}

// GetServerIP
// @Desc：获取服务IP
// @return：string
// @return：error
func GetServerIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	if err != nil {
		return "", err
	}
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	ip := fmt.Sprintf("%s", ipAddress.IP.String())
	return ip, nil
}

// HashPassword
// @Desc：用户密码加密
// @param：password
// @return：string
// @return：error
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash
// @Desc：用户密码验证
// @param：password
// @param：hash
// @return：bool
// CheckPasswordHash compares a hashed password with a plain text password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("Password comparison failed:", err)
		return false
	}
	return true
}
