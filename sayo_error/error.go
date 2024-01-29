package sayoerror

import "fmt"

// return the error code and the error message
// If the error is not registered, then return the ErrInternalServer's code and message
func GetErrMsgByErr(err error) (int32, string) {
	code, ok := errorMp[err]
	if !ok {
		return errorMp[ErrInternalServer], ErrInternalServer.Error()
	}

	return code, err.Error()
}

// internal server error
var (
	ErrRegisterJobFailed = fmt.Errorf("register job failed")
	ErrRunModulesFailed  = fmt.Errorf("register job run modules failed")
)

// web info error
var (
	ErrInternalServer      = fmt.Errorf("internal server error")
	ErrDuplicateIdentifier = fmt.Errorf("duplicate identifier")
	ErrRegisterFailed      = fmt.Errorf("modules register failed")
	ErrUnknownType         = fmt.Errorf("unknown register type")
	ErrInvalidRole         = fmt.Errorf("invalid register role")
	ErrAIChatFailed        = fmt.Errorf("ai chat failed")
	ErrCoreGetRoleFailed   = fmt.Errorf("core get role failed")
)

const (
	SuccessCode = 200
	SuccessMsg  = "success"
)

var errorMp map[error]int32 = map[error]int32{
	ErrRegisterJobFailed: 500,
	ErrRunModulesFailed:  501,

	ErrInternalServer:      1000,
	ErrDuplicateIdentifier: 1001,
	ErrRegisterFailed:      1002,
	ErrUnknownType:         1003,
	ErrInvalidRole:         1004,
	ErrAIChatFailed:        1005,
	ErrCoreGetRoleFailed:   1006,
}

func Msg(err error, format string, a ...interface{}) error {
	return fmt.Errorf("%w, [msg]: %s", err, fmt.Sprintf(format, a...))
}
