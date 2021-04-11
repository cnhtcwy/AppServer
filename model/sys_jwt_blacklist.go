package model

import (
	"cnhtc/gin-vue-admin/AppServer/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
