package request

import (
	"minerpserver/model/common/request"
	"minerpserver/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
