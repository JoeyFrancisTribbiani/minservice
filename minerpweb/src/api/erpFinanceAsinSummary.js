import service from '@/utils/request'

// @Tags ErpFinanceAsinSummary
// @Summary 创建ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErpFinanceAsinSummary true "创建ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /erpFinanceAsinSummary/createErpFinanceAsinSummary [post]
export const createErpFinanceAsinSummary = (data) => {
  return service({
    url: '/erpFinanceAsinSummary/createErpFinanceAsinSummary',
    method: 'post',
    data
  })
}

// @Tags ErpFinanceAsinSummary
// @Summary 删除ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErpFinanceAsinSummary true "删除ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /erpFinanceAsinSummary/deleteErpFinanceAsinSummary [delete]
export const deleteErpFinanceAsinSummary = (data) => {
  return service({
    url: '/erpFinanceAsinSummary/deleteErpFinanceAsinSummary',
    method: 'delete',
    data
  })
}

// @Tags ErpFinanceAsinSummary
// @Summary 删除ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /erpFinanceAsinSummary/deleteErpFinanceAsinSummary [delete]
export const deleteErpFinanceAsinSummaryByIds = (data) => {
  return service({
    url: '/erpFinanceAsinSummary/deleteErpFinanceAsinSummaryByIds',
    method: 'delete',
    data
  })
}

// @Tags ErpFinanceAsinSummary
// @Summary 更新ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErpFinanceAsinSummary true "更新ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /erpFinanceAsinSummary/updateErpFinanceAsinSummary [put]
export const updateErpFinanceAsinSummary = (data) => {
  return service({
    url: '/erpFinanceAsinSummary/updateErpFinanceAsinSummary',
    method: 'put',
    data
  })
}

// @Tags ErpFinanceAsinSummary
// @Summary 用id查询ErpFinanceAsinSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ErpFinanceAsinSummary true "用id查询ErpFinanceAsinSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /erpFinanceAsinSummary/findErpFinanceAsinSummary [get]
export const findErpFinanceAsinSummary = (params) => {
  return service({
    url: '/erpFinanceAsinSummary/findErpFinanceAsinSummary',
    method: 'get',
    params
  })
}

// @Tags ErpFinanceAsinSummary
// @Summary 分页获取ErpFinanceAsinSummary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ErpFinanceAsinSummary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /erpFinanceAsinSummary/getErpFinanceAsinSummaryList [get]
export const getErpFinanceAsinSummaryList = (params) => {
  return service({
    url: '/erpFinanceAsinSummary/getErpFinanceAsinSummaryList',
    method: 'get',
    params
  })
}
