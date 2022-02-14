import service from '@/utils/request'

// @Tags SpConfig
// @Summary 创建SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpConfig true "创建SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spConfig/createSpConfig [post]
export const createSpConfig = (data) => {
  return service({
    url: '/spConfig/createSpConfig',
    method: 'post',
    data
  })
}

// @Tags SpConfig
// @Summary 删除SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpConfig true "删除SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /spConfig/deleteSpConfig [delete]
export const deleteSpConfig = (data) => {
  return service({
    url: '/spConfig/deleteSpConfig',
    method: 'delete',
    data
  })
}

// @Tags SpConfig
// @Summary 删除SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /spConfig/deleteSpConfig [delete]
export const deleteSpConfigByIds = (data) => {
  return service({
    url: '/spConfig/deleteSpConfigByIds',
    method: 'delete',
    data
  })
}

// @Tags SpConfig
// @Summary 更新SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpConfig true "更新SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /spConfig/updateSpConfig [put]
export const updateSpConfig = (data) => {
  return service({
    url: '/spConfig/updateSpConfig',
    method: 'put',
    data
  })
}

// @Tags SpConfig
// @Summary 用id查询SpConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SpConfig true "用id查询SpConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /spConfig/findSpConfig [get]
export const findSpConfig = (params) => {
  return service({
    url: '/spConfig/findSpConfig',
    method: 'get',
    params
  })
}

// @Tags SpConfig
// @Summary 分页获取SpConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SpConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spConfig/getSpConfigList [get]
export const getSpConfigList = (params) => {
  return service({
    url: '/spConfig/getSpConfigList',
    method: 'get',
    params
  })
}
