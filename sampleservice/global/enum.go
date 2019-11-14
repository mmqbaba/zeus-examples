package global

import (
	"net/http"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/enum"
)

// 每个子项目特有的错误码定义，避免使用 0 ~ 19999，与公共库冲突
const (
	ECodeSampleServiceOK enum.ErrorCode = iota + 20000
	ECodeSampleServiceErr
)

func init() {
	// ECodeMsg and ECodeStatus
	enum.ECodeMsg[ECodeSampleServiceOK] = "ECodeSampleServiceOK"
	enum.ECodeStatus[ECodeSampleServiceOK] = http.StatusOK

	enum.ECodeMsg[ECodeSampleServiceErr] = "ECodeSampleServiceErr"
	enum.ECodeStatus[ECodeSampleServiceErr] = http.StatusInternalServerError
}
