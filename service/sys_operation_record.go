package service

import (
	"cnhtc/gin-vue-admin/AppServer/global"
	"cnhtc/gin-vue-admin/AppServer/model"
	"cnhtc/gin-vue-admin/AppServer/model/request"
	"encoding/json"
	"github.com/jmoiron/sqlx"
)

//@author: [granty1](https://github.com/granty1)
//@function: CreateSysOperationRecord
//@description: 创建记录
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error

func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	//sqlStr := "insert into sys_operation_records (created_at,updated_at,ip,method,path,status,latency,agent,error_message,body,resp,user_id) values (?,?,?,?,?,?,?,?,?,?,?,?)"
	sqlStr := "insert into sys_operation_records (created_at,updated_at,ip,method,path,status,latency,agent,error_message,body,resp,user_id) values (now(),now(),:ip,:method,:path,:status,:latency,:agent,:error_message,:body,:resp,:user_id)"
	m := make(map[string]interface{})
	j,_ := json.Marshal(sysOperationRecord)
	json.Unmarshal(j,&m)
	_,err = global.GVA_DB.NamedExec(sqlStr, m)
	return err
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysOperationRecordByIds
//@description: 批量删除记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	query, args, err := sqlx.In("DELETE FROM sys_operation_records WHERE id IN (?)", ids)
	if err != nil {
		return err
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	//query = db.Rebind(query)
	//fmt.Println(query)
	_,err = global.GVA_DB.Exec(query, args...)
	return err
	//err = global.GVA_DB.Delete(&[]model.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	//return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteSysOperationRecord
//@description: 删除操作记录
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error

func DeleteSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	_,err = global.GVA_DB.Exec("delete from sys_operation_records where id =?",sysOperationRecord.ID)
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteSysOperationRecord
//@description: 根据id获取单条操作记录
//@param: id uint
//@return: err error, sysOperationRecord model.SysOperationRecord

func GetSysOperationRecord(id uint) (err error, sysOperationRecord model.SysOperationRecord) {
	err = global.GVA_DB.Get(&sysOperationRecord,"select * from sys_operation_records where id = ?", id)
	return err,sysOperationRecord
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysOperationRecordInfoList
//@description: 分页获取操作记录列表
//@param: info request.SysOperationRecordSearch
//@return: err error, list interface{}, total int64

func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var sysOperationRecords []model.SysOperationRecord
	//sqlStr := "select * from sys_operation_records"
	err = global.GVA_DB.Get(&total,"select count(*) from sys_operation_records")
	err = global.GVA_DB.Select(&sysOperationRecords,"select * from sys_operation_records limit ?,?",offset,limit)
	// 创建db
	//db := global.GVA_DB.Model(&model.SysOperationRecord{})
	//var sysOperationRecords []model.SysOperationRecord
	//// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Method != "" {
	//	db = db.Where("method = ?", info.Method)
	//}
	//if info.Path != "" {
	//	db = db.Where("path LIKE ?", "%"+info.Path+"%")
	//}
	//if info.Status != 0 {
	//	db = db.Where("status = ?", info.Status)
	//}
	//err = db.Count(&total).Error
	//err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
