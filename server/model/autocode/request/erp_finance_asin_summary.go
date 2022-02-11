package request

import (
	"minerpserver/model/autocode"
	"minerpserver/model/common/request"
)

type ErpFinanceAsinSummarySearch struct {
	autocode.ErpFinanceAsinSummary
	request.PageInfo
}
