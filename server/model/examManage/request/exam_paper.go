package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
)

type ExamPaperSearch struct {
	EpSearch
	request.PageInfo
}
type PaperDistribution struct {
	PlanId uint `json:"planId" form:"planId"`
}
type EpSearch struct {
	PlanId     int    `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	TemplateId int    `json:"templateId" form:"templateId" gorm:"column:template_id;comment:试卷模板Id;size:32;"`
	TermId     uint   `json:"termId" form:"termId"`
	LessonId   uint   `json:"lessonId" form:"lessonId"`
	UserId     uint   `json:"user_id" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
}
