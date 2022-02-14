package request

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
