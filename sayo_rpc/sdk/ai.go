package sdk

import (
	"github.com/sagayosa/sayo_utils/constant"
	sayorpc "github.com/sagayosa/sayo_utils/sayo_rpc"
	aitypes "github.com/sagayosa/sayo_utils/types/ai"
)

func PostAICompletion(aiAddr string, content string) (result interface{}, err error) {
	req := &aitypes.CompletionsReq{
		Messages: []aitypes.Messages{
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	return sayorpc.DirectPost(aiAddr, constant.AICompletionsURL, req)
}
