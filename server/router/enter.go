package router

import (
	"github.com/prl26/exam-system/server/router/basicdata"
	"github.com/prl26/exam-system/server/router/examManage"
	"github.com/prl26/exam-system/server/router/lessondata"
	"github.com/prl26/exam-system/server/router/oj"
	"github.com/prl26/exam-system/server/router/questionBank"
	"github.com/prl26/exam-system/server/router/system"
	"github.com/prl26/exam-system/server/router/teachplan"
)

type RouterGroup struct {
	System       system.RouterGroup
	Basicdata    basicdata.RouterGroup
	Lessondata   lessondata.RouterGroup
	Teachplan    teachplan.RouterGroup
	Exammanage   examManage.RouterGroup
	QuestionBank questionBank.RouterGroup
	Oj           oj.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
