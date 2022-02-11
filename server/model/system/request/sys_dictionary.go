package request

import (
	"minerpserver/model/common/request"
	"minerpserver/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
