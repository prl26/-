package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type TeachAttendanceRecordSearch struct {
	teachplan.TeachAttendanceRecord
	request.PageInfo
}
