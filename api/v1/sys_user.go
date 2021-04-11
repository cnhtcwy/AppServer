package v1

import (
	"cnhtc/gin-vue-admin/AppServer/global"
	"cnhtc/gin-vue-admin/AppServer/middleware"
	"cnhtc/gin-vue-admin/AppServer/model"
	"cnhtc/gin-vue-admin/AppServer/model/request"
	"cnhtc/gin-vue-admin/AppServer/model/response"
	"cnhtc/gin-vue-admin/AppServer/service"
	"cnhtc/gin-vue-admin/AppServer/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)
//type SysUser struct {
//	//global.GVA_MODEL
//	UUID        string    `json:"uuid" db:"uuid" gorm:"comment:用户UUID"`
//	Username    string       `json:"userName" db:"username" gorm:"comment:用户登录名"`
//	Password    string       `json:"-" db:"password" gorm:"comment:用户登录密码"`
//	NickName    string       `json:"nickName" db:"nick_name" gorm:"default:系统用户;comment:用户昵称" `
//	HeaderImg   string       `json:"headerImg" db:"header_img" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
//	AuthorityId string       `json:"authorityId" db:"authority_id" gorm:"default:888;comment:用户角色ID"`
//}

type SysUser struct {
	UUID     string `json:"uuid" db:"uuid"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	NickName string `json:"nickname" db:"nick_name"`
}

func Login(c *gin.Context)  {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	fmt.Println(L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		//验证码匹配之后验证密码正确性
		var user SysUser
		err := global.GVA_DB.Get(&user,"select uuid,username,password,nick_name from sys_users where username = ? and password = ?",L.Username,L.Password)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		} else {
			//tokenNext(c, *user)
			response.OkWithData(user,c)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

//登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.GVA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := service.GetRedisJWT(user.Username); err == redis.Nil {
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}