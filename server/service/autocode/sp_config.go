package autocode

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/global"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
    autoCodeReq "github.com/JoeyFrancisTribbiani/minerpserver/model/autocode/request"
)

type SpConfigService struct {
}

// CreateSpConfig 创建SpConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (spConfigService *SpConfigService) CreateSpConfig(spConfig autocode.SpConfig) (err error) {
	err = global.GVA_DB.Create(&spConfig).Error
	return err
}

// DeleteSpConfig 删除SpConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (spConfigService *SpConfigService)DeleteSpConfig(spConfig autocode.SpConfig) (err error) {
	err = global.GVA_DB.Delete(&spConfig).Error
	return err
}

// DeleteSpConfigByIds 批量删除SpConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (spConfigService *SpConfigService)DeleteSpConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SpConfig{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSpConfig 更新SpConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (spConfigService *SpConfigService)UpdateSpConfig(spConfig autocode.SpConfig) (err error) {
	err = global.GVA_DB.Save(&spConfig).Error
	return err
}

// GetSpConfig 根据id获取SpConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (spConfigService *SpConfigService)GetSpConfig(id uint) (err error, spConfig autocode.SpConfig) {
	err = global.GVA_DB.Where("id = ?", id).First(&spConfig).Error
	return
}

// GetSpConfigInfoList 分页获取SpConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (spConfigService *SpConfigService)GetSpConfigInfoList(info autoCodeReq.SpConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.SpConfig{})
    var spConfigs []autocode.SpConfig
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&spConfigs).Error
	return err, spConfigs, total
}
