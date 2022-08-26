package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type QuestionBankProgrammCaseRouter struct {
}

// InitQuestionBankProgrammCaseRouter 初始化 QuestionBankProgrammCase 路由信息
func (s *QuestionBankProgrammCaseRouter) InitQuestionBankProgrammCaseRouter(Router *gin.RouterGroup) {
	questionBankProgrammCaseRouter := Router.Group("questionBankProgrammCase").Use(middleware.OperationRecord())
	questionBankProgrammCaseRouterWithoutRecord := Router.Group("questionBankProgrammCase")
	var questionBankProgrammCaseApi = v1.ApiGroupApp.QuestionBankApiGroup.QuestionBankProgrammCaseApi
	{
		questionBankProgrammCaseRouter.POST("createQuestionBankProgrammCase", questionBankProgrammCaseApi.CreateQuestionBankProgrammCase)             // 新建QuestionBankProgrammCase
		questionBankProgrammCaseRouter.DELETE("deleteQuestionBankProgrammCase", questionBankProgrammCaseApi.DeleteQuestionBankProgrammCase)           // 删除QuestionBankProgrammCase
		questionBankProgrammCaseRouter.DELETE("deleteQuestionBankProgrammCaseByIds", questionBankProgrammCaseApi.DeleteQuestionBankProgrammCaseByIds) // 批量删除QuestionBankProgrammCase
		questionBankProgrammCaseRouter.PUT("updateQuestionBankProgrammCase", questionBankProgrammCaseApi.UpdateQuestionBankProgrammCase)              // 更新QuestionBankProgrammCase
	}
	{
		questionBankProgrammCaseRouterWithoutRecord.GET("findQuestionBankProgrammCase", questionBankProgrammCaseApi.FindQuestionBankProgrammCase)       // 根据ID获取QuestionBankProgrammCase
		questionBankProgrammCaseRouterWithoutRecord.GET("getQuestionBankProgrammCaseList", questionBankProgrammCaseApi.GetQuestionBankProgrammCaseList) // 获取QuestionBankProgrammCase列表
	}
}
