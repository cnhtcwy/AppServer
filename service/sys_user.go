package service

import (
	"cnhtc/gin-vue-admin/AppServer/global"
	"cnhtc/gin-vue-admin/AppServer/model"
	"cnhtc/gin-vue-admin/AppServer/model/request"
	"cnhtc/gin-vue-admin/AppServer/utils"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	//var user model.SysUser
	var i int
	err = global.GVA_DB.Get(&i,"select count(*) from sys_users where username=?",u.Username)
	//if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
	//	return errors.New("用户名已注册"), userInter
	//}
	if err != nil {
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	//err = global.GVA_DB.Create(&u).Error
	_,err = global.GVA_DB.Exec("insert into sys_users(created_at,updated_at,uuid,username,password,nick_name,header_img,authority_id) values(?,?,?,?,?,?,?,?)",time.Now(),time.Now(),u.UUID,u.Username,u.Password,u.NickName,u.HeaderImg,u.AuthorityId)
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Get(&user,"select * from sys_users where username =? and password =?",u.Username, u.Password)
	//err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	//var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	_,err = global.GVA_DB.Exec("update sys_users set password = ? where username = ? AND password = ? ",utils.MD5V([]byte(newPassword)),u.Username, u.Password)
	//err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var userList []model.SysUser
	err = global.GVA_DB.Select(&userList,"select * from sys_users limit ?,?",offset,limit)
	err = global.GVA_DB.Get(&total,"select count(*) from sys_users")
	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	_,err = global.GVA_DB.Exec("update sys_users set authority_id = ? where uuid =?",authorityId,uuid)
	//err = global.GVA_DB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func DeleteUser(id float64) (err error) {
	//var user model.SysUser
	_,err = global.GVA_DB.Exec("delete from sys_users where id = ?",id)
	//err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func SetUserInfo(reqUser model.SysUser) (err error, user model.SysUser) {
	_,err = global.GVA_DB.Exec("update sys_users set nick_name =?, header_img=? where id =?",reqUser.NickName,reqUser.AuthorityId,reqUser.ID)
	//err = global.GVA_DB.Updates(&reqUser).Error
	return err, reqUser
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = global.GVA_DB.Get(&u,"select * from sys_users where id = ?",id)
	//err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.GVA_DB.Get(&u,"select * from sys_users where uuid = ?",uuid); err != nil {
		return errors.New("用户不存在"), &u
	}
	//if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
	//	return errors.New("用户不存在"), &u
	//}
	return nil, &u
}
