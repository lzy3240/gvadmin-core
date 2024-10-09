package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gvadmin_v3/core/config"
	"gvadmin_v3/core/global/E"
	"gvadmin_v3/core/global/R"
	"net/http"
	"strings"
	"time"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 根据实际情况取TOKEN, 这里从request header取
		header := c.Request.Header
		tokenStr := header.Get(E.HeaderSignToken)
		if len(tokenStr) < 1 {
			R.ErrorResp(c).SetMsg("参数错误").SetCode(http.StatusInternalServerError).WriteJsonExit()
			c.Abort()
			return
		}

		token, err := VerifyToken(tokenStr)
		if err != nil {
			R.ErrorResp(c).SetMsg("认证失败").SetCode(http.StatusUnauthorized).WriteJsonExit()
			c.Abort()
			return
		}
		userId := token.Claims.(jwt.MapClaims)["user_id"]
		deptId := token.Claims.(jwt.MapClaims)["dept_id"]
		userName := token.Claims.(jwt.MapClaims)["user_name"]

		// 角色权限验证

		// 此处已经通过了, 可以把Claims中的有效信息拿出来放入上下文使用
		c.Set("userId", userId)
		c.Set("deptId", deptId)
		c.Set("userName", userName)
		c.Next()
	}
}

func CreateToken(UserName string, UserId int, DeptId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name": UserName, // 为了添加createBy 字段
		"user_id":   UserId,
		"dept_id":   DeptId,
		"exp":       time.Now().Unix() + int64(config.Instance().Jwt.Ttl),
		"iss":       "gvadmin_v3",
	})

	mySigningKey := []byte(config.Instance().Jwt.Secret)

	return token.SignedString(mySigningKey)
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	mySigningKey := []byte(config.Instance().Jwt.Secret)
	tokenStr = strings.ReplaceAll(tokenStr, E.HeaderSignTokenStr, "")
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
}
