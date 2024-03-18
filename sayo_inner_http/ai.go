package sayoinnerhttp

import (
	"encoding/json"
	"net/http"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
	"github.com/grteen/sayo_utils/utils"
)

type AICompletionsReq struct {
	Model       string     `json:"model"`
	Messages    []Messages `json:"messages"`
	Temperature int        `json:"temperature"`
	Stream      bool       `json:"stream"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AIDecisionResp struct {
	Root  string                 `json:"root"`
	Argvs map[string]interface{} `json:"argvs"`
}

func PostAICompletion(aiAddr string, content string) (result interface{}, err error) {
	req := &AICompletionsReq{
		Messages: []Messages{
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	code, body, err := sayorpc.Post(utils.StringPlus("http://", aiAddr, constant.AICompletionsURL), req)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return nil, sayoerror.ErrorInStatusCode(sayoerror.ErrAIChatFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return nil, sayoerror.ErrorInMsgCode(sayoerror.ErrAIChatFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data, nil
}
