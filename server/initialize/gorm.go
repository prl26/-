package initialize

import (
	"github.com/prl26/exam-system/server/global"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	"github.com/prl26/exam-system/server/model/system"
	"os"

	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/lessondata"
	"github.com/prl26/exam-system/server/model/teachplan"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		// 自动化模块表
		// Code generated by github.com/prl26/frontExam-system/server Begin; DO NOT EDIT.

		basicdata.Lesson{},
		basicdata.Chapter{},
		basicdata.Resource{},
		basicdata.Student{},
		basicdata.Class{},
		basicdata.College{},
		basicdata.LearnResourcesChapterMerge{},
		basicdata.Professional{},
		basicdata.TeachClass{},
		basicdata.StudentAndTeachClass{},
		basicdata.Term{},

		lessondata.VideoResources{},
		lessondata.ArticleResources{},
		lessondata.ResourcePractice{},
		lessondata.ResourcesTest{},
		lessondata.Knowledge{},

		teachplan.TeachAttendance{},
		teachplan.TeachAttendanceRecord{},
		teachplan.Score{},
		teachplan.ExamPlan{},

		examManage.PaperQuestionMerge{},
		examManage.ExamPaper{},
		examManage.PaperTemplateItem{},
		examManage.PaperTemplate{},
		examManage.ExamStudentPaper{},

		basicdata.Student{},
		basicdata.TeachClass{},
		//questionBank.ProgrammLanguageMerge{},
		examManage.ExamStudentPaper{},
		questionBankPo.PublicProgram{},
		// Code generated by github.com/prl26/frontExam-system/server End; DO NOT EDIT.
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
