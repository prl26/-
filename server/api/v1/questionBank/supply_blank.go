package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type QuestionBankSupplyBlankApi struct {
}

var questionBank_supply_blankService = service.ServiceGroupApp.QuestionBankServiceGroup.SupplyBlankService

// CreateQuestionBankSupplyBlank 创建QuestionBankSupplyBlank
// @Tags QuestionBankSupplyBlank
// @Summary 创建QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.SupplyBlank true "创建QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankSupplyBlank/createQuestionBankSupplyBlank [post]
func (questionBank_supply_blankApi *QuestionBankSupplyBlankApi) CreateQuestionBankSupplyBlank(c *gin.Context) {
	var questionBank_supply_blank questionBank.SupplyBlank
	_ = c.ShouldBindJSON(&questionBank_supply_blank)
	verify := utils.Rules{
		"Describe": {utils.NotEmpty()},
		"Is_order": {utils.NotEmpty()},
	}
	if err := utils.Verify(questionBank_supply_blank, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionBank_supply_blankService.CreateQuestionBankSupplyBlank(questionBank_supply_blank); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionBankSupplyBlank 删除QuestionBankSupplyBlank
// @Tags QuestionBankSupplyBlank
// @Summary 删除QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.SupplyBlank true "删除QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBankSupplyBlankdeleteQuestionBankSupplyBlank [delete]
func (questionBank_supply_blankApi *QuestionBankSupplyBlankApi) DeleteQuestionBankSupplyBlank(c *gin.Context) {
	var questionBank_supply_blank questionBank.SupplyBlank
	_ = c.ShouldBindJSON(&questionBank_supply_blank)
	if err := questionBank_supply_blankService.DeleteQuestionBankSupplyBlank(questionBank_supply_blank); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionBankSupplyBlankByIds 批量删除QuestionBankSupplyBlank
// @Tags QuestionBankSupplyBlank
// @Summary 批量删除QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionBankSupplyBlank/deleteQuestionBankSupplyBlankByIds [delete]
func (questionBank_supply_blankApi *QuestionBankSupplyBlankApi) DeleteQuestionBankSupplyBlankByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := questionBank_supply_blankService.DeleteQuestionBankSupplyBlankByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionBankSupplyBlank 更新QuestionBankSupplyBlank
// @Tags QuestionBankSupplyBlank
// @Summary 更新QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.SupplyBlank true "更新QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBankSupplyBlank/updateQuestionBankSupplyBlank [put]
func (questionBank_supply_blankApi *QuestionBankSupplyBlankApi) UpdateQuestionBankSupplyBlank(c *gin.Context) {
	var questionBank_supply_blank questionBank.SupplyBlank
	_ = c.ShouldBindJSON(&questionBank_supply_blank)
	verify := utils.Rules{
		"Describe": {utils.NotEmpty()},
		"Is_order": {utils.NotEmpty()},
	}
	if err := utils.Verify(questionBank_supply_blank, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionBank_supply_blankService.UpdateQuestionBankSupplyBlank(questionBank_supply_blank); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionBankSupplyBlank 用id查询QuestionBankSupplyBlank
// @Tags QuestionBankSupplyBlank
// @Summary 用id查询QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBank.SupplyBlank true "用id查询QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBankSupplyBlankfindQuestionBankSupplyBlank [get]
func (questionBank_supply_blankApi *QuestionBankSupplyBlankApi) FindQuestionBankSupplyBlank(c *gin.Context) {
	var questionBank_supply_blank questionBank.SupplyBlank
	_ = c.ShouldBindQuery(&questionBank_supply_blank)
	if requestionBank_supply_blank, err := questionBank_supply_blankService.GetQuestionBankSupplyBlank(questionBank_supply_blank.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionBank_supply_blank": requestionBank_supply_blank}, c)
	}
}

// GetQuestionBankSupplyBlankList 分页获取QuestionBankSupplyBlank列表
// @Tags QuestionBankSupplyBlank
// @Summary 分页获取QuestionBankSupplyBlank列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.QuestionBankSupplyBlankSearch true "分页获取QuestionBankSupplyBlank列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankSupplyBlank/getQuestionBankSupplyBlankList [get]
func (questionBank_supply_blankApi *QuestionBankSupplyBlankApi) GetQuestionBankSupplyBlankList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankSupplyBlankSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBank_supply_blankService.GetQuestionBankSupplyBlankInfoList(pageInfo); err != nil {
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
