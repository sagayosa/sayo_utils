package sayoinnerhttp

import (
	"net/http"

	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	"github.com/grteen/sayo_utils/utils"
)

func PostPlugin(plugin *module.Module, decision *AIDecisionResp) error {
	uri := ""
	for _, r := range plugin.Declare {
		if r.Root == decision.Root {
			uri = r.URL
		}
	}
	if uri == "" {
		return sayoerror.Msg(sayoerror.ErrPostPluginNoUri, "root = %v", decision.Root)
	}

	url := utils.StringPlus("http://", plugin.GetIPInfo(), uri)
	code, _, err := utils.Post(url, decision.Argvs)
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return sayoerror.ErrorInStatusCode(sayoerror.ErrAIChatFailed, code)
	}

	return nil
}
