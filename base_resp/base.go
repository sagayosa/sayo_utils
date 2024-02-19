package baseresp

import (
	"fmt"

	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	"github.com/grteen/sayo_utils/utils"
)

type BaseResp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
	Data interface{} `json:"data"`
}

func (r *BaseResp) WithData(data interface{}) *BaseResp {
	r.Data = data
	return r
}

func (r *BaseResp) AppendMsg(format string, a ...interface{}) *BaseResp {
	r.Msg = utils.StringPlus(r.Msg, fmt.Sprintf(format, a...))
	return r
}

// func (r *BaseResp) ToString() string {
// 	bts, err := json.Marshal(r)
// 	sayolog.Msg("BaseResp Marshal failed").Err(err).Error()
// 	return string(bts)
// }

func NewBaseRespByError(err error) *BaseResp {
	code, msg := sayoerror.GetErrMsgByErr(err)

	return &BaseResp{
		Code: code,
		Msg:  msg,
		Err:  err.Error(),
	}
}

func NewSuccessResp(data interface{}) *BaseResp {
	return &BaseResp{
		Code: sayoerror.SuccessCode,
		Msg:  sayoerror.SuccessMsg,
		Data: data,
	}
}
