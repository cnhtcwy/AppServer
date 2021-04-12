package service

import (
	"cnhtc/gin-vue-admin/AppServer/global"
	"cnhtc/gin-vue-admin/AppServer/model"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	sqlStr := "insert into jwt_blacklists(created_at,updated_at, jwt) values (now(),now(),?)"
	_, err = global.GVA_DB.Exec(sqlStr, jwtList.Jwt)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func IsBlacklist(jwt string) bool {
	var i int
	sqlStr := "select count(*) from jwt_blacklists where jwt =?"
	err := global.GVA_DB.Get(&i, sqlStr, jwt)
	if err != nil {
		return false
	}
	if i > 0 {
		return true
	}
	return false
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, redisJWT
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: userName string
//@return: err error, redisJWT string

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
