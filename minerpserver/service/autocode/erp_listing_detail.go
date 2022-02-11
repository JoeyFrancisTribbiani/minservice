package autocode

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/global"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	autoCodeReq "github.com/JoeyFrancisTribbiani/minerpserver/model/autocode/request"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
)

type ErpListingDetailService struct {
}

// CreateErpListingDetail 创建ErpListingDetail记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpListingDetailService *ErpListingDetailService) CreateErpListingDetail(erpListingDetail autocode.ErpListingDetail) (err error) {
	err = global.GVA_DB.Create(&erpListingDetail).Error
	return err
}

// DeleteErpListingDetail 删除ErpListingDetail记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpListingDetailService *ErpListingDetailService) DeleteErpListingDetail(erpListingDetail autocode.ErpListingDetail) (err error) {
	err = global.GVA_DB.Delete(&erpListingDetail).Error
	return err
}

// DeleteErpListingDetailByIds 批量删除ErpListingDetail记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpListingDetailService *ErpListingDetailService) DeleteErpListingDetailByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ErpListingDetail{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateErpListingDetail 更新ErpListingDetail记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpListingDetailService *ErpListingDetailService) UpdateErpListingDetail(erpListingDetail autocode.ErpListingDetail) (err error) {
	err = global.GVA_DB.Save(&erpListingDetail).Error
	return err
}

// GetErpListingDetail 根据id获取ErpListingDetail记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpListingDetailService *ErpListingDetailService) GetErpListingDetail(id uint) (err error, erpListingDetail autocode.ErpListingDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&erpListingDetail).Error
	return
}

// GetErpListingDetailInfoList 分页获取ErpListingDetail记录
// Author [piexlmax](https://github.com/piexlmax)
func (erpListingDetailService *ErpListingDetailService) GetErpListingDetailInfoList(info autoCodeReq.ErpListingDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.ErpListingDetail{})
	var erpListingDetails []autocode.ErpListingDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&erpListingDetails).Error
	return err, erpListingDetails, total
}
