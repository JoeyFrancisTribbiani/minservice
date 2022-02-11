package request

import (
	"minerpserver/model/common/request"
	"minerpserver/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
