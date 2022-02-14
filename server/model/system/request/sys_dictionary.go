package request

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
