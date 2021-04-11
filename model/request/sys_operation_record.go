package request

import "cnhtc/gin-vue-admin/AppServer/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
