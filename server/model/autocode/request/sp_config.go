package request

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
)

type SpConfigSearch struct{
    autocode.SpConfig
    request.PageInfo
}