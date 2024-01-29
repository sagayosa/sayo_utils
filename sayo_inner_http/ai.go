package sayoinnerhttp

import (
	"encoding/json"
	"fmt"
	"net/http"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
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

func GenerateAIDecisionRootCommandText(root []*module.Plugin, userCommand string) string {
	format := "Hello, I will provide a list of commands and the user's commands. Please respond to the user's commands based on what you think they want to call and the parameters you think they want to enter, in JSON format {root: the command you think the user wants to call, argvs: the parameters you think the user wants to fill in}. Note that the user's commands are the result of speech-to-text, so you may need to decide based on pronunciation. Thank you. Here is the list of commands:\n%v\nuser's command is: %v"
	return fmt.Sprintf(format, root, userCommand)
}

func PostAIDecisionRootCommand(aiAddr string, root []*module.Plugin, userCommand string) (result *AIDecisionResp, err error) {
	text := GenerateAIDecisionRootCommandText(root, userCommand)
	req := &AICompletionsReq{
		Messages: []Messages{
			{
				Role:    "user",
				Content: text,
			},
		},
	}
	code, body, err := utils.Post(utils.StringPlus("http://", aiAddr, constant.AICompletionsURL), req)
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

	mp := make(map[string]string)
	result = &AIDecisionResp{}
	if err := utils.UnMarshalUnknownAny(resp.Data, &mp); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(mp["content"]), result); err != nil {
		return nil, err
	}

	return result, nil
}
