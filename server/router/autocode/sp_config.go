package autocode

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/api/v1"
	"github.com/JoeyFrancisTribbiani/minerpserver/middleware"
	"github.com/gin-gonic/gin"
)

type SpConfigRouter struct {
}

// InitSpConfigRouter 初始化 SpConfig 路由信息
func (s *SpConfigRouter) InitSpConfigRouter(Router *gin.RouterGroup) {
	spConfigRouter := Router.Group("spConfig").Use(middleware.OperationRecord())
	spConfigRouterWithoutRecord := Router.Group("spConfig")
	var spConfigApi = v1.ApiGroupApp.AutoCodeApiGroup.SpConfigApi
	{
		spConfigRouter.POST("createSpConfig", spConfigApi.CreateSpConfig)   // 新建SpConfig
		spConfigRouter.DELETE("deleteSpConfig", spConfigApi.DeleteSpConfig) // 删除SpConfig
		spConfigRouter.DELETE("deleteSpConfigByIds", spConfigApi.DeleteSpConfigByIds) // 批量删除SpConfig
		spConfigRouter.PUT("updateSpConfig", spConfigApi.UpdateSpConfig)    // 更新SpConfig
	}
	{
		spConfigRouterWithoutRecord.GET("findSpConfig", spConfigApi.FindSpConfig)        // 根据ID获取SpConfig
		spConfigRouterWithoutRecord.GET("getSpConfigList", spConfigApi.GetSpConfigList)  // 获取SpConfig列表
	}
}
