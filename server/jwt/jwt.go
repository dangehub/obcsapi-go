package jwt

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"obcsapi-go/tools"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var MySecret = []byte(GenerateSecret()) // 生成签名的密钥

const TokenExpireDuration = time.Hour * 720

var db = &User{Id: 0, UserName: tools.ConfigGetString("user"), Password: tools.ConfigGetString("password")}

type UserInfo struct {
	Id       int
	UserName string
}
type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type MyClaims struct {
	User UserInfo
	jwt.RegisteredClaims
}

// 登录成功后调用，传入UserInfo结构体
func GenerateToken(userInfo UserInfo) (string, error) {
	//expirationTime := time.Now().Add(TokenExpireDuration) // 两个小时有效期
	expirationTime := jwt.NewNumericDate(time.Now().Add(TokenExpireDuration))
	claims := &MyClaims{
		User: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
			Issuer:    "kkbt",
		},
	}
	// 生成Token，指定签名算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	if tokenString, err := token.SignedString(MySecret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}

}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token无法解析")
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			c.Status(401)
			return
		}
		// 校验token，只要出错直接拒绝请求
		_, err := ParseToken(auth)
		if err != nil {
			c.Abort()
			log.Println("ParseToken Error:", err)
			if strings.Contains(err.Error(), "expired") {
				c.String(401, "Token expired")
				return
			}
			c.Status(401)
			return
		}
		c.Next()
	}
}

func NewInfo(user User) *UserInfo {
	return &UserInfo{Id: user.Id, UserName: user.UserName}
}

func LoginHandler(c *gin.Context) {
	var userVo User
	if c.ShouldBindJSON(&userVo) != nil {
		c.String(http.StatusOK, "参数错误")
		return
	}
	if userVo.UserName == db.UserName && userVo.Password == db.Password {
		info := NewInfo(*db)
		tokenString, err := GenerateToken(*info)
		if err != nil {
			log.Println(err)
			c.Status(500)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  201,
			"token": tokenString,
			"msg":   "登录成功",
		})
		return
	}
	c.String(400, "登录失败")
}

func GenerateSecret() string {
	md5 := md5.New()
	fileStr, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	md5.Write(fileStr)
	MD5Str := hex.EncodeToString(md5.Sum(nil))
	fmt.Println(MD5Str)
	return MD5Str
}
