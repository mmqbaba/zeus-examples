package errdef

import (
	"net/http"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"
)

// 每个子项目特有的错误码定义，避免使用 0 ~ 19999，与公共库冲突
const (
	// 由公共库proto定义
	ECodeSuccessed errors.ErrorCode = 0
	ECodeSystem errors.ErrorCode = 10001

	// sampleservice
	ECodeSampleServiceErr errors.ErrorCode = 20001

	// sealservice
	ECodeSealServiceErr errors.ErrorCode = 30001
)

func init() {
	// ECodeMsg and ECodeStatus
	errors.ECodeMsg[ECodeSuccessed] = "ok"
	errors.ECodeStatus[ECodeSuccessed] = http.StatusOK
	errors.ECodeMsg[ECodeSystem] = "系统错误"
	errors.ECodeStatus[ECodeSystem] = http.StatusOK

	errors.ECodeMsg[ECodeSampleServiceErr] = "SampleServiceErr"
	errors.ECodeStatus[ECodeSampleServiceErr] = http.StatusInternalServerError

	errors.ECodeMsg[ECodeSealServiceErr] = "SealServiceErr"
	errors.ECodeStatus[ECodeSealServiceErr] = http.StatusInternalServerError
}