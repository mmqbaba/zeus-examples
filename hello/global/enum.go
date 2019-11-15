package global

import (
	"net/http"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"
)

// 每个子项目特有的错误码定义，避免使用 0 ~ 19999，与公共库冲突
const (
	ECodeSampleServiceOK errors.ErrorCode = iota + 20000
	ECodeSampleServiceErr
)

func init() {
	// ECodeMsg and ECodeStatus
	errors.ECodeMsg[ECodeSampleServiceOK] = "ECodeSampleServiceOK"
	errors.ECodeStatus[ECodeSampleServiceOK] = http.StatusOK

	errors.ECodeMsg[ECodeSampleServiceErr] = "ECodeSampleServiceErr"
	errors.ECodeStatus[ECodeSampleServiceErr] = http.StatusInternalServerError
}
