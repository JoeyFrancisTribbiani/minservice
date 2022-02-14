package request

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
)

type ErpFinanceAsinSummarySearch struct {
	autocode.ErpFinanceAsinSummary
	request.PageInfo
}
