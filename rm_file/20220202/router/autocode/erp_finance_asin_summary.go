package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ErpFinanceAsinSummaryRouter struct {
}

// InitErpFinanceAsinSummaryRouter 初始化 ErpFinanceAsinSummary 路由信息
func (s *ErpFinanceAsinSummaryRouter) InitErpFinanceAsinSummaryRouter(Router *gin.RouterGroup) {
	erpFinanceAsinSummaryRouter := Router.Group("erpFinanceAsinSummary").Use(middleware.OperationRecord())
	erpFinanceAsinSummaryRouterWithoutRecord := Router.Group("erpFinanceAsinSummary")
	var erpFinanceAsinSummaryApi = v1.ApiGroupApp.AutoCodeApiGroup.ErpFinanceAsinSummaryApi
	{
		erpFinanceAsinSummaryRouter.POST("createErpFinanceAsinSummary", erpFinanceAsinSummaryApi.CreateErpFinanceAsinSummary)   // 新建ErpFinanceAsinSummary
		erpFinanceAsinSummaryRouter.DELETE("deleteErpFinanceAsinSummary", erpFinanceAsinSummaryApi.DeleteErpFinanceAsinSummary) // 删除ErpFinanceAsinSummary
		erpFinanceAsinSummaryRouter.DELETE("deleteErpFinanceAsinSummaryByIds", erpFinanceAsinSummaryApi.DeleteErpFinanceAsinSummaryByIds) // 批量删除ErpFinanceAsinSummary
		erpFinanceAsinSummaryRouter.PUT("updateErpFinanceAsinSummary", erpFinanceAsinSummaryApi.UpdateErpFinanceAsinSummary)    // 更新ErpFinanceAsinSummary
	}
	{
		erpFinanceAsinSummaryRouterWithoutRecord.GET("findErpFinanceAsinSummary", erpFinanceAsinSummaryApi.FindErpFinanceAsinSummary)        // 根据ID获取ErpFinanceAsinSummary
		erpFinanceAsinSummaryRouterWithoutRecord.GET("getErpFinanceAsinSummaryList", erpFinanceAsinSummaryApi.GetErpFinanceAsinSummaryList)  // 获取ErpFinanceAsinSummary列表
	}
}
