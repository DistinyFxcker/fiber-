package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"

	"net/http"
)

// 跨域
func MiddlewareCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization,userID,token, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//  允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                   //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}

//管理员身份检验
func MiddlewareNeedAdmin(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Next()
		//return
		var (
			tokenString string
		)
		//获取token
		tokenString = c.GetHeader("Authorization")
		if len(tokenString) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.Response{Message: "用户未登录"})
			return
		}
		//token是否存在缓存
		if _, existed := localCache.Get(tokenString); !existed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.Response{Message: "用户令牌检验失败,请重新登录"})
			return
		}
		//校验token合法性
		tokenID, ok := token.CheckToken(tokenString)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.Response{Message: "用户令牌检验失败,请重新登录"})
			return
		}
		//字符串切割，合法格式:角色@用户id@用戶名 => 1@dbsdsfasfsa-adasdasas-ada3e33333@admin
		tokenStrArry := strings.Split(tokenID, "@")
		if len(tokenStrArry) != 5 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.Response{Message: "用户令牌检验失败,请重新登录"})
			return
		}
		//校验身份
		roles, _ := strconv.Atoi(tokenStrArry[0])
		switch roles {
		case user.Role_管理员:
			break
		case user.Role_普通会员:
			if user.Role_普通会员 != role {
				c.AbortWithStatusJSON(http.StatusUnauthorized, handler.Response{Message: "暂无权限"})
				return
			}
			break
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.Response{Message: "非法的身份"})
			return
		}
		c.Set("role", tokenStrArry[0])
		c.Set("operatorId", tokenStrArry[1])
		c.Set("nike", tokenStrArry[2])
		c.Set("primaryDomainId", tokenStrArry[3])
		c.Set("userGroup", tokenStrArry[4])
		c.Next()
	}
}
