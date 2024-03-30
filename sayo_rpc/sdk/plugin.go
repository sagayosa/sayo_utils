package sdk

import (
	"fmt"
	"net/http"

	"github.com/sagayosa/sayo_utils/module"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
	sayorpc "github.com/sagayosa/sayo_utils/sayo_rpc"
	"github.com/sagayosa/sayo_utils/utils"
)

func PostPlugin(plugin *module.Module, decision *AIDecisionResp) error {
	uri := ""
	for _, r := range plugin.Declare {
		if r.Root == decision.Root {
			uri = r.URL
		}
	}
	if uri == "" {
		return sayoerror.ErrMsg(sayoerror.ErrPostPluginNoUri, fmt.Sprintf("root = %v", decision.Root))
	}

	url := utils.StringPlus("http://", plugin.GetIPInfo(), uri)
	code, _, err := sayorpc.Post(url, decision.Argvs)
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrPostPluginFailed, code)
	}

	return nil
}
