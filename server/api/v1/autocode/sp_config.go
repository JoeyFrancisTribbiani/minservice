package autocode

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/global"
    "github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
    "github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
    autocodeReq "github.com/JoeyFrancisTribbiani/minerpserver/model/autocode/request"
    "github.com/JoeyFrancisTribbiani/minerpserver/model/common/response"
    "github.com/JoeyFrancisTribbiani/minerpserver/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SpConfigApi struct {
}

var spConfigService = service.ServiceGroupApp.AutoCodeServiceGroup.SpConfigService


// CreateSpConfig 创建SpConfig
// @Tags SpConfig
// @Summary 创建SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SpConfig true "创建SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spConfig/createSpConfig [post]
func (spConfigApi *SpConfigApi) CreateSpConfig(c *gin.Context) {
	var spConfig autocode.SpConfig
	_ = c.ShouldBindJSON(&spConfig)
	if err := spConfigService.CreateSpConfig(spConfig); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSpConfig 删除SpConfig
// @Tags SpConfig
// @Summary 删除SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SpConfig true "删除SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /spConfig/deleteSpConfig [delete]
func (spConfigApi *SpConfigApi) DeleteSpConfig(c *gin.Context) {
	var spConfig autocode.SpConfig
	_ = c.ShouldBindJSON(&spConfig)
	if err := spConfigService.DeleteSpConfig(spConfig); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSpConfigByIds 批量删除SpConfig
// @Tags SpConfig
// @Summary 批量删除SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /spConfig/deleteSpConfigByIds [delete]
func (spConfigApi *SpConfigApi) DeleteSpConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := spConfigService.DeleteSpConfigByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSpConfig 更新SpConfig
// @Tags SpConfig
// @Summary 更新SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SpConfig true "更新SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /spConfig/updateSpConfig [put]
func (spConfigApi *SpConfigApi) UpdateSpConfig(c *gin.Context) {
	var spConfig autocode.SpConfig
	_ = c.ShouldBindJSON(&spConfig)
	if err := spConfigService.UpdateSpConfig(spConfig); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSpConfig 用id查询SpConfig
// @Tags SpConfig
// @Summary 用id查询SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SpConfig true "用id查询SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /spConfig/findSpConfig [get]
func (spConfigApi *SpConfigApi) FindSpConfig(c *gin.Context) {
	var spConfig autocode.SpConfig
	_ = c.ShouldBindQuery(&spConfig)
	if err, respConfig := spConfigService.GetSpConfig(spConfig.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"respConfig": respConfig}, c)
	}
}

// GetSpConfigList 分页获取SpConfig列表
// @Tags SpConfig
// @Summary 分页获取SpConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SpConfigSearch true "分页获取SpConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spConfig/getSpConfigList [get]
func (spConfigApi *SpConfigApi) GetSpConfigList(c *gin.Context) {
	var pageInfo autocodeReq.SpConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := spConfigService.GetSpConfigInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
