package service

import (
	"cnhtc/gin-vue-admin/AppServer/global"
	"cnhtc/gin-vue-admin/AppServer/model"
	"cnhtc/gin-vue-admin/AppServer/model/request"
	"cnhtc/gin-vue-admin/AppServer/utils"
	"errors"
	"github.com/jmoiron/sqlx"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

func CreateApi(api model.SysApi) (err error) {
	sqlStr := "INSERT INTO sys_apis(created_at, updated_at, path, description, api_group, method) VALUES ( now(), now(), :path, :description, :api_group, :method)"
	//if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
	//	return errors.New("存在相同api")
	//}
	var i int
	global.GVA_DB.Get(&i, "select count(*) sys_apis where path = ? AND method = ?", api.Path, api.Method)
	if i > 0 {
		return errors.New("存在相同api")
	}
	m, _ := utils.ToMapByJson(api)
	_, err = global.GVA_DB.NamedExec(sqlStr, m)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func DeleteApi(api model.SysApi) (err error) {
	sqlStr := "delete from sys_apis where id =?"
	_, err = global.GVA_DB.Exec(sqlStr, api.ID)
	ClearCasbin(1, api.Path, api.Method)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func GetAPIInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//db := global.GVA_DB.Model(&model.SysApi{})
	var apiList []model.SysApi
	sqlStr := "select * from sys_api limit ?,? "
	//if api.Path != "" {
	//
	//	db = db.Where("path LIKE ?", "%"+api.Path+"%")
	//}
	//
	//if api.Description != "" {
	//	db = db.Where("description LIKE ?", "%"+api.Description+"%")
	//}
	//
	//if api.Method != "" {
	//	db = db.Where("method = ?", api.Method)
	//}
	//
	//if api.ApiGroup != "" {
	//	db = db.Where("api_group = ?", api.ApiGroup)
	//}

	err = global.GVA_DB.Get(&total, "select count(*) from sys_apis ")
	global.GVA_DB.Select(&apiList, sqlStr, offset, limit)
	//if err != nil {
	//	return err, apiList, total
	//} else {
	//	db = db.Limit(limit).Offset(offset)
	//	if order != "" {
	//		var OrderStr string
	//		if desc {
	//			OrderStr = order + " desc"
	//		} else {
	//			OrderStr = order
	//		}
	//		err = db.Order(OrderStr).Find(&apiList).Error
	//	} else {
	//		err = db.Order("api_group").Find(&apiList).Error
	//	}
	//}
	return err, apiList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func GetAllApis() (err error, apis []model.SysApi) {
	//var apis []model.SysApi
	err = global.GVA_DB.Select(&apis, "select * from sys_apis ")
	return err, apis
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func GetApiById(id float64) (err error, api model.SysApi) {
	//err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	err = global.GVA_DB.Get(&api, "select * from sys_apis where id = ?", id)
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func UpdateApi(api model.SysApi) (err error) {
	var oldA model.SysApi
	err = global.GVA_DB.Get("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(&api).Error
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func DeleteApisByIds(ids request.IdsReq) (err error) {
	//err = global.GVA_DB.Delete(&[]model.SysApi{}, "id in ?", ids.Ids).Error
	query, args, err := sqlx.In("DELETE FROM sys_apis WHERE id IN (?)", ids)
	if err != nil {
		return err
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	//query = db.Rebind(query)
	//fmt.Println(query)
	_, err = global.GVA_DB.Exec(query, args...)
	return err
}
