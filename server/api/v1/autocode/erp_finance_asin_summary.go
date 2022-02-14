package autocode

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/global"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	autocodeReq "github.com/JoeyFrancisTribbiani/minerpserver/model/autocode/request"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/request"
	"github.com/JoeyFrancisTribbiani/minerpserver/model/common/response"
	"github.com/JoeyFrancisTribbiani/minerpserver/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ErpFinanceAsinSummaryApi struct {
}

var erpFinanceAsinSummaryService = service.ServiceGroupApp.AutoCodeServiceGroup.ErpFinanceAsinSummaryService

// CreateErpFinanceAsinSummary 创建ErpFinanceAsinSummary
// @Tags ErpFinanceAsinSummary
// @Summary 创建ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ErpFinanceAsinSummary true "创建ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /erpFinanceAsinSummary/createErpFinanceAsinSummary [post]
func (erpFinanceAsinSummaryApi *ErpFinanceAsinSummaryApi) CreateErpFinanceAsinSummary(c *gin.Context) {
	var erpFinanceAsinSummary autocode.ErpFinanceAsinSummary
	_ = c.ShouldBindJSON(&erpFinanceAsinSummary)
	if err := erpFinanceAsinSummaryService.CreateErpFinanceAsinSummary(erpFinanceAsinSummary); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteErpFinanceAsinSummary 删除ErpFinanceAsinSummary
// @Tags ErpFinanceAsinSummary
// @Summary 删除ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ErpFinanceAsinSummary true "删除ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /erpFinanceAsinSummary/deleteErpFinanceAsinSummary [delete]
func (erpFinanceAsinSummaryApi *ErpFinanceAsinSummaryApi) DeleteErpFinanceAsinSummary(c *gin.Context) {
	var erpFinanceAsinSummary autocode.ErpFinanceAsinSummary
	_ = c.ShouldBindJSON(&erpFinanceAsinSummary)
	if err := erpFinanceAsinSummaryService.DeleteErpFinanceAsinSummary(erpFinanceAsinSummary); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteErpFinanceAsinSummaryByIds 批量删除ErpFinanceAsinSummary
// @Tags ErpFinanceAsinSummary
// @Summary 批量删除ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /erpFinanceAsinSummary/deleteErpFinanceAsinSummaryByIds [delete]
func (erpFinanceAsinSummaryApi *ErpFinanceAsinSummaryApi) DeleteErpFinanceAsinSummaryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := erpFinanceAsinSummaryService.DeleteErpFinanceAsinSummaryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateErpFinanceAsinSummary 更新ErpFinanceAsinSummary
// @Tags ErpFinanceAsinSummary
// @Summary 更新ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ErpFinanceAsinSummary true "更新ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /erpFinanceAsinSummary/updateErpFinanceAsinSummary [put]
func (erpFinanceAsinSummaryApi *ErpFinanceAsinSummaryApi) UpdateErpFinanceAsinSummary(c *gin.Context) {
	var erpFinanceAsinSummary autocode.ErpFinanceAsinSummary
	_ = c.ShouldBindJSON(&erpFinanceAsinSummary)
	if err := erpFinanceAsinSummaryService.UpdateErpFinanceAsinSummary(erpFinanceAsinSummary); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindErpFinanceAsinSummary 用id查询ErpFinanceAsinSummary
// @Tags ErpFinanceAsinSummary
// @Summary 用id查询ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ErpFinanceAsinSummary true "用id查询ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /erpFinanceAsinSummary/findErpFinanceAsinSummary [get]
func (erpFinanceAsinSummaryApi *ErpFinanceAsinSummaryApi) FindErpFinanceAsinSummary(c *gin.Context) {
	var erpFinanceAsinSummary autocode.ErpFinanceAsinSummary
	_ = c.ShouldBindQuery(&erpFinanceAsinSummary)
	if err, reerpFinanceAsinSummary := erpFinanceAsinSummaryService.GetErpFinanceAsinSummary(erpFinanceAsinSummary.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reerpFinanceAsinSummary": reerpFinanceAsinSummary}, c)
	}
}

// GetErpFinanceAsinSummaryList 分页获取ErpFinanceAsinSummary列表
// @Tags ErpFinanceAsinSummary
// @Summary 分页获取ErpFinanceAsinSummary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ErpFinanceAsinSummarySearch true "分页获取ErpFinanceAsinSummary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /erpFinanceAsinSummary/getErpFinanceAsinSummaryList [get]
func (erpFinanceAsinSummaryApi *ErpFinanceAsinSummaryApi) GetErpFinanceAsinSummaryList(c *gin.Context) {
	var pageInfo autocodeReq.ErpFinanceAsinSummarySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := erpFinanceAsinSummaryService.GetErpFinanceAsinSummaryInfoList(pageInfo); err != nil {
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
