package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api/Backstage"
)

type JwtRouter struct{}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt")
	jwtApi := v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		jwtRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
