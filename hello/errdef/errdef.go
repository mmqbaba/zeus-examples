package errdef

import (
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"
)

var (
	ErrOK = errors.New(errors.ECodeSuccessed, "OK", "")
)

