import service from '@/utils/request'

// @Tags ErpListingDetail
// @Summary 创建ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErpListingDetail true "创建ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /erpListingDetail/createErpListingDetail [post]
export const createErpListingDetail = (data) => {
  return service({
    url: '/erpListingDetail/createErpListingDetail',
    method: 'post',
    data
  })
}

// @Tags ErpListingDetail
// @Summary 删除ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErpListingDetail true "删除ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /erpListingDetail/deleteErpListingDetail [delete]
export const deleteErpListingDetail = (data) => {
  return service({
    url: '/erpListingDetail/deleteErpListingDetail',
    method: 'delete',
    data
  })
}

// @Tags ErpListingDetail
// @Summary 删除ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /erpListingDetail/deleteErpListingDetail [delete]
export const deleteErpListingDetailByIds = (data) => {
  return service({
    url: '/erpListingDetail/deleteErpListingDetailByIds',
    method: 'delete',
    data
  })
}

// @Tags ErpListingDetail
// @Summary 更新ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErpListingDetail true "更新ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /erpListingDetail/updateErpListingDetail [put]
export const updateErpListingDetail = (data) => {
  return service({
    url: '/erpListingDetail/updateErpListingDetail',
    method: 'put',
    data
  })
}

// @Tags ErpListingDetail
// @Summary 用id查询ErpListingDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ErpListingDetail true "用id查询ErpListingDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /erpListingDetail/findErpListingDetail [get]
export const findErpListingDetail = (params) => {
  return service({
    url: '/erpListingDetail/findErpListingDetail',
    method: 'get',
    params
  })
}

// @Tags ErpListingDetail
// @Summary 分页获取ErpListingDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ErpListingDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /erpListingDetail/getErpListingDetailList [get]
export const getErpListingDetailList = (params) => {
  return service({
    url: '/erpListingDetail/getErpListingDetailList',
    method: 'get',
    params
  })
}
