package request

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
)

type ErpListingDetailSearch struct {
	autocode.ErpListingDetail
	request.PageInfo
}
