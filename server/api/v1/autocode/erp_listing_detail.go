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

type ErpListingDetailApi struct {
}

var erpListingDetailService = service.ServiceGroupApp.AutoCodeServiceGroup.ErpListingDetailService

// CreateErpListingDetail 创建ErpListingDetail
// @Tags ErpListingDetail
// @Summary 创建ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ErpListingDetail true "创建ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /erpListingDetail/createErpListingDetail [post]
func (erpListingDetailApi *ErpListingDetailApi) CreateErpListingDetail(c *gin.Context) {
	var erpListingDetail autocode.ErpListingDetail
	_ = c.ShouldBindJSON(&erpListingDetail)
	if err := erpListingDetailService.CreateErpListingDetail(erpListingDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteErpListingDetail 删除ErpListingDetail
// @Tags ErpListingDetail
// @Summary 删除ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ErpListingDetail true "删除ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /erpListingDetail/deleteErpListingDetail [delete]
func (erpListingDetailApi *ErpListingDetailApi) DeleteErpListingDetail(c *gin.Context) {
	var erpListingDetail autocode.ErpListingDetail
	_ = c.ShouldBindJSON(&erpListingDetail)
	if err := erpListingDetailService.DeleteErpListingDetail(erpListingDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteErpListingDetailByIds 批量删除ErpListingDetail
// @Tags ErpListingDetail
// @Summary 批量删除ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /erpListingDetail/deleteErpListingDetailByIds [delete]
func (erpListingDetailApi *ErpListingDetailApi) DeleteErpListingDetailByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := erpListingDetailService.DeleteErpListingDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateErpListingDetail 更新ErpListingDetail
// @Tags ErpListingDetail
// @Summary 更新ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ErpListingDetail true "更新ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /erpListingDetail/updateErpListingDetail [put]
func (erpListingDetailApi *ErpListingDetailApi) UpdateErpListingDetail(c *gin.Context) {
	var erpListingDetail autocode.ErpListingDetail
	_ = c.ShouldBindJSON(&erpListingDetail)
	if err := erpListingDetailService.UpdateErpListingDetail(erpListingDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindErpListingDetail 用id查询ErpListingDetail
// @Tags ErpListingDetail
// @Summary 用id查询ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ErpListingDetail true "用id查询ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /erpListingDetail/findErpListingDetail [get]
func (erpListingDetailApi *ErpListingDetailApi) FindErpListingDetail(c *gin.Context) {
	var erpListingDetail autocode.ErpListingDetail
	_ = c.ShouldBindQuery(&erpListingDetail)
	if err, reerpListingDetail := erpListingDetailService.GetErpListingDetail(erpListingDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reerpListingDetail": reerpListingDetail}, c)
	}
}

// GetErpListingDetailList 分页获取ErpListingDetail列表
// @Tags ErpListingDetail
// @Summary 分页获取ErpListingDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ErpListingDetailSearch true "分页获取ErpListingDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /erpListingDetail/getErpListingDetailList [get]
func (erpListingDetailApi *ErpListingDetailApi) GetErpListingDetailList(c *gin.Context) {
	var pageInfo autocodeReq.ErpListingDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := erpListingDetailService.GetErpListingDetailInfoList(pageInfo); err != nil {
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
