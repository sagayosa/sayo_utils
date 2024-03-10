package sayoiris

import (
	"runtime/debug"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	sayolog "github.com/grteen/sayo_utils/sayo_log"
	"github.com/kataras/iris/v12"
)

type HandlerFunc func(iris.Context)

func IrisCtxJSONWrap(f func(ctx iris.Context) (*baseresp.BaseResp, error)) HandlerFunc {
	return func(ctx iris.Context) {
		resp, err := f(ctx)
		if err != nil {
			sayolog.Err(err).Msg("stack: %v", debug.Stack()).Error()
		}
		ctx.JSON(resp)
	}
}
