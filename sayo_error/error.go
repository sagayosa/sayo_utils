package sayoerror

import (
	"fmt"
)

// GetErrMsgByErr return the error code and the error message
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
	ErrRegisterJobFailed            = fmt.Errorf("register job failed")
	ErrRunModulesFailed             = fmt.Errorf("register job run modules failed")
	ErrGetAvailablePortTimesLimited = fmt.Errorf("get available port failed")
	ErrModuleRestart                = fmt.Errorf("module heart failed, try to restart")
	ErrModuleHeartFailed            = fmt.Errorf("module heart failed")
)

// web info error
var (
	ErrInternalServer             = fmt.Errorf("internal server error")
	ErrDuplicateIdentifier        = fmt.Errorf("duplicate identifier")
	ErrRegisterFailed             = fmt.Errorf("modules register failed")
	ErrUnknownType                = fmt.Errorf("unknown register type")
	ErrInvalidRole                = fmt.Errorf("invalid register role")
	ErrAIChatFailed               = fmt.Errorf("ai chat failed")
	ErrCoreGetRoleFailed          = fmt.Errorf("core get role failed")
	ErrDuplicateRootCommand       = fmt.Errorf("duplicate root command")
	ErrNoVoiceRecognizeModule     = fmt.Errorf("no voice recognize module")
	ErrNoVoiceGenerateModule      = fmt.Errorf("no voice generate module")
	ErrNoAIModule                 = fmt.Errorf("no ai module")
	ErrVoiceRecognizeFailed       = fmt.Errorf("voice recognize module failed")
	ErrPullCenterFailed           = fmt.Errorf("pull center failed")
	ErrPostPluginNoUri            = fmt.Errorf("no uri of this root command")
	ErrNoPluginOfRoot             = fmt.Errorf("no plugin of this root command")
	ErrCallCoreToPullCenterFailed = fmt.Errorf("call core to pull center failed")
	ErrCallCoreVoiceCommandFailed = fmt.Errorf("call core voice command failed")
	ErrNoCoreModule               = fmt.Errorf("no core module")
	ErrOpenNoIdentifier           = fmt.Errorf("open plugin dose not have this identifier")
	ErrOpenIdentifierIsNotAllowed = fmt.Errorf("this identifier in open plugin is disallowed")
	ErrOpenIdentifierDuplicated   = fmt.Errorf("this identifier in open plugin is duplicated")
	ErrOpenFileSelectorFailed     = fmt.Errorf("open file selector failed")
	ErrNoDesktopModule            = fmt.Errorf("no desktop module")
	ErrNoModule                   = fmt.Errorf("no module")
	ErrRegisterHotKeyFailed       = fmt.Errorf("register hot key failed")
	ErrNewWindowFailed            = fmt.Errorf("open new window failed")
	ErrTranslateFailed            = fmt.Errorf("translate failed")
	ErrWindowHideFailed           = fmt.Errorf("window hide failed")
	ErrWindowShowFailed           = fmt.Errorf("window show failed")
)

const (
	SuccessCode = 200
	SuccessMsg  = "success"
)

var errorMp map[error]int32 = map[error]int32{
	ErrRegisterJobFailed:            500,
	ErrRunModulesFailed:             501,
	ErrGetAvailablePortTimesLimited: 502,
	ErrModuleRestart:                503,
	ErrModuleHeartFailed:            504,

	ErrInternalServer:             1000,
	ErrDuplicateIdentifier:        1001,
	ErrRegisterFailed:             1002,
	ErrUnknownType:                1003,
	ErrInvalidRole:                1004,
	ErrAIChatFailed:               1005,
	ErrCoreGetRoleFailed:          1006,
	ErrDuplicateRootCommand:       1007,
	ErrNoVoiceRecognizeModule:     1008,
	ErrNoVoiceGenerateModule:      1009,
	ErrNoAIModule:                 1010,
	ErrPullCenterFailed:           1011,
	ErrPostPluginNoUri:            1012,
	ErrNoPluginOfRoot:             1013,
	ErrCallCoreToPullCenterFailed: 1014,
	ErrCallCoreVoiceCommandFailed: 1015,
	ErrNoCoreModule:               1016,
	ErrOpenNoIdentifier:           1017,
	ErrOpenIdentifierIsNotAllowed: 1018,
	ErrOpenIdentifierDuplicated:   1019,
	ErrOpenFileSelectorFailed:     1020,
	ErrNoDesktopModule:            1021,
	ErrNoModule:                   1022,
	ErrRegisterHotKeyFailed:       1023,
	ErrNewWindowFailed:            1024,
	ErrTranslateFailed:            1025,
	ErrWindowHideFailed:           1026,
}

func ErrMsg(err error, msg string) error {
	return fmt.Errorf("%w: %v", err, msg)
}
