package global

import (
	"net/http"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"
)

// 每个子项目特有的错误码定义，避免使用 0 ~ 19999，与公共库冲突
const (
	ECodeSampleServiceOK errors.ErrorCode = 20000
	ECodeSampleServiceErr errors.ErrorCode = 20001
)

func init() {
	// ECodeMsg and ECodeStatus
	errors.ECodeMsg[ECodeSampleServiceOK] = "ECodeSampleServiceOK"
	errors.ECodeStatus[ECodeSampleServiceOK] = http.StatusOK

	errors.ECodeMsg[ECodeSampleServiceErr] = "ECodeSampleServiceErr"
	errors.ECodeStatus[ECodeSampleServiceErr] = http.StatusInternalServerError
}
