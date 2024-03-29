package sdk

import (
	"github.com/grteen/sayo_utils/constant"
	sayorpc "github.com/grteen/sayo_utils/sayo_rpc"
)

func PostAICompletion(aiAddr string, content string) (result interface{}, err error) {
	req := &AICompletionsReq{
		Messages: []Messages{
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	return sayorpc.DirectPost(aiAddr, constant.AICompletionsURL, req)
}
