package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
)

type PublicProgramApi struct {
}

var publicProgramService = service.ServiceGroupApp.QuestionBankServiceGroup.PublicProgramService

// Create 创建公共编程题
func (p *PublicProgramApi) Create(c *gin.Context) {
	var req questionBankReq.PublicProgramCreate
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req.BasicModel, questionBankReq.BaseVerify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	programPo := questionBankPo.PublicProgram{}
	if len(req.ProgramCases) != 0 {
		programCaseStr, err := req.ProgramCases.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ProgramCases = programCaseStr
	} else {
		questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		return
	}
	if len(req.LanguageSupports) != 0 {
		languageSupportStr, err := req.LanguageSupports.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.LanguageSupports = languageSupportStr
	} else {
		questionBankResp.ErrorHandle(c, fmt.Errorf("未输入语言支持"))
		return
	}
	if len(req.ReferenceAnswers) != 0 {
		languageSupportStr, err := req.ReferenceAnswers.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ReferenceAnswers = languageSupportStr
	}
	if len(req.DefaultCodes) != 0 {
		languageSupportStr, err := req.DefaultCodes.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.DefaultCodes = languageSupportStr
	}
	programPo.BasicModel = req.BasicModel
	if err := publicProgramService.Create(&programPo); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	}
}

func (api *PublicProgramApi) FindList(c *gin.Context) {
	var pageInfo questionBankReq.PublicProgramSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := publicProgramService.FindList(pageInfo.PublicProgramSearchCriteria, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, err)
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (api *PublicProgramApi) FindDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if detail, err := publicProgramService.FindDetail(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, err)
	} else {
		if detail == nil {
			questionBankResp.NotFind(c)
			return
		} else {
			programDetail := questionBankResp.PublicProgramDetail{}
			programDetail.BasicModel = detail.BasicModel
			if err := programDetail.ProgramCases.Deserialize(detail.ProgramCases); err != nil {
				global.GVA_LOG.Error(err.Error())
				questionBankResp.ErrorHandle(c, err)
				return
			}
			if err := programDetail.LanguageSupports.Deserialization(detail.LanguageSupports); err != nil {
				global.GVA_LOG.Error(err.Error())
				questionBankResp.ErrorHandle(c, err)
				return
			}
			questionBankResp.OkWithDetailed(programDetail, "获取成功", c)
		}
	}
}

func (api *PublicProgramApi) Update(c *gin.Context) {
	var req questionBankReq.ProgramUpdate
	_ = c.ShouldBindJSON(&req)
	if req.Id == 0 {
		questionBankResp.CheckHandle(c, fmt.Errorf("请输入修改ID"))
		return
	}
	programPo := questionBankPo.Program{}
	if len(req.ProgramCases) != 0 {
		programCaseStr, err := req.ProgramCases.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ProgramCases = programCaseStr
	} else {
		// 修改的时候不一定修改编程题用例
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		//return
	}
	if len(req.LanguageSupports) != 0 {
		languageSupportStr, err := req.LanguageSupports.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.LanguageSupports = languageSupportStr
	} else {
		// 修改的时候不一定修改语言支持
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		//return
	}
	programPo.BasicModel = req.BasicModel
	programPo.CourseSupport = req.CourseSupport
	if err := programService.Update(&programPo); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	}
}
