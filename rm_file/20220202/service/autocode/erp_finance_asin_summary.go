package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ErpFinanceAsinSummaryService struct {
}

// CreateErpFinanceAsinSummary 创建ErpFinanceAsinSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpFinanceAsinSummaryService *ErpFinanceAsinSummaryService) CreateErpFinanceAsinSummary(erpFinanceAsinSummary autocode.ErpFinanceAsinSummary) (err error) {
	err = global.GVA_DB.Create(&erpFinanceAsinSummary).Error
	return err
}

// DeleteErpFinanceAsinSummary 删除ErpFinanceAsinSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpFinanceAsinSummaryService *ErpFinanceAsinSummaryService)DeleteErpFinanceAsinSummary(erpFinanceAsinSummary autocode.ErpFinanceAsinSummary) (err error) {
	err = global.GVA_DB.Delete(&erpFinanceAsinSummary).Error
	return err
}

// DeleteErpFinanceAsinSummaryByIds 批量删除ErpFinanceAsinSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpFinanceAsinSummaryService *ErpFinanceAsinSummaryService)DeleteErpFinanceAsinSummaryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ErpFinanceAsinSummary{},"id in ?",ids.Ids).Error
	return err
}

// UpdateErpFinanceAsinSummary 更新ErpFinanceAsinSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpFinanceAsinSummaryService *ErpFinanceAsinSummaryService)UpdateErpFinanceAsinSummary(erpFinanceAsinSummary autocode.ErpFinanceAsinSummary) (err error) {
	err = global.GVA_DB.Save(&erpFinanceAsinSummary).Error
	return err
}

// GetErpFinanceAsinSummary 根据id获取ErpFinanceAsinSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpFinanceAsinSummaryService *ErpFinanceAsinSummaryService)GetErpFinanceAsinSummary(id uint) (err error, erpFinanceAsinSummary autocode.ErpFinanceAsinSummary) {
	err = global.GVA_DB.Where("id = ?", id).First(&erpFinanceAsinSummary).Error
	return
}

// GetErpFinanceAsinSummaryInfoList 分页获取ErpFinanceAsinSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpFinanceAsinSummaryService *ErpFinanceAsinSummaryService)GetErpFinanceAsinSummaryInfoList(info autoCodeReq.ErpFinanceAsinSummarySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.ErpFinanceAsinSummary{})
    var erpFinanceAsinSummarys []autocode.ErpFinanceAsinSummary
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&erpFinanceAsinSummarys).Error
	return err, erpFinanceAsinSummarys, total
}
