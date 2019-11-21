// Code generated by zeus-gen. DO NOT EDIT.
package errdef

import (
	net/http
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"
)

// 每个子项目特有的错误码定义，避免使用 0 ~ 19999，与公共库冲突
const (
	ErrCodeLogicBegin errors.ErrorCode = 40001
	ECodeSampleServiceErr errors.ErrorCode = 20001
	ECodeSealServiceErr errors.ErrorCode = 30001

)

func init() {
	// ECodeMsg and ECodeStatus
	errors.ECodeMsg[ErrCodeLogicBegin] = "错误码起始"
	errors.ECodeMsg[ECodeSampleServiceErr] = "SampleServiceErr"
	errors.ECodeStatus[ECodeSampleServiceErr] = http.StatusInternalServerError
	errors.ECodeMsg[ECodeSealServiceErr] = "SampleServiceErr"
	errors.ECodeStatus[ECodeSealServiceErr] = http.StatusInternalServerError

}
