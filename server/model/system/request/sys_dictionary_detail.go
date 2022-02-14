package request

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
