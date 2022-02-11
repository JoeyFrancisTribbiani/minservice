package autocode

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/api/v1"
	"github.com/JoeyFrancisTribbiani/minerpserver/middleware"
	"github.com/gin-gonic/gin"
)

type ErpListingDetailRouter struct {
}

// InitErpListingDetailRouter 初始化 ErpListingDetail 路由信息
func (s *ErpListingDetailRouter) InitErpListingDetailRouter(Router *gin.RouterGroup) {
	erpListingDetailRouter := Router.Group("erpListingDetail").Use(middleware.OperationRecord())
	erpListingDetailRouterWithoutRecord := Router.Group("erpListingDetail")
	var erpListingDetailApi = v1.ApiGroupApp.AutoCodeApiGroup.ErpListingDetailApi
	{
		erpListingDetailRouter.POST("createErpListingDetail", erpListingDetailApi.CreateErpListingDetail)             // 新建ErpListingDetail
		erpListingDetailRouter.DELETE("deleteErpListingDetail", erpListingDetailApi.DeleteErpListingDetail)           // 删除ErpListingDetail
		erpListingDetailRouter.DELETE("deleteErpListingDetailByIds", erpListingDetailApi.DeleteErpListingDetailByIds) // 批量删除ErpListingDetail
		erpListingDetailRouter.PUT("updateErpListingDetail", erpListingDetailApi.UpdateErpListingDetail)              // 更新ErpListingDetail
	}
	{
		erpListingDetailRouterWithoutRecord.GET("findErpListingDetail", erpListingDetailApi.FindErpListingDetail)       // 根据ID获取ErpListingDetail
		erpListingDetailRouterWithoutRecord.GET("getErpListingDetailList", erpListingDetailApi.GetErpListingDetailList) // 获取ErpListingDetail列表
	}
}
